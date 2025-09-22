package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 系统内置角色名字常量
const (
	DefaultRoleSuperAdmin = "超级管理员"
	DefaultRoleOrgAdmin   = "组织管理员"
	DefaultRoleNormalUser = "普通用户"
)

// 系统内置角色及权限分配
var DefaultRoles = []*Role{
	{
		OrgID:     DefaultOrgID,
		Name:      DefaultRoleSuperAdmin,
		CreatorID: DefaultSuperAdminID,
	},
	{
		OrgID:     DefaultOrgID,
		Name:      DefaultRoleOrgAdmin,
		CreatorID: DefaultSuperAdminID,
		PermissionPoints: []PermissionPointCode{
			RoleCreate,
			RoleEdit,
			RoleDelete,
			RoleView,
			RoleList,
			RoleAssign,
			UserCreate,
			UserEdit,
			UserDelete,
			UserView,
			UserList,
		},
	},
	{
		OrgID:     DefaultOrgID,
		Name:      DefaultRoleNormalUser,
		CreatorID: DefaultSuperAdminID,
		PermissionPoints: []PermissionPointCode{
			UserView,
		},
	},
}

type Role struct {
	ID        int64       `json:"id" dc:"角色ID"`
	PID       int64       `json:"pid" dc:"父级ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	Name      string      `json:"name" dc:"角色名称"`
	CreatorID string      `json:"creator_id" dc:"创建人ID"`
	DeletorID string      `json:"deletor_id" dc:"删除人ID"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" dc:"删除时间"`

	PermissionPoints []PermissionPointCode `json:"permission_points" dc:"资源ID列表"`
}

// 前端资源（角色管理）
const (
	DirRoleManager  = "dir:role:manage"
	MenuRoleManager = "menu:role:list"

	ButtonRoleCreate = "button:role:create"
	ButtonRoleDelete = "button:role:delete"
	ButtonRoleEdit   = "button:role:edit"
	ButtonRoleView   = "button:role:view"
)

// API 资源（角色管理）
const (
	APIRoleCreate = "api:role:create"
	APIRoleDelete = "api:role:delete"
	APIRoleEdit   = "api:role:edit"
	APIRoleView   = "api:role:view"
	APIRoleList   = "api:role:list"
)

// 角色管理权限点常量
const (
	RoleCreate PermissionPointCode = "role:create"
	RoleEdit   PermissionPointCode = "role:edit"
	RoleDelete PermissionPointCode = "role:delete"
	RoleView   PermissionPointCode = "role:view"
	RoleList   PermissionPointCode = "role:list"

	RoleAssign PermissionPointCode = "role:assign"
)

var RoleResources []*Resource = []*Resource{
	{Type: ResourceTypeDir, Code: DirRoleManager},
	{Type: ResourceTypeMenu, Code: MenuRoleManager},
	{Type: ResourceTypeButton, Code: ButtonRoleCreate},
	{Type: ResourceTypeButton, Code: ButtonRoleDelete},
	{Type: ResourceTypeButton, Code: ButtonRoleEdit},
	{Type: ResourceTypeButton, Code: ButtonRoleView},
	{Type: ResourceTypeAPI, Code: APIRoleCreate},
	{Type: ResourceTypeAPI, Code: APIRoleDelete},
	{Type: ResourceTypeAPI, Code: APIRoleEdit},
	{Type: ResourceTypeAPI, Code: APIRoleView},
	{Type: ResourceTypeAPI, Code: APIRoleList},
}

// 角色管理权限点清单
var RolePermissions []*PermissionPoint = []*PermissionPoint{
	{
		Code:     RoleCreate,
		CodeName: "创建角色",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirRoleManager},
			{Type: ResourceTypeMenu, Code: MenuRoleManager},
			{Type: ResourceTypeAPI, Code: APIRoleList},
			{Type: ResourceTypeButton, Code: ButtonRoleCreate},
			{Type: ResourceTypeAPI, Code: APIRoleCreate},
		},
	},
	{
		Code:     RoleDelete,
		CodeName: "删除角色",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirRoleManager},
			{Type: ResourceTypeMenu, Code: MenuRoleManager},
			{Type: ResourceTypeAPI, Code: APIRoleList},
			{Type: ResourceTypeButton, Code: ButtonRoleDelete},
			{Type: ResourceTypeAPI, Code: APIRoleDelete},
		},
	},
	{
		Code:     RoleEdit,
		CodeName: "编辑角色",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirRoleManager},
			{Type: ResourceTypeMenu, Code: MenuRoleManager},
			{Type: ResourceTypeAPI, Code: APIRoleList},
			{Type: ResourceTypeButton, Code: ButtonRoleEdit},
			{Type: ResourceTypeAPI, Code: APIRoleEdit},
		},
	},
	{
		Code:     RoleView,
		CodeName: "角色详情",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirRoleManager},
			{Type: ResourceTypeMenu, Code: MenuRoleManager},
			{Type: ResourceTypeAPI, Code: APIRoleList},
			{Type: ResourceTypeButton, Code: ButtonRoleView},
			{Type: ResourceTypeAPI, Code: APIRoleView},
		},
	},
	{
		Code:     RoleList,
		CodeName: "角色列表",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirRoleManager},
			{Type: ResourceTypeMenu, Code: MenuRoleManager},
			{Type: ResourceTypeAPI, Code: APIRoleList},
		},
	},
	{
		Code:     RoleAssign,
		CodeName: "角色分配",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirRoleManager},
			{Type: ResourceTypeMenu, Code: MenuRoleManager},
			{Type: ResourceTypeAPI, Code: APIRoleList},
			{Type: ResourceTypeButton, Code: ButtonRoleEdit},
			{Type: ResourceTypeAPI, Code: APIRoleEdit},
		},
	},
}
