package router

import (
	websocketRouter "IdentifyService/internal/app/websocket/router"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindWebsocketModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	websocketRouter.R.BindController(ctx, group)
}
