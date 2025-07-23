package service

import (
	"context"

	"IdentifyService/internal/app/system/model"

	"IdentifyService/api/v1/system"
)

type (
	IRole interface {
		Add(ctx context.Context, req *system.RoleAddReq) (id int64, err error)

		DeleteByIDs(ctx context.Context, ids []int64) (err error)

		Edit(ctx context.Context, req *system.RoleEditReq) (err error)

		// 根据角色ID获取角色详细信息
		Get(ctx context.Context, id int64) (out *model.Role, err error)
		// 根据组织ID获取角色列表
		ListByOrgID(ctx context.Context, orgID string) (out []*model.Role, err error)

		// 根据用户ID获取拥有的角色ID列表
		GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error)

		// 根据角色名称获取角色信息
		GetByName(ctx context.Context, roleName string) (res *model.Role, err error)

		GetFilteredNamedPolicy(ctx context.Context, id int64) (gpSlice []int64, err error)

		FindSonByParentId(roleList []*model.Role, id int64) []*model.Role
		FindSonIDsByParentID(roleList []*model.Role, id int64) []int64
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
