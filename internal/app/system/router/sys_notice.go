package router

import (
	"context"

	"IdentifyService/internal/app/system/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysNoticeController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysNotice", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysNotice,
		)
	})
}
