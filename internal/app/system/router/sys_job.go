package router

import (
	"IdentifyService/internal/app/system/controller"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysJobController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysJob", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysJob,
		)
	})
}
