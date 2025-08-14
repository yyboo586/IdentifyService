package libWebsocket

import (
	systemModel "IdentifyService/internal/app/system/model"
	"context"
	"runtime/debug"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 5 * 60
)

// 用户登录
type login struct {
	UserId string
	Client *Client
}

// GetKey 读取客户端数据
func (l *login) GetKey() string {
	return GetUserKey(l.UserId)
}

// Client 客户端连接
type Client struct {
	Addr              string                   // 客户端地址
	ID                string                   // 连接唯一标识
	Socket            *websocket.Conn          // 用户连接
	Send              chan *WResponse          // 待发送的数据
	SendClose         bool                     // 发送是否关闭
	CloseSignal       chan struct{}            // 关闭信号
	FirstConnectTime  int64                    // 首次连接时间
	LastHeartbeatTime int64                    // 用户上次心跳时间
	User              *systemModel.ContextUser // 用户信息
	context           context.Context          // Custom context for internal usage purpose.
	IP                string                   // 客户端IP
	UserAgent         string                   // 用户代理
}

// 读取客户端数据
func (c *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(mctx, "client read err: %+v, stack:%+v, user:%+v", r, string(debug.Stack()), c.User)
		}
	}()

	defer c.close()

	for {
		messageType, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		// 处理消息
		handleMessage(c, messageType, message)
	}
}

// 向客户端写数据
func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(mctx, "client write err: %+v, stack:%+v, user:%+v", r, string(debug.Stack()), c.User)
		}
	}()
	defer func() {
		clientManager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		select {
		case <-c.CloseSignal:
			g.Log().Infof(mctx, "websocket client quit, user:%+v", c.User)
			return
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				g.Log().Warningf(mctx, "client write message, user:%+v", c.User)
				return
			}
			_ = c.Socket.WriteJSON(message)
		}
	}
}

// SendMsg 发送数据
func (c *Client) SendMsg(msg *WResponse) {
	if c == nil || c.SendClose {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			g.Log().Infof(mctx, "SendMsg err:%+v, stack:%+v", r, string(debug.Stack()))
		}
	}()
	c.Send <- msg
}

// Context is alias for function GetCtx.
func (c *Client) Context() context.Context {
	if c.context == nil {
		c.context = gctx.New()
	}
	return c.context
}

// Heartbeat 心跳更新
func (c *Client) SetHeartbeat(currentTime int64) {
	c.LastHeartbeatTime = currentTime
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout(currentTime int64) (timeout bool) {
	if c.LastHeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 关闭客户端
func (c *Client) close() {
	if c.SendClose {
		return
	}
	c.SendClose = true
	c.CloseSignal <- struct{}{}
}

// Close 关闭指定客户端连接
func Close(client *Client) {
	client.close()
}

// SendSuccess 发送成功消息
func SendSuccess(client *Client, event WebSocketEvent, data ...interface{}) {
	d := interface{}(nil)
	if len(data) > 0 {
		d = data[0]
	}
	client.SendMsg(&WResponse{
		Event:     event,
		Data:      d,
		Code:      gcode.CodeOK.Code(),
		Timestamp: gtime.Now().Unix(),
	})
	before(client)
}

// SendError 发送错误消息
func SendError(client *Client, event WebSocketEvent, err error) {
	client.SendMsg(&WResponse{
		Event:     event,
		Code:      gcode.CodeNil.Code(),
		ErrorMsg:  err.Error(),
		Timestamp: gtime.Now().Unix(),
	})
	before(client)
}

// before
func before(client *Client) {
	client.SetHeartbeat(gtime.Now().Unix())
}
