package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysNoticeRead is the golang structure for table sys_notice_read.
type SysNoticeRead struct {
	gmeta.Meta `orm:"table:sys_notice_read"`
	ID         int64       `orm:"id,primary" json:"id"`        // id
	NoticeId   int64       `orm:"notice_id" json:"noticeId"`   // 信息id
	UserId     int64       `orm:"user_id" json:"userId"`       // 用户id
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 阅读时间
}
