package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Resource is the golang structure for table t_resource.
type Resource struct {
	ID        int64       `orm:"id"`
	Type      int         `orm:"type"`
	Code      string      `orm:"code"`
	CreatedAt *gtime.Time `orm:"created_at"`
	UpdatedAt *gtime.Time `orm:"updated_at"`
}
