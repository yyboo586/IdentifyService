package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type OrgStatus int64

const (
	_                 OrgStatus = iota
	OrgStatusEnabled            // 启用
	OrgStatusDisabled           // 禁用
)

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
