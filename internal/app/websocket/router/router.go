package router

import (
	"IdentifyService/internal/app/system/controller"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"IdentifyService/library/libWebsocket"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/websocket", func(group *ghttp.RouterGroup) {
		group.GET("/", WsHandler)
	})
	// 启动websocket监听
	libWebsocket.Start()

	// 注册消息路由
	libWebsocket.RegisterMessageHandlers(libWebsocket.EventHandlers{
		"ping": controller.Ping.Ping, // 心跳
	})
}

// WsPageFinal 最终的WebSocket连接入口
func WsHandler(r *ghttp.Request) {
	// 1. 验证令牌
	userInfo, err := validateTokenAndGetUserInfoFinal(r)
	if err != nil {
		g.Log().Warningf(r.Context(), "WebSocket连接令牌验证失败: %v", err)
		r.Response.Status = http.StatusUnauthorized
		r.Response.WriteJson(g.Map{
			"code": http.StatusUnauthorized,
			"msg":  fmt.Sprintf("连接失败: %v", err),
		})
		return
	}

	// 2. 升级HTTP连接为WebSocket连接
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // 在生产环境中应该进行更严格的跨域检查
		},
	}

	conn, err := upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		g.Log().Errorf(r.Context(), "WebSocket升级失败: %v", err)
		return
	}

	// 3. 创建WebSocket客户端并绑定用户信息
	currentTime := gtime.Now().Unix()
	client := NewClientWithUserInfo(r, conn, currentTime, userInfo)

	// 4. 启动读写协程
	go client.Read()
	go client.Write()

	// 5. 注册客户端连接
	libWebsocket.Register(client)

	// 6. 记录连接日志
	g.Log().Infof(r.Context(), "用户建立WebSocket连接: userID=%s, userName=%s, clientID=%s, IP=%s",
		userInfo.ID, userInfo.Name, client.ID, client.IP)

	// 7. 推送连接成功消息
	data := &libWebsocket.DataConnectionEstablished{
		ClientID:         client.ID,
		UserID:           userInfo.ID,
		UserName:         userInfo.Name,
		FirstConnectTime: currentTime,
	}
	libWebsocket.SendSuccess(client, libWebsocket.EventConnectionEstablished, data)
}

// validateTokenAndGetUserInfoFinal 验证令牌并获取用户信息
func validateTokenAndGetUserInfoFinal(r *ghttp.Request) (*model.ContextUser, error) {
	// 1. 获取Authorization头
	authHeader := r.Request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("缺少Authorization头")
	}

	// 2. 解析Bearer令牌
	token := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		token = authHeader // 兼容直接传递令牌的情况
	}

	if token == "" {
		return nil, fmt.Errorf("令牌格式错误")
	}

	// 3. 验证令牌有效性
	ctx := r.Context()
	introspectRes, err := service.TokenService().Introspect(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("令牌验证失败: %v", err)
	}

	g.Log().Infof(ctx, "introspectRes: %+v", introspectRes)

	// 4. 检查用户信息
	if introspectRes == nil || introspectRes.UserID == "" {
		return nil, fmt.Errorf("用户信息无效")
	}

	// 5. 获取完整的用户信息（这里需要根据实际情况调整）
	// 由于IntrospectRes只包含基本信息，我们创建一个基本的用户信息
	wsUser := &model.ContextUser{
		UserLoginRes: &model.UserLoginRes{
			ID:       introspectRes.UserID,
			Name:     introspectRes.UserName,
			Nickname: introspectRes.UserName, // 使用用户名作为昵称
			Mobile:   "",                     // 需要从其他地方获取
			Status:   1,                      // 假设1表示启用状态
			IsAdmin:  false,                  // 需要从其他地方获取
			Avatar:   "",                     // 需要从其他地方获取
			OrgID:    introspectRes.OrdID,    // 注意这里是OrdID
		},
	}

	return wsUser, nil
}

// NewClientWithUserInfo 创建带有用户信息的WebSocket客户端
func NewClientWithUserInfo(r *ghttp.Request, socket *websocket.Conn, firstTime int64, userInfo *model.ContextUser) *libWebsocket.Client {
	client := &libWebsocket.Client{
		Addr:              socket.RemoteAddr().String(),
		ID:                guid.S(),
		Socket:            socket,
		Send:              make(chan *libWebsocket.WResponse, 100),
		SendClose:         false,
		CloseSignal:       make(chan struct{}, 1),
		FirstConnectTime:  firstTime,
		LastHeartbeatTime: firstTime,
		User:              userInfo,
		IP:                libUtils.GetClientIp(r.Context()),
		UserAgent:         r.UserAgent(),
	}

	return client
}

// 连接统计和管理
type ConnectionStats struct {
	TotalConnections int
	ActiveUsers      int
	ConnectionsByOrg map[string]int
}

// GetConnectionStats 获取连接统计信息
func GetConnectionStats() *ConnectionStats {
	stats := &ConnectionStats{
		TotalConnections: 0,
		ActiveUsers:      0,
		ConnectionsByOrg: make(map[string]int),
	}

	clients := libWebsocket.Manager().GetClients()
	userMap := make(map[string]bool)
	orgMap := make(map[string]int)

	for client := range clients {
		stats.TotalConnections++

		if client.User != nil && client.User.ID != "" {
			userMap[client.User.ID] = true

			if client.User.OrgID != "" {
				orgMap[client.User.OrgID]++
			}
		}
	}

	stats.ActiveUsers = len(userMap)
	stats.ConnectionsByOrg = orgMap

	return stats
}

/*
// DisconnectUser 断开指定用户的所有连接
func DisconnectUser(userID string) error {
	clients := libWebsocket.Manager().GetUserClient(userID)

	for _, client := range clients {
		// 发送断开连接消息
		libWebsocket.SendSuccess(client, "force_disconnect", map[string]interface{}{
			"reason": "管理员强制断开连接",
			"time":   gtime.Now().Unix(),
		})

		// 关闭连接
		libWebsocket.Close(client)
	}

	g.Log().Infof(context.Background(), "强制断开用户连接: userID=%s, 连接数=%d", userID, len(clients))
	return nil
}

// BroadcastToOrg 向指定组织推送消息
func BroadcastToOrg(orgID string, event string, data interface{}) error {
	response := &libWebsocket.WResponse{
		Event:     event,
		Data:      data,
		Code:      200,
		Timestamp: gtime.Now().Unix(),
	}

	clients := libWebsocket.Manager().GetClients()
	count := 0

	for client := range clients {
		if client.User != nil && client.User.OrgID == orgID {
			client.SendMsg(response)
			count++
		}
	}

	g.Log().Infof(context.Background(), "向组织推送消息: orgID=%s, event=%s, 目标连接数=%d", orgID, event, count)
	return nil
}
*/
