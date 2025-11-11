package system

import (
	commonApi "IdentifyService/api/v1/common"
	"IdentifyService/internal/app/common/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentSearchReq 分页请求参数
type SysAttachmentSearchReq struct {
	g.Meta `path:"/list" tags:"附件管理" method:"get" summary:"附件管理列表"`
	commonApi.Author
	model.SysAttachmentSearchReq
}

// SysAttachmentSearchRes 列表返回结果
type SysAttachmentSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysAttachmentSearchRes
}

// SysAttachmentAddReq 添加操作请求参数
type SysAttachmentAddReq struct {
	g.Meta `path:"/add" tags:"附件管理" method:"post" summary:"附件管理添加"`
	commonApi.Author
	*model.SysAttachmentAddReq
}

// SysAttachmentAddRes 添加操作返回结果
type SysAttachmentAddRes struct {
	commonApi.EmptyRes
}

// SysAttachmentEditReq 修改操作请求参数
type SysAttachmentEditReq struct {
	g.Meta `path:"/edit" tags:"附件管理" method:"put" summary:"附件管理修改"`
	commonApi.Author
	*model.SysAttachmentEditReq
}

// SysAttachmentEditRes 修改操作返回结果
type SysAttachmentEditRes struct {
	commonApi.EmptyRes
}

// SysAttachmentGetReq 获取一条数据请求
type SysAttachmentGetReq struct {
	g.Meta `path:"/get" tags:"附件管理" method:"get" summary:"获取附件管理信息"`
	commonApi.Author
	Id int64 `p:"id" v:"required#主键必须"` //通过主键获取
}

// SysAttachmentGetRes 获取一条数据结果
type SysAttachmentGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysAttachmentInfoRes
}

// SysAttachmentDeleteReq 删除数据请求
type SysAttachmentDeleteReq struct {
	g.Meta `path:"/delete" tags:"附件管理" method:"delete" summary:"删除附件管理"`
	commonApi.Author
	Ids []int64 `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SysAttachmentDeleteRes 删除数据返回
type SysAttachmentDeleteRes struct {
	commonApi.EmptyRes
}

// 附件管理状态修改（状态）
type SysAttachmentStatusSwitchReq struct {
	g.Meta `path:"/changeStatus" tags:"附件管理" method:"put" summary:"修改状态"`
	commonApi.Author
	Id     int64 `p:"id" v:"required#主键必须"`     //通过主键修改
	Status bool  `p:"status" v:"required#状态必须"` //通过主键获取
}
type SysAttachmentStatusSwitchRes struct {
	commonApi.EmptyRes
}
