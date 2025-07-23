package service

import (
	"context"

	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IContext interface {
	Init(r *ghttp.Request, customCtx *model.Context)
	Set(ctx context.Context, customCtx *model.ContextUser)
	Get(ctx context.Context) *model.Context
	GetUser(ctx context.Context) *model.ContextUser
}

var localContext IContext

func ContextService() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContextService() {
	localContext = &sContext{}
}

type sContext struct {
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.CtxKey, customCtx)
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) Set(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) (out *model.Context) {
	value := ctx.Value(consts.CtxKey)
	if value == nil {
		g.Log().Info(ctx, "value is nil")
		return nil
	}

	out, ok := value.(*model.Context)
	if !ok {
		return nil
	}
	return
}

func (s *sContext) GetUser(ctx context.Context) *model.ContextUser {
	return s.Get(ctx).User
}
