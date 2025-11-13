package system

import (
	"IdentifyService/internal/app/system/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type ListUserLogReq struct {
	g.Meta `path:"/logs/user" tags:"日志管理" method:"get" summary:"日志列表"`
	model.PageReq
	model.Author
}

type ListUserLogRes struct {
	g.Meta `mime:"application/json"`

	List []*LogItem `json:"list"`
	model.PageRes
}

type LogItem struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`

	OperatorID string    `json:"operator_id"`
	IP         string    `json:"ip"`
	CreateTime time.Time `json:"create_time"`
}
