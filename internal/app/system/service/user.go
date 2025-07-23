package service

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
)

type (
	IUser interface {
		// 创建用户
		Create(ctx context.Context, req *system.UserCreateReq) (id string, err error)
		// 前台用户自注册
		Register(ctx context.Context, req *system.UserRegisterReq) (userID string, err error)

		// 删除用户
		Delete(ctx context.Context, ids []string) (err error)

		// 修改个人资料
		EditPersonalInfo(ctx context.Context, req *system.EditUserPersonalInfoReq) (err error)
		// 修改用户角色
		EditUserRoles(ctx context.Context, req *system.EditUserRolesReq) (err error)
		// 修改用户状态
		EditUserStatus(ctx context.Context, req *system.EditUserStatusReq) (err error)
		// 重置用户密码
		ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)

		// 通过id获取用户信息
		GetByID(ctx context.Context, id string) (user *model.User, err error)
		// 通过用户名获取用户信息
		GetByUsername(ctx context.Context, username string) (user *model.User, err error)
		// 获取用户列表
		List(ctx context.Context, req *system.UserListReq) (total interface{}, out []*model.User, err error)

		// 判断用户是否是超级管理员
		IsSuperAdmin(ctx context.Context, userID string) bool
		// 验证用户名和密码
		ValidateUsernameAndPassword(ctx context.Context, hashPassword, salt, password string) (err error)
	}
)

var (
	localSysUser IUser
)

func User() IUser {
	if localSysUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localSysUser
}

func RegisterUser(i IUser) {
	localSysUser = i
}
