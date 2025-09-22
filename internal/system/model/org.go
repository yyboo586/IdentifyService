package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

const (
	DefaultOrgID = "00000000-0000-0000-0000-000000000000"
)

type OrgStatus int64

const (
	_                 OrgStatus = iota
	OrgStatusEnabled            // 启用
	OrgStatusDisabled           // 禁用
)

func IsValidOrgStatus(status OrgStatus) bool {
	return status == OrgStatusEnabled || status == OrgStatusDisabled
}

type Org struct {
	ID          string      `json:"id" dc:"组织ID"`
	PID         string      `json:"pid" dc:"父级ID"`
	Name        string      `json:"name" dc:"组织名称"`
	ManagerID   string      `json:"manager_id" dc:"负责人ID"`
	ManagerName string      `json:"manager_name" dc:"负责人名称"`
	Status      OrgStatus   `json:"status" dc:"组织状态"`
	CreatedAt   *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at" dc:"修改时间"`
}

type OrgConfig struct {
	OrgID string `json:"org_id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 前端资源（组织管理）
const (
	DirOrgManager  = "dir:org:manage"
	MenuOrgManager = "menu:org:list"

	ButtonOrgCreate = "button:org:create"
	ButtonOrgDelete = "button:org:delete"
	ButtonOrgEdit   = "button:org:edit"
	ButtonOrgView   = "button:org:view"
)

// API 资源（组织管理）
const (
	APIOrgCreate = "api:org:create"
	APIOrgDelete = "api:org:delete"
	APIOrgEdit   = "api:org:edit"
	APIOrgView   = "api:org:view"
	APIOrgList   = "api:org:list"
)

// 组织管理权限点常量
const (
	OrgView   PermissionPoint = "org:view"
	OrgCreate PermissionPoint = "org:create"
	OrgEdit   PermissionPoint = "org:edit"
	OrgDelete PermissionPoint = "org:delete"
)

// 组织管理权限点清单
var OrgPermissions []*Permission = []*Permission{
	{
		Code:     OrgCreate,
		CodeName: "创建组织",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirOrgManager},
			{Type: ResourceTypeMenu, Code: MenuOrgManager},
			{Type: ResourceTypeAPI, Code: APIOrgList},
			{Type: ResourceTypeButton, Code: ButtonOrgCreate},
			{Type: ResourceTypeAPI, Code: APIOrgCreate},
		},
	},
	{
		Code:     OrgDelete,
		CodeName: "删除组织",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirOrgManager},
			{Type: ResourceTypeMenu, Code: MenuOrgManager},
			{Type: ResourceTypeAPI, Code: APIOrgList},
			{Type: ResourceTypeButton, Code: ButtonOrgDelete},
			{Type: ResourceTypeAPI, Code: APIOrgDelete},
		},
	},
	{
		Code:     OrgEdit,
		CodeName: "编辑组织",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirOrgManager},
			{Type: ResourceTypeMenu, Code: MenuOrgManager},
			{Type: ResourceTypeAPI, Code: APIOrgList},
			{Type: ResourceTypeButton, Code: ButtonOrgEdit},
			{Type: ResourceTypeAPI, Code: APIOrgEdit},
		},
	},
	{
		Code:     OrgView,
		CodeName: "组织详情",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirOrgManager},
			{Type: ResourceTypeMenu, Code: MenuOrgManager},
			{Type: ResourceTypeAPI, Code: APIOrgList},
			{Type: ResourceTypeButton, Code: ButtonOrgView},
			{Type: ResourceTypeAPI, Code: APIOrgView},
		},
	},
}
