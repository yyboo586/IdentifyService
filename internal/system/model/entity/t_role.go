package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table t_role.
type Role struct {
	ID        int64       `orm:"id"`
	Pid       int64       `orm:"pid"`
	OrgID     string      `orm:"org_id"`
	Name      string      `orm:"name"`
	CreatorID string      `orm:"creator_id"`
	DeletorID string      `orm:"deletor_id"`
	CreatedAt *gtime.Time `orm:"created_at"`
	UpdatedAt *gtime.Time `orm:"updated_at"`
	DeletedAt *gtime.Time `orm:"deleted_at"`
}
