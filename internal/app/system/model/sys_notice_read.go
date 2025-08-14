package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysNoticeReadInfoRes is the golang structure for table sys_notice_read.
type SysNoticeReadInfoRes struct {
	gmeta.Meta `orm:"table:sys_notice_read"`
	ID         int64       `orm:"id,primary" json:"id" dc:"id"`          // id
	NoticeId   int64       `orm:"notice_id" json:"noticeId" dc:"信息id"`   // 信息id
	UserId     int64       `orm:"user_id" json:"userId" dc:"用户id"`       // 用户id
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"更新时间"` // 更新时间
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"阅读时间"` // 阅读时间
}

type SysNoticeReadListRes struct {
	ID        int64       `json:"id" dc:"id"`
	NoticeId  int64       `json:"noticeId" dc:"信息id"`
	UserId    int64       `json:"userId" dc:"用户id"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"阅读时间"`
}

// SysNoticeReadSearchReq 分页请求参数
type SysNoticeReadSearchReq struct {
	PageReq
}

// SysNoticeReadSearchRes 列表返回结果
type SysNoticeReadSearchRes struct {
	PageRes
	List []*SysNoticeReadListRes `json:"list"`
}

// SysNoticeReadAddReq 添加操作请求参数
type SysNoticeReadAddReq struct {
	NoticeId  int64       `json:"noticeId" dc:"信息id"`
	UserId    string      `json:"userId" dc:"用户id"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"阅读时间"`
}

// SysNoticeReadAddReq 添加操作请求参数
type SysNoticeReadNoticeReq struct {
	NoticeId  int64 `p:"noticeId" dc:"信息id"`
	UserId    string
	CreatedAt *gtime.Time
}

// SysNoticeReadEditReq 修改操作请求参数
type SysNoticeReadEditReq struct {
	ID int64 `p:"id" v:"required#主键ID不能为空" dc:"id"`
}
