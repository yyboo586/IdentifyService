package router

import (
	"context"

	"IdentifyService/internal/app/system/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysAttachmentController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysAttachment", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysAttachment,
		)
	})
}
