package router

import (
	"context"

	"IdentifyService/internal/app/system/controller"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libRouter"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/identify-service", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.AuthController,
		)
		group.Middleware(InjectCtx)
		if err := libRouter.RouterAutoBindBefore(ctx, router, group); err != nil {
			panic(err)
		}
		// group.Hook("/*", ghttp.HookAfterOutput, service.Log().OperationLog)
		group.Bind(
			controller.UserController,
			controller.MenuController,
			controller.RoleController,
			controller.OrgController,
			controller.LogController,
		)
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}

func InjectCtx(r *ghttp.Request) {
	ctx := r.GetCtx()
	data, err := service.TokenService().ParseToken(r)
	if err != nil {
		g.Log().Error(ctx, err)
		r.Middleware.Next()
		return
	}
	if data != nil {
		customCtx := new(model.Context)
		err = gconv.Struct(data.Data, &customCtx.User)
		if err != nil {
			g.Log().Error(ctx, err)
			r.Middleware.Next()
			return
		}
		service.ContextService().Init(r, customCtx)
	}
	r.Middleware.Next()
}
