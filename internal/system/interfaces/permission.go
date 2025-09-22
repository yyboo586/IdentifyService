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
		GetRolesByUserID(ctx context.Context, userID string) (out []*model.Role, err error)
		GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error)

		// 角色--权限
		AssignRolePermissions(ctx context.Context, roleID int64, permissionIDs []int64) (err error)
		RemoveRolePermissions(ctx context.Context, roleID int64, permissionIDs []int64) (err error)
		RemoveRoleAllPermissions(ctx context.Context, roleIDs []int64) (err error)

		// 权限检查
		FilterRoleIDsByUserID(ctx context.Context, roleIds []int64, userID string) (out []int64, err error)
		FilterPermissionPointIDsByUserID(ctx context.Context, permissionPointIDs []int64, userID string) (out []int64, err error)
		// FilterPermissionPointIDsByRoleID(ctx context.Context, permissionPointIDs []int64, roleID int64) (out []int64, err error)
		IsSuperAdmin(ctx context.Context, userID string) bool
		IsOrgAdmin(ctx context.Context, userID string, orgInfo *model.Org) bool
		HasPermission(ctx context.Context, permission model.PermissionPoint, userID string, orgInfo *model.Org) (hasPermission bool, err error)

		// 权限点管理相关
		GetAllPermissionPonits(ctx context.Context) (permissions []*model.Permission, err error)
		GetPermissionPointsByUserID(ctx context.Context, userID string) (permissionPointList []*model.Permission, err error)
		GetPermissionPointsByRoleID(ctx context.Context, roleID int64) (permissionPointList []*model.Permission, err error)

		// 注册权限点
		RegisterPermissionPoints(ctx context.Context, in []*model.Permission)
		// 初始化权限点
		Init(ctx context.Context) (err error)
	}
)
