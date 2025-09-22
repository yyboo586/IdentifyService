package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TLog is the golang structure for table t_log.
type TLog struct {
	ID        int64       `orm:"id"`
	OrgID     string      `orm:"org_id"`
	UserID    string      `orm:"user_id"`
	UserName  string      `orm:"user_name"`
	IP        string      `orm:"ip"`
	Type      int         `orm:"type"`
	Content   string      `orm:"content"`
	CreatedAt *gtime.Time `orm:"created_at"`
}
