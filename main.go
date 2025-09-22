package main

import (
	commonDomain "IdentifyService/internal/common/domain"
	commonModel "IdentifyService/internal/common/model"
	commonAppService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/domain"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/service"
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	g.Log().SetTimeFormat("2006-01-02 15:04:05")

	s := g.Server()
	s.SetAddr(g.Cfg().MustGet(context.Background(), "server.address").String())
	s.SetOpenApiPath(g.Cfg().MustGet(context.Background(), "server.openapiPath").String())
	s.SetSwaggerPath(g.Cfg().MustGet(context.Background(), "server.swaggerPath").String())

	commonAppService.RegisterMQService()
	commonAppService.RegisterContextService()
	commonAppService.RegisterLog(commonDomain.NewLog())
	commonAppService.RegisterTokenService(commonDomain.NewToken())

	domainRole := domain.NewRole(nil)
	domainUser := domain.NewUser(nil)
	domainOrg := domain.NewOrg(domainUser)
	domainPermission := domain.NewPermission(nil)

	// 初始化权限点数据
	domainPermission.RegisterPermissionPoints(context.Background(), model.UserPermissions)
	domainPermission.RegisterPermissionPoints(context.Background(), model.RolePermissions)
	domainPermission.RegisterPermissionPoints(context.Background(), model.OrgPermissions)

	err := domainPermission.Init(context.Background())
	if err != nil {
		panic(err)
	}

	authAppService := service.NewAuthService(domainUser, domainOrg, domainPermission)
	orgAppService := service.NewOrgService(domainUser, domainOrg, domainRole, domainPermission)

	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(commonAppService.CORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Group("/identify-service", func(group *ghttp.RouterGroup) {
			group.Bind(
				authAppService,
			)
			group.Middleware(InjectCtx)
			group.Hook("/*", ghttp.HookAfterOutput, commonAppService.Log().RecordLog)
			group.Bind(
				orgAppService,
				service.LogService,
			)
		})
	})

	s.Run()
}

func GetBearerTokenFromRequest(r *ghttp.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", fmt.Errorf("缺少Authorization头")
	}
	return authorization, nil
}

func InjectCtx(r *ghttp.Request) {
	ctx := r.GetCtx()

	excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
	if slices.Contains(excludePaths, r.URL.Path) {
		r.Middleware.Next()
		return
	}

	tokenStr, err := GetBearerTokenFromRequest(r)
	if err != nil {
		r.ExitAll()
	}
	customClaims, err := commonAppService.TokenService().Parse(ctx, tokenStr)
	if err != nil {
		if errors.Is(err, commonAppService.ErrTokenInvalid) {
			r.Response.Status = http.StatusOK
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}
		r.ExitAll()
	}

	err = commonAppService.TokenService().Validate(ctx, customClaims)
	if err != nil {
		if errors.Is(err, commonAppService.ErrTokenExpired) {
			r.Response.Status = http.StatusOK
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "token not active",
			})
		} else if errors.Is(err, commonAppService.ErrTokenNotActive) {
			r.Response.Status = http.StatusOK
			r.Response.WriteJson(g.Map{
				"code":    http.StatusForbidden,
				"message": "no permission",
			})
		}
		r.ExitAll()
	}

	commonAppService.ContextService().Init(r, &commonModel.ContextUser{
		UserID:       customClaims.CustomData["user_id"].(string),
		UserName:     customClaims.CustomData["user_name"].(string),
		UserNickname: customClaims.CustomData["user_nickname"].(string),
		OrgID:        customClaims.CustomData["org_id"].(string),
		RoleIDs:      customClaims.CustomData["role_ids"].([]int64),
	})
	r.Middleware.Next()
}
