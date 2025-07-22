package system

import (
	v1 "IdentifyService/api/v1"
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysOperLogSearchReq 分页请求参数
type SysOperLogSearchReq struct {
	g.Meta        `path:"/operLog/list" tags:"系统后台/操作日志" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"`         //系统模块
	RequestMethod string `p:"requestMethod"` //请求方式
	OperName      string `p:"operName"`      //操作人员
	v1.PageReq
	v1.Author
}

// SysOperLogSearchRes 列表返回结果
type SysOperLogSearchRes struct {
	g.Meta `mime:"application/json"`
	v1.ListRes
	List []*model.SysOperLogListRes `json:"list"`
}

// SysOperLogGetReq 获取一条数据请求
type SysOperLogGetReq struct {
	g.Meta `path:"/operLog/get" tags:"系统后台/操作日志" method:"get" summary:"获取操作日志信息"`
	v1.Author
	OperId uint64 `p:"operId" v:"required#主键必须"` //通过主键获取
}

// SysOperLogGetRes 获取一条数据结果
type SysOperLogGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysOperLogInfoRes
}

// SysOperLogDeleteReq 删除数据请求
type SysOperLogDeleteReq struct {
	g.Meta `path:"/operLog/delete" tags:"系统后台/操作日志" method:"delete" summary:"删除操作日志"`
	v1.Author
	OperIds []uint64 `p:"operIds" v:"required#主键必须"` //通过主键删除
}

// SysOperLogDeleteRes 删除数据返回
type SysOperLogDeleteRes struct {
	v1.EmptyRes
}

type SysOperLogClearReq struct {
	g.Meta `path:"/operLog/clear" tags:"系统后台/操作日志" method:"delete" summary:"清除日志"`
	v1.Author
}

type SysOperLogClearRes struct {
	v1.EmptyRes
}
