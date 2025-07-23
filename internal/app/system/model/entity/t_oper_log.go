package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TOperLog is the golang structure for table t_oper_log.
type TOperLog struct {
	ID         uint64      `orm:"id"`
	OrgID      string      `orm:"org_id"`
	OperName   string      `orm:"oper_name"`
	OperUrl    string      `orm:"oper_url"`
	OperMethod string      `orm:"oper_method"`
	OperIP     string      `orm:"oper_ip"`
	OperTime   *gtime.Time `orm:"oper_time"`
	CreatedAt  *gtime.Time `orm:"created_at"`
}
