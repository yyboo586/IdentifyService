package router

import (
	"context"
	"errors"
	"net/http"

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
		group.Hook("/*", ghttp.HookAfterOutput, service.Log().OperationLog)
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
	data, err := service.TokenService().Parse(r)
	if err != nil {
		if errors.Is(err, service.ErrTokenInvalid) || errors.Is(err, service.ErrTokenExpired) {
			r.Response.Status = http.StatusOK
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		} else if errors.Is(err, service.ErrTokenNotActive) {
			r.Response.Status = http.StatusOK
			r.Response.WriteJson(g.Map{
				"code":    http.StatusForbidden,
				"message": "Forbidden",
			})
		}
		g.Log().Error(r.GetCtx(), err)
		r.ExitAll()
	}
	g.Log().Debug(r.Context(), "InjectCtx: ", data)
	if data != nil {
		customCtx := new(model.Context)
		err = gconv.Struct(data, &customCtx.User)
		if err != nil {
			g.Log().Error(r.GetCtx(), err)
			r.ExitAll()
		}
		service.ContextService().Init(r, customCtx)
	}
	r.Middleware.Next()
}
