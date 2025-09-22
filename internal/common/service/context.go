package service

import (
	"context"
	"errors"

	"IdentifyService/internal/common/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IContext interface {
	Init(r *ghttp.Request, info *model.ContextUser)
	SetUserInfo(r *ghttp.Request, info *model.ContextUser)
	GetUserID(ctx context.Context) (string, error)
	GetUserName(ctx context.Context) (string, error)
	GetUserNickname(ctx context.Context) (string, error)
	GetOrgID(ctx context.Context) (string, error)
	GetRoleIDs(ctx context.Context) ([]int64, error)
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
func (s *sContext) Init(r *ghttp.Request, info *model.ContextUser) {
	r.SetCtxVar(model.CtxKey, info)
}

func (s *sContext) SetUserInfo(r *ghttp.Request, info *model.ContextUser) {
	r.SetCtxVar(model.CtxKey, info)
}

func (s *sContext) IsLogin(ctx context.Context) bool {
	value := ctx.Value(model.CtxKey)
	return value != nil
}

func (s *sContext) getUserInfo(ctx context.Context) (*model.ContextUser, error) {
	value := ctx.Value(model.CtxKey)
	if value == nil {
		return nil, errors.New("user info not found")
	}
	return value.(*model.ContextUser), nil
}

func (s *sContext) GetUserID(ctx context.Context) (string, error) {
	userInfo, err := s.getUserInfo(ctx)
	if err != nil {
		return "", err
	}
	return userInfo.UserID, nil
}

func (s *sContext) GetUserName(ctx context.Context) (string, error) {
	userInfo, err := s.getUserInfo(ctx)
	if err != nil {
		return "", err
	}
	return userInfo.UserName, nil
}

func (s *sContext) GetUserNickname(ctx context.Context) (string, error) {
	userInfo, err := s.getUserInfo(ctx)
	if err != nil {
		return "", err
	}
	return userInfo.UserNickname, nil
}

func (s *sContext) GetOrgID(ctx context.Context) (string, error) {
	userInfo, err := s.getUserInfo(ctx)
	if err != nil {
		return "", err
	}
	return userInfo.OrgID, nil
}

func (s *sContext) GetRoleIDs(ctx context.Context) ([]int64, error) {
	userInfo, err := s.getUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return userInfo.RoleIDs, nil
}
