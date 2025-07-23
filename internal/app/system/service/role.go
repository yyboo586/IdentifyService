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
		EditStatus(ctx context.Context, id int64, enabled bool) (err error)

		// 根据角色ID获取角色详细信息
		Get(ctx context.Context, id int64) (out *model.Role, err error)
		// 根据组织ID获取角色列表
		ListByOrgID(ctx context.Context, orgID string) (out []*model.Role, err error)

		// 根据角色ID获取角色树形结构
		GetTreeByRoleID(ctx context.Context, id int64) (out *system.RoleNode, err error)
		// 根据组织ID获取角色列表(树形结构)
		ListTreesByOrgID(ctx context.Context, orgID string) (out []*system.RoleNode, err error)

		// 根据用户ID获取拥有的角色ID列表
		GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error)
		// 根据用户ID获取拥有的角色列表
		GetRolesByUserID(ctx context.Context, userID string) (out []*model.Role, err error)
		// 过滤角色ID，确保只能分配自己有权限的角色ID
		FilterRoleIDs(ctx context.Context, roleIds []int64, userID string, includeChildren bool) (out []int64, err error)
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
