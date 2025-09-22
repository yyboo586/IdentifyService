package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionPoint is the golang structure for table t_permission.
type PermissionPoint struct {
	ID          int64       `orm:"id"`          // 权限点ID
	Code        string      `orm:"code"`        // 权限点代码
	CodeName    string      `orm:"code_name"`   // 权限点名称
	ResourceID  int64       `orm:"resource_id"` // 资源ID
	Description string      `orm:"description"` // 权限点描述
	CreatedAt   *gtime.Time `orm:"created_at"`  // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"`  // 更新时间
}
