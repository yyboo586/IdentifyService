package router

import (
	"context"
	"errors"

	"IdentifyService/internal/app/system/controller"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libResponse"
	"IdentifyService/library/libRouter"
	"IdentifyService/library/liberr"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/identify-service", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.AuthController,
		)
		// context拦截器
		group.Middleware(InjectCtx)
		// 自动绑定定义的控制器
		if err := libRouter.RouterAutoBindBefore(ctx, router, group); err != nil {
			panic(err)
		}
		// group.Middleware(Auth)
		// 后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.Log().OperationLog)
		group.Bind(
			controller.UserController,
			controller.MenuController,
			controller.RoleController,
			controller.OrgController,
			controller.LogController,
		)
		// 自动绑定定义的控制器
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
	g.Log().Info(ctx, "[DEBUG] data: ", "data", data)
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

func Auth(r *ghttp.Request) {
	ctx := r.GetCtx()
	operatorInfo := service.ContextService().Get(ctx)
	url := gstr.TrimLeft(r.Request.URL.Path, "/")

	if service.User().IsSuperAdmin(ctx, operatorInfo.User.ID) {
		r.Middleware.Next()
		return
	}

	// 获取地址对应的菜单id
	menuList, err := service.AuthRule().GetAllAuthRules(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		libResponse.FailJson(true, r, "请求数据失败")
	}
	var menu *model.AuthRule
	for _, m := range menuList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == url || ms == url {
			menu = m
			break
		}
	}
	//只验证存在数据库中的规则
	if menu != nil {
		//若是不登录能访问的接口则不判断权限
		excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
		for _, p := range excludePaths {
			if gstr.Equal(menu.Name, gstr.TrimLeft(p, "/")) {
				r.Middleware.Next()
				return
			}
		}
		menuId := menu.ID
		//菜单没存数据库不验证权限
		if menuId != 0 {
			//判断权限操作
			err = checkAuth(ctx, operatorInfo.User.ID, menuId)
			if err != nil {
				libResponse.FailJson(true, r, err.Error())
			}
		}
	}
	r.Middleware.Next()
}

// checkAuth 检查用户是否有访问指定菜单的权限
// 这是CASBIN权限检查的核心方法
//
// CASBIN权限检查流程：
// 1. 获取用户的所有角色ID
// 2. 使用Enforce方法检查角色是否有访问指定资源的权限
// 3. Enforce(role, resource, action) 返回true表示有权限，false表示无权限
//
// 参数说明：
// - ctx: 上下文
// - adminId: 用户ID
// - menuId: 菜单ID（资源ID）
//
// 返回值：
// - err: 错误信息，如果用户没有权限则返回错误
func checkAuth(ctx context.Context, userID string, menuId int64) (err error) {
	var (
		roleIds    []int64
		roleIdsMap = gmap.New()
		enforcer   *casbin.SyncedEnforcer
		b          bool
	)

	err = g.Try(ctx, func(ctx context.Context) {
		// 1. 获取用户的所有角色ID
		roleIds, err = service.Role().GetRoleIDsByUserID(ctx, userID)
		liberr.ErrIsNil(ctx, err)

		// 2. 将角色ID转换为map，用于快速查找
		for _, v := range roleIds {
			roleIdsMap.Set(v, v)
		}

		// 3. 获取CASBIN执行器实例
		enforcer, err = service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err)

		// 4. 遍历用户的每个角色，检查是否有访问权限
		roleIdsMap.Iterator(func(k interface{}, v interface{}) bool {
			// 使用Enforce方法检查权限
			// Enforce(role, resource, action) 检查角色是否有对资源的指定操作权限
			// 参数说明：
			// - gconv.String(v): 角色ID（字符串格式）
			// - gconv.String(menuId): 菜单ID（资源ID，字符串格式）
			// - "All": 操作类型，表示所有操作
			b, err = enforcer.Enforce(gconv.String(v), gconv.String(menuId), "All")
			liberr.ErrIsNil(ctx, err)

			// 如果当前角色有权限，继续检查下一个角色
			// 如果当前角色没有权限，继续检查下一个角色
			// 只有当所有角色都没有权限时，才返回false
			return !b
		})

		// 5. 如果所有角色都没有权限，则返回错误
		if !b {
			liberr.ErrIsNil(ctx, errors.New("没有权限"))
		}
	})

	return
}
