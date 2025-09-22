package system

import (
	"IdentifyService/internal/common/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ListLoginLogReq struct {
	g.Meta `path:"/logs/login-log" tags:"日志管理" method:"get" summary:"登录日志列表"`
	model.Author
	OrgID     string      `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	StartTime *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" dc:"结束时间"`
	model.PageReq
}

type ListLoginLogRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.Log `json:"list"`
	model.PageRes
}

type ListOperLogReq struct {
	g.Meta    `path:"/logs/oper-log" tags:"日志管理" method:"get" summary:"操作日志列表"`
	OrgID     string      `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	StartTime *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" dc:"结束时间"`
	model.PageReq
}

type ListOperLogRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.Log `json:"list"`
	model.PageRes
}
