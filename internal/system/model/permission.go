package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 权限点
type PermissionPoint string

// 资源类型
type ResourceType uint

const (
	_                  ResourceType = iota
	ResourceTypeDir                 // 目录
	ResourceTypeMenu                // 菜单
	ResourceTypeButton              // 按钮
	ResourceTypeAPI                 // 接口
)

type Resource struct {
	ID        int64        `json:"id"`         // 资源ID
	Type      ResourceType `json:"type"`       // 资源类型
	Code      string       `json:"code"`       // 资源代码
	CreatedAt *gtime.Time  `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time  `json:"updated_at"` // 更新时间
}

// Permission 权限点模型
type Permission struct {
	ID          int64           `json:"id"`          // 权限点ID
	Code        PermissionPoint `json:"code"`        // 权限点代码
	CodeName    string          `json:"code_name"`   // 权限点名称
	Description string          `json:"description"` // 权限点描述
	Resources   []*Resource     `json:"resource"`    // 资源
	CreatedAt   *gtime.Time     `json:"created_at"`  // 创建时间
	UpdatedAt   *gtime.Time     `json:"updated_at"`  // 更新时间
}
