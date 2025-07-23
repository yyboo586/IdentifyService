package service

import (
	"context"

	"IdentifyService/internal/app/system/model"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model/entity"
)

type (
	ISysRole interface {
		GetRoleListSearch(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
		GetRoleList(ctx context.Context) (list []*entity.SysRole, err error)
		AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error)
		DelRoleRule(ctx context.Context, roleId int64) (err error)
		AddRole(ctx context.Context, req *system.RoleAddReq) (err error)
		Get(ctx context.Context, id uint) (res *model.RoleInfoRes, err error)
		Get2(ctx context.Context, roleID uint) (res *model.RoleInfoRes, err error)
		GetByName(ctx context.Context, roleName string) (res *model.RoleInfoRes, err error)
		GetFilteredNamedPolicy(ctx context.Context, id uint) (gpSlice []int, err error)
		EditRole(ctx context.Context, req *system.RoleEditReq) (err error)
		DeleteByIds(ctx context.Context, ids []int64) (err error)
		RoleDeptTreeSelect(ctx context.Context) (res *system.RoleDeptTreeSelectRes, err error)
		RoleDataScope(ctx context.Context, req *system.DataScopeReq) error
		FindSonByParentId(roleList []*entity.SysRole, id uint) []*entity.SysRole
		FindSonIdsByParentId(roleList []*entity.SysRole, id uint) []uint
		GetRoleDataScope(ctx context.Context, roleId uint) (data []*model.ScopeAuthData, err error)
		GetRoleMenuScope(ctx context.Context, roleIds []uint, menuId uint) (data []*model.ScopeAuthData, err error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
