package model

import "github.com/gogf/gf/v2/os/gtime"

type LoginStatus int

const (
	_                  LoginStatus = iota
	LoginStatusSuccess             //  登录成功
	LoginStatusFailed              // 登录失败
)

type LoginLog struct {
	ID        int64       `json:"id" dc:"日志ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	LoginName string      `json:"login_name" dc:"登录名"`
	IP        string      `json:"ip" dc:"IP地址"`
	Browser   string      `json:"browser" dc:"浏览器"`
	Success   bool        `json:"success" dc:"是否成功"`
	Message   string      `json:"message" dc:"消息"`
	LoginTime *gtime.Time `json:"login_time" dc:"登录时间"`
	CreatedAt *gtime.Time `json:"created_at" dc:"记录创建时间"`
}

type OperLog struct {
	ID         uint64      `json:"id" dc:"日志ID"`
	OrgID      string      `json:"org_id" dc:"组织ID"`
	OperName   string      `json:"oper_name" dc:"操作人员"`
	OperUrl    string      `json:"oper_url" dc:"操作URL"`
	OperMethod string      `json:"oper_method" dc:"操作方法"`
	OperIP     string      `json:"oper_ip" dc:"操作IP"`
	OperTime   *gtime.Time `json:"oper_time" dc:"操作时间"`
	CreatedAt  *gtime.Time `json:"created_at" dc:"记录创建时间"`
}
