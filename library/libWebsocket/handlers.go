package libWebsocket

import (
	"context"
	"runtime/debug"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
)

// handleMessage 处理消息
func handleMessage(client *Client, messageType int, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(mctx, "handlerMsg recover, err:%+v, stack:%+v", r, string(debug.Stack()))
		}
	}()

	switch messageType {
	case websocket.PingMessage:
		handlePingMessage(client, message)
	case websocket.CloseMessage:
		handleCloseMessage(client, message)
	case websocket.BinaryMessage:
		handleBinaryMessage(client, message)
	default:
		g.Log().Warningf(mctx, "handlerMsg messageType:%v, message:%+v", messageType, string(message))
	}

}

func handlePingMessage(client *Client, message []byte) {
	client.SetHeartbeat(gtime.Now().Unix())
}

func handleCloseMessage(client *Client, message []byte) {
	Manager().DelClients(client)
	client.close()
}

func handleBinaryMessage(client *Client, message []byte) {
	var request *WRequest
	if err := gconv.Struct(message, &request); err != nil {
		g.Log().Warningf(mctx, "handlerMsg 数据解析失败,err:%+v, message:%+v", err, string(message))
		return
	}

	if request.Event == "" {
		g.Log().Warning(mctx, "handlerMsg request.Event is null")
		return
	}

	fun, ok := routers[WebSocketEvent(request.Event)]
	if !ok {
		g.Log().Warningf(mctx, "handlerMsg function id %v: not registered", request.Event)
		return
	}

	err := msgGo.AddWithRecover(mctx,
		func(ctx context.Context) {
			fun(client, request)
		},
		func(ctx context.Context, err error) {
			g.Log().Warningf(mctx, "handlerMsg msgGo exec err:%+v", err)
		},
	)

	if err != nil {
		g.Log().Warningf(mctx, "handlerMsg msgGo Add err:%+v", err)
		return
	}
}

// RegisterMessageHandlers 注册消息器
func RegisterMessageHandlers(handlers EventHandlers) {
	for event, f := range handlers {
		if _, ok := routers[event]; ok {
			g.Log().Fatalf(mctx, "RegisterMsg function id %v: already registered", event)
			return
		}
		routers[event] = f
	}
}
