package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 系统内置角色名字常量
const (
	RoleSuperAdmin  = "超级管理员"
	RoleOrgMatainer = "前台组织所有者" // 不一定是后台管理员
	RoleUser        = "普通用户"
)

type RoleStatus int

const (
	_                  RoleStatus = iota
	RoleStatusEnabled             // 启用
	RoleStatusDisabled            // 禁用
)

type Role struct {
	ID        int64       `json:"id" dc:"角色ID"`
	PID       int64       `json:"pid" dc:"父级ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	Name      string      `json:"name" dc:"角色名称"`
	Status    RoleStatus  `json:"status" dc:"状态"`
	CreatorID string      `json:"creator_id" dc:"创建人ID"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`

	MenuIDs []int64 `json:"menu_ids" dc:"资源ID列表"`
}
