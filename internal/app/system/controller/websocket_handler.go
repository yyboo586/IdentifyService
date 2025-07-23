package controller

import (
	"IdentifyService/library/libWebsocket"
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var WebSocketHandler = new(websocketHandler)

type websocketHandler struct{}

// ==================== 基础消息处理 ====================

// Login 用户登录
// 使用场景：
// 1. 客户端建立WebSocket连接后，需要绑定用户身份
// 2. 用户切换账号时，重新绑定新的用户信息
// 3. 多设备登录场景，每个设备都需要独立的用户身份绑定
// 4. 临时访客用户需要升级为正式用户身份
func (c *websocketHandler) Login(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	userID := req.Data["user_id"].(string)
	if userID == "" {
		libWebsocket.SendError(client, req.Event, errors.New("用户ID不能为空"))
		return
	}

	// 设置用户信息
	client.User.ID = userID
	client.User.Name = req.Data["user_name"].(string)

	// 注册用户登录
	// login := &libWebsocket.Login{
	// 	UserId: userID,
	// 	Client: client,
	// }
	// libWebsocket.Manager().Login <- login

	// 发送登录成功响应
	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"user_id":    userID,
		"login_time": gtime.Now().Unix(),
		"client_id":  client.ID,
	})

	g.Log().Infof(context.Background(), "用户登录成功: userID=%s, clientID=%s", userID, client.ID)
}

// Logout 用户登出
// 使用场景：
// 1. 用户主动退出登录，清理连接状态
// 2. 用户切换账号前，先登出当前账号
// 3. 管理员强制用户下线时，客户端主动登出
// 4. 用户会话过期，需要重新登录
// 5. 多设备登录限制，新设备登录时旧设备自动登出
func (c *websocketHandler) Logout(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	if client.User == nil || client.User.ID == "" {
		libWebsocket.SendError(client, req.Event, errors.New("用户未登录"))
		return
	}

	userID := client.User.ID
	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"user_id":     userID,
		"logout_time": gtime.Now().Unix(),
	})

	g.Log().Infof(context.Background(), "用户登出: userID=%s, clientID=%s", userID, client.ID)
}

// Heartbeat 心跳处理
// 使用场景：
// 1. 保持WebSocket连接活跃，防止连接超时断开
// 2. 检测客户端在线状态，及时发现断线用户
// 3. 网络质量监控，通过心跳延迟判断网络状况
// 4. 负载均衡场景，通过心跳判断客户端健康状态
// 5. 移动端应用，网络切换时保持连接稳定
func (c *websocketHandler) Heartbeat(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	currentTime := gtime.Now().Unix()
	client.SetHeartbeat(currentTime)

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"timestamp": currentTime,
		"client_id": client.ID,
	})
}

// GetClientInfo 获取客户端信息
// 使用场景：
// 1. 客户端需要获取自己的连接状态和配置信息
// 2. 调试和故障排查，查看客户端详细连接信息
// 3. 客户端重新连接后，验证连接状态是否正确
// 4. 管理员查看特定客户端的连接详情
// 5. 客户端应用启动时，获取连接配置信息
func (c *websocketHandler) GetClientInfo(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	info := map[string]interface{}{
		"client_id":           client.ID,
		"user_id":             client.User.ID,
		"user_name":           client.User.Name,
		"ip":                  client.IP,
		"user_agent":          client.UserAgent,
		"first_connect_time":  client.FirstConnectTime,
		"last_heartbeat_time": client.LastHeartbeatTime,
	}

	libWebsocket.SendSuccess(client, req.Event, info)
}

// ==================== 订阅管理 ====================

// Subscribe 订阅事件
// 使用场景：
// 1. 用户订阅系统通知，接收重要消息推送
// 2. 用户订阅特定业务事件，如订单状态变更
// 3. 用户加入聊天室，订阅聊天消息
// 4. 用户订阅实时数据更新，如股票价格、天气信息
// 5. 用户订阅工作流通知，如审批流程状态
// 6. 用户订阅安全告警，如异常登录提醒
func (c *websocketHandler) Subscribe(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	event := req.Data["event"].(string)
	if event == "" {
		libWebsocket.SendError(client, req.Event, errors.New("事件名称不能为空"))
		return
	}

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"event":     event,
		"client_id": client.ID,
	})

	g.Log().Infof(context.Background(), "客户端订阅事件: clientID=%s, event=%s", client.ID, event)
}

// Unsubscribe 取消订阅
// 使用场景：
// 1. 用户取消订阅系统通知，减少消息干扰
// 2. 用户退出聊天室，取消订阅聊天消息
// 3. 用户不再关注特定业务事件
// 4. 用户关闭实时数据更新功能
// 5. 用户完成工作流，取消相关通知
// 6. 用户调整通知偏好，取消某些类型的通知
func (c *websocketHandler) Unsubscribe(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	event := req.Data["event"].(string)
	if event == "" {
		libWebsocket.SendError(client, req.Event, errors.New("事件名称不能为空"))
		return
	}

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"event":     event,
		"client_id": client.ID,
	})

	g.Log().Infof(context.Background(), "客户端取消订阅: clientID=%s, event=%s", client.ID, event)
}

// ==================== 消息历史 ====================

// GetMessageHistory 获取消息历史
// 使用场景：
// 1. 用户重新连接后，获取离线期间错过的消息
// 2. 用户查看历史通知和系统消息
// 3. 聊天应用中，用户查看历史聊天记录
// 4. 用户搜索特定时间段的消息
// 5. 管理员查看系统消息发送历史
// 6. 审计和合规要求，查看消息发送记录
func (c *websocketHandler) GetMessageHistory(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	// 这里可以实现获取消息历史的逻辑
	// 目前返回空列表
	history := []map[string]interface{}{}

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"history":   history,
		"client_id": client.ID,
	})
}

// ==================== 在线状态 ====================

// GetOnlineStatus 获取在线状态
// 使用场景：
// 1. 用户查看好友或同事的在线状态
// 2. 聊天应用中显示联系人是否在线
// 3. 协作工具中显示团队成员在线状态
// 4. 客服系统中查看客服人员是否在线
// 5. 游戏应用中查看好友是否在线
// 6. 管理员监控用户在线情况
func (c *websocketHandler) GetOnlineStatus(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	userID := req.Data["user_id"].(string)
	if userID == "" {
		libWebsocket.SendError(client, req.Event, errors.New("用户ID不能为空"))
		return
	}

	// 检查用户是否在线
	clients := libWebsocket.Manager().GetUserClient(userID)
	online := len(clients) > 0

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"user_id": userID,
		"online":  online,
	})
}

// GetOnlineUsers 获取在线用户列表
// 使用场景：
// 1. 聊天室显示当前在线用户列表
// 2. 协作工具显示团队成员在线状态
// 3. 客服系统显示可用的客服人员
// 4. 游戏大厅显示在线玩家列表
// 5. 管理员查看系统在线用户统计
// 6. 实时监控系统显示活跃用户
func (c *websocketHandler) GetOnlineUsers(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	// 这里可以实现获取在线用户列表的逻辑
	// 目前返回空列表
	users := []map[string]interface{}{}

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"users":     users,
		"client_id": client.ID,
	})
}

// ==================== 错误处理 ====================

// HandleError 处理错误
// 使用场景：
// 1. 客户端遇到异常情况，向服务器报告错误
// 2. 客户端网络异常，报告连接问题
// 3. 客户端应用崩溃，报告错误信息
// 4. 客户端功能异常，报告具体错误
// 5. 客户端性能问题，报告性能指标
// 6. 客户端安全事件，报告可疑行为
func (c *websocketHandler) HandleError(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	errorMsg := req.Data["message"].(string)
	if errorMsg == "" {
		errorMsg = "未知错误"
	}

	g.Log().Warningf(context.Background(), "客户端报告错误: clientID=%s, error=%s", client.ID, errorMsg)

	libWebsocket.SendSuccess(client, req.Event, map[string]interface{}{
		"received":  true,
		"client_id": client.ID,
	})
}

// ==================== 调试信息 ====================

// Debug 调试信息
// 使用场景：
// 1. 开发调试阶段，查看客户端详细状态
// 2. 故障排查，分析连接问题
// 3. 性能分析，查看连接性能指标
// 4. 安全审计，查看连接安全信息
// 5. 系统监控，查看客户端运行状态
// 6. 技术支持，帮助用户解决连接问题
func (c *websocketHandler) Debug(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	debugInfo := map[string]interface{}{
		"client_id":           client.ID,
		"user_id":             client.User.ID,
		"user_name":           client.User.Name,
		"ip":                  client.IP,
		"user_agent":          client.UserAgent,
		"first_connect_time":  client.FirstConnectTime,
		"last_heartbeat_time": client.LastHeartbeatTime,
		"send_close":          client.SendClose,
		"timestamp":           gtime.Now().Unix(),
	}

	libWebsocket.SendSuccess(client, req.Event, debugInfo)
}
