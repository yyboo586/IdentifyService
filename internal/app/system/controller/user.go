package controller

import (
	"context"

	"IdentifyService/api/v1/system"

	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"

	"github.com/gogf/gf/v2/util/gconv"
)

var (
	UserController = userController{}
)

type userController struct {
}

func (c *userController) Create(ctx context.Context, req *system.UserCreateReq) (res *system.UserCreateRes, err error) {
	_, err = service.User().Create(ctx, req)
	return
}

func (c *userController) Register(ctx context.Context, req *system.UserRegisterReq) (res *system.UserRegisterRes, err error) {
	_, err = service.User().Register(ctx, req)
	return
}

func (c *userController) Delete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.IDs)
	return
}

func (c *userController) EditPersonalInfo(ctx context.Context, req *system.EditUserPersonalInfoReq) (res *system.EditUserPersonalInfoRes, err error) {
	err = service.User().EditPersonalInfo(ctx, req)
	return
}

func (c *userController) EditUserPermission(ctx context.Context, req *system.EditUserPermissionReq) (res *system.EditUserPermissionRes, err error) {
	err = service.User().EditUserPermission(ctx, req)
	return
}

func (c *userController) ResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error) {
	err = service.User().ResetUserPwd(ctx, req)
	return
}

func (c *userController) GetUserInfo(ctx context.Context, req *system.GetUserInfoReq) (res *system.GetUserInfoRes, err error) {
	user, err := service.User().GetUserInfoByID(ctx, req.ID)
	if err != nil {
		return
	}
	res.User = c.format(user)
	return
}

func (c *userController) List(ctx context.Context, req *system.UserListReq) (res *system.UserListRes, err error) {
	total, out, err := service.User().List(ctx, req)
	if err != nil {
		return
	}

	res = new(system.UserListRes)
	res.Total = gconv.Int(total)
	for _, v := range out {
		res.List = append(res.List, c.format(v))
	}
	return
}

func (c *userController) format(in *model.User) (out *system.User) {
	out = &system.User{
		ID:        in.ID,
		Name:      in.Name,
		Nickname:  in.Nickname,
		Mobile:    in.Mobile,
		Email:     in.Email,
		Enabled:   in.Status == model.UserStatusEnabled,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
