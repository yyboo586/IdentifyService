package main

import (
	commonDomain "IdentifyService/internal/common/domain"
	commonModel "IdentifyService/internal/common/model"
	commonAppService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/domain"
	"IdentifyService/internal/system/service"
	"context"
	"errors"
	"mime"
	"net/http"
	"slices"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	ctx := context.Background()
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

	casbin, err := domain.NewCasbinEnforcer(ctx)
	if err != nil {
		panic(err)
	}

	domainRole := domain.NewRole(casbin)
	domainUser := domain.NewUser(casbin)
	domainOrg := domain.NewOrg(domainUser)
	domainPermission := domain.NewPermission(casbin)

	// 初始化权限点数据
	err = domainPermission.Init(context.Background())
	if err != nil {
		panic(err)
	}

	authAppService := service.NewAuthService(domainRole, domainUser, domainOrg, domainPermission)
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

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

var (
	// streamContentType is the content types for stream response.
	streamContentType = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}
)

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0 {
		return
	}

	// It does not output common response content if it is stream response.
	mediaType, _, _ := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))
	for _, ct := range streamContentType {
		if mediaType == ct {
			return
		}
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates an error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
		msg = code.Message()
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}

func InjectCtx(r *ghttp.Request) {
	ctx := r.GetCtx()

	excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
	if slices.Contains(excludePaths, r.URL.Path) {
		r.Middleware.Next()
		return
	}

	tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	customClaims, err := commonAppService.TokenService().Parse(ctx, tokenStr)
	if err != nil {
		g.Log().Error(ctx, err)
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
		g.Log().Error(ctx, err)
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

	roleIDs := make([]int64, 0, len(customClaims.CustomData["role_ids"].([]interface{})))
	for _, roleID := range customClaims.CustomData["role_ids"].([]interface{}) {
		roleIDs = append(roleIDs, int64(roleID.(float64)))
	}
	commonAppService.ContextService().Init(r, &commonModel.ContextUser{
		UserID:       customClaims.CustomData["user_id"].(string),
		UserName:     customClaims.CustomData["user_name"].(string),
		UserNickname: customClaims.CustomData["user_nickname"].(string),
		OrgID:        customClaims.CustomData["org_id"].(string),
		RoleIDs:      roleIDs,
	})
	r.Middleware.Next()
}
