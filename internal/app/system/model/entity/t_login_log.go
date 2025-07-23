package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TLoginLog is the golang structure for table t_login_log.
type TLoginLog struct {
	ID        int64       `orm:"id"`
	OrgID     string      `orm:"org_id"`
	LoginName string      `orm:"login_name"`
	IP        string      `orm:"ip"`
	Browser   string      `orm:"browser"`
	Status    int         `orm:"status"`
	Msg       string      `orm:"msg"`
	LoginTime *gtime.Time `orm:"login_time"`
	CreatedAt *gtime.Time `orm:"created_at"`
}
