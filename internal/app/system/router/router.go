package router

import (
	"context"

	"IdentifyService/internal/app/system/controller"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libRouter"

	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(
			//登录
			controller.Login,
		)
		//context拦截器
		group.Middleware(service.Middleware().Ctx)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBindBefore(ctx, router, group); err != nil {
			panic(err)
		}
		group.Middleware(service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		group.Bind(
			controller.User,
			controller.Menu,
			controller.Role,
			controller.Dept,
			controller.Post,
			controller.DictType,
			controller.DictData,
			controller.Config,
			controller.Monitor,
			controller.LoginLog,
			controller.OperLog,
			controller.ToolsGenTable,
			controller.Personal,
			controller.UserOnline,
			controller.Cache,   // 缓存处理
			controller.Upload,  // 普通文件上传
			controller.UEditor, //编辑器
			controller.File,    // 文件管理
			controller.Log,
			controller.Agreement,
			controller.UserDevice,
		)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
