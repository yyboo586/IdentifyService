package interfaces

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IUser interface {
		// 创建用户
		Create(ctx context.Context, tx gdb.TX, in *model.User) (err error)
		// 删除用户
		Delete(ctx context.Context, tx gdb.TX, ids []string) (err error)

		// 修改个人资料
		Edit(ctx context.Context, req *system.UserEditReq) (err error)
		// 重置用户密码
		ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)

		// 通过id获取用户信息
		GetByID(ctx context.Context, id string) (user *model.User, err error)
		// 通过用户名获取用户信息
		GetByUsername(ctx context.Context, username string) (user *model.User, err error)
		// 获取用户列表
		Search(ctx context.Context, req *system.UserSearchReq) (total interface{}, out []*model.User, err error)

		// 验证用户名和密码
		ValidateUsernameAndPassword(ctx context.Context, hashPassword, salt, password string) (err error)

		// 检查账户状态
		IsEnabled(ctx context.Context, in *model.User) (err error)

		// 检查组织下属是否还有用户
		IsOrgHasUsers(ctx context.Context, orgID string) (hasUsers bool, err error)
	}
)
