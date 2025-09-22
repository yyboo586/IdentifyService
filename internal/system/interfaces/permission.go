package interfaces

import (
	"IdentifyService/internal/system/model"
	"context"
)

type (
	IPermission interface {
		// 用户--角色
		AssignUserRoles(ctx context.Context, userID string, roleIDs []int64) (err error)
		RemoveUserRoles(ctx context.Context, userID string, roleIDs []int64) (err error)
		RemoveUserAllRoles(ctx context.Context, userIDs []string) (err error)

		// 角色--权限
		AssignRolePermissions(ctx context.Context, roleID int64, permissionPointCodes []model.PermissionPointCode) (err error)
		RemoveRolePermissions(ctx context.Context, roleID int64, permissionPointCodes []model.PermissionPointCode) (err error)
		RemoveRoleAllPermissions(ctx context.Context, roleIDs []int64) (err error)
		GetPermissionPointsByRoleID(ctx context.Context, roleIDs []int64) (out []*model.PermissionPoint, err error)

		// 权限检查
		IsSuperAdmin(ctx context.Context, userID string) bool
		IsOrgAdmin(ctx context.Context, userID string, orgInfo *model.Org) bool
		FilterRoleIDsByUserID(ctx context.Context, roleIDs []int64, userID string) (out []int64, err error)
		FilterPermissionPointsByUserID(ctx context.Context, permissionPoints []model.PermissionPointCode, userID string) (out []model.PermissionPointCode, err error)
		// FilterPermissionPointIDsByRoleID(ctx context.Context, permissionPointIDs []int64, roleID int64) (out []int64, err error)
		HasPermission(ctx context.Context, permission model.PermissionPointCode, userID string, orgInfo *model.Org) (hasPermission bool, err error)
		GetPermissionPointsByUserID(ctx context.Context, userID string) (out []*model.PermissionPoint, err error)
		GetPermissionPointCodesByUserID(ctx context.Context, userID string) (out []model.PermissionPointCode, err error)
		GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error)

		// 初始化权限点
		Init(ctx context.Context) (err error)
	}
)
