package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysNoticeSearchReq 分页请求参数
type SysNoticeSearchReq struct {
	g.Meta `path:"/list" tags:"系统后台/通知私信" method:"get" summary:"通知公告列表(管理)"`
	model.Author
	model.SysNoticeSearchReq
}

// SysNoticeSearchRes 列表返回结果
type SysNoticeSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeSearchRes
}

// SysNoticeSearchReq 分页请求参数
type SysNoticeShowSearchReq struct {
	g.Meta `path:"/listShow" tags:"系统后台/通知私信" method:"get" summary:"通知公告列表(浏览)"`
	model.Author
	model.SysNoticeSearchReq
}

// SysNoticeSearchRes 列表返回结果
type SysNoticeShowSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeSearchRes
}

// SysNoticeAddReq 添加操作请求参数
type SysNoticeAddReq struct {
	g.Meta `path:"/add" tags:"系统后台/通知私信" method:"post" summary:"通知公告添加"`
	model.Author
	*model.SysNoticeAddReq
}

// SysNoticeAddRes 添加操作返回结果
type SysNoticeAddRes struct {
	model.EmptyRes
}

// SysNoticeEditReq 修改操作请求参数
type SysNoticeEditReq struct {
	g.Meta `path:"/edit" tags:"系统后台/通知私信" method:"put" summary:"通知公告修改"`
	model.Author
	*model.SysNoticeEditReq
}

// SysNoticeEditRes 修改操作返回结果
type SysNoticeEditRes struct {
	model.EmptyRes
}

// SysNoticeGetReq 获取一条数据请求
type SysNoticeGetReq struct {
	g.Meta `path:"/get" tags:"系统后台/通知私信" method:"get" summary:"获取通知公告信息"`
	model.Author
	ID int64 `p:"id" v:"required#主键必须"` //通过主键获取
}

// SysNoticeGetRes 获取一条数据结果
type SysNoticeGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeInfoRes
}

// SysNoticeDeleteReq 删除数据请求
type SysNoticeDeleteReq struct {
	g.Meta `path:"/delete" tags:"系统后台/通知私信" method:"delete" summary:"删除通知公告"`
	model.Author
	Ids []int64 `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SysNoticeDeleteRes 删除数据返回
type SysNoticeDeleteRes struct {
	model.EmptyRes
}

type SysNoticeIndexDataReq struct {
	g.Meta `path:"getIndexData" tags:"系统后台/通知私信" method:"get" summary:"获取首页通知预览数据"`
	model.Author
}
type SysNoticeIndexDataRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeIndexRes
}
type SysNoticeUnReadCountReq struct {
	g.Meta `path:"unReadCount" tags:"系统后台/通知私信" method:"get" summary:"获取未读数量"`
	model.Author
}
type SysNoticeUnReadCountRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeUnreadCount
}

// SysNoticeShowSearchReq 通知展示分页请求参数
type SysNoticeUserSearchReq struct {
	g.Meta `path:"/userList" tags:"系统后台/通知私信" method:"get" summary:"获取待指定的用户"`
	model.Author
	UserNickName string `p:"userNickname"`
}

// SysNoticeSearchRes 通知展示列表返回结果
type SysNoticeUserSearchRes struct {
	g.Meta   `mime:"application/json"`
	UserList []*model.SysNoticeUserNickname `json:"userList"`
}

// SysNoticeShowSearchReq 通知展示分页请求参数
type SysNoticeUserReadReq struct {
	g.Meta `path:"/userList" tags:"系统后台/通知私信" method:"get" summary:"获取待指定的用户"`
	model.Author
	string `p:"nickName"`
}

// SysNoticeSearchRes 通知展示列表返回结果
type SysNoticeUserReadRes struct {
	model.EmptyRes
}
