package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysNoticeReadController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysNoticeRead", func(group *ghttp.RouterGroup) {
		group.Bind(
		//controller.SysNoticeRead,
		)
	})
}
