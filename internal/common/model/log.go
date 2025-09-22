package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type LogType int

const (
	_           LogType = iota
	LogTypeAuth         // 认证
	LogTypeOper         // 操作
)

type ListLogReq struct {
	OrgID     string      `json:"org_id" dc:"组织ID"`
	StartTime *gtime.Time `json:"start_time" dc:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" dc:"结束时间"`
	PageReq
}

type ListLogRes struct {
	List []*Log
	PageRes
}

type Log struct {
	ID        int64       `json:"id" dc:"日志ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	UserID    string      `json:"user_id" dc:"用户ID"`
	UserName  string      `json:"user_name" dc:"用户名"`
	IP        string      `json:"ip" dc:"IP地址"`
	Type      LogType     `json:"type" dc:"日志类型"`
	Content   interface{} `json:"content" dc:"日志内容"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
}
