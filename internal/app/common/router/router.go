package router

import (
	"IdentifyService/internal/app/common/controller"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/pub", func(group *ghttp.RouterGroup) {
		group.Group("/captcha", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.Captcha,
			)
		})
	})
}
