package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Org is the golang structure for table t_org.
type Org struct {
	ID          string      `orm:"id"`
	PID         string      `orm:"pid"`
	Name        string      `orm:"name"`
	ManagerID   string      `orm:"manager_id"`
	ManagerName string      `orm:"manager_name"`
	Status      int         `orm:"status"`
	CreatedAt   *gtime.Time `orm:"created_at"`
	UpdatedAt   *gtime.Time `orm:"updated_at"`
}
