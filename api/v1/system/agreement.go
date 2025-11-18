package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AgreementCreateReq struct {
	g.Meta `path:"/agreement" tags:"协议管理" method:"post" summary:"创建协议"`
	model.Author

	Name       string `json:"name" v:"required#协议名称不能为空" dc:"协议名称(用户服务协议/隐私保护协议)"`
	Version    string `json:"version" v:"required#版本号不能为空" dc:"版本号(格式: 1.0.0)"`
	Content    string `json:"content" v:"required#协议内容不能为空"`
	PublishNow bool   `json:"publish_now" dc:"是否立即发布"`
}

type AgreementCreateRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id"`
}

type AgreementUpdateReq struct {
	g.Meta `path:"/agreement/{id}" tags:"协议管理" method:"put" summary:"更新协议"`
	model.Author

	ID         int64  `p:"id" v:"required#协议ID不能为空"`
	Content    string `json:"content" v:"required#协议内容不能为空"`
	PublishNow bool   `json:"publish_now" dc:"是否立即发布"`
	Status     *int   `json:"status" dc:"0:草稿 1:已发布 2:已归档"`
}

type AgreementUpdateRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"新创建的协议版本ID"`
}

type AgreementDeleteReq struct {
	g.Meta `path:"/agreement/{id}" tags:"协议管理" method:"delete" summary:"删除协议"`
	model.Author

	ID int64 `p:"id" v:"required#协议ID不能为空"`
}

type AgreementDeleteRes struct {
	model.EmptyRes
}

type AgreementGetReq struct {
	g.Meta `path:"/agreement/{id}" tags:"协议管理" method:"get" summary:"获取协议详情"`
	model.Author

	ID int64 `p:"id" v:"required#协议ID不能为空"`
}

type AgreementGetRes struct {
	g.Meta `mime:"application/json"`
	*AgreementItem
}

type AgreementListReq struct {
	g.Meta `path:"/agreements" tags:"协议管理" method:"get" summary:"协议列表"`
	model.Author
	model.PageReq

	Name   string `p:"name"`
	Status *int   `p:"status"`
}

type AgreementListRes struct {
	g.Meta `mime:"application/json"`

	List []*AgreementItem `json:"list"`
	*model.PageRes
}

type AgreementItem struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	MajorVersion int    `json:"major_version"`
	MinorVersion int    `json:"minor_version"`
	PatchVersion int    `json:"patch_version"`
	VersionCode  int    `json:"version_code"`
	Status       int    `json:"status"`
	Content      string `json:"content"`
	Version      string `json:"version"`
	PublishedAt  int64  `json:"published_at"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type AgreementGetLatestReq struct {
	g.Meta `path:"/agreement/latest" tags:"协议管理" method:"get" summary:"获取最新协议"`
	model.Author

	Name string `p:"name" v:"required#协议名称不能为空"`
}

type AgreementGetLatestRes struct {
	g.Meta `mime:"application/json"`
	*AgreementItem
}

type UserAgreementListReq struct {
	g.Meta `path:"/user/agreement" tags:"协议管理" method:"get" summary:"获取用户同意的协议列表"`
	model.Author
	UserID string `json:"user_id" v:"required#用户ID不能为空"`
}

type UserAgreementListRes struct {
	g.Meta `mime:"application/json"`

	List []*UserAgreementItem `json:"list"`
}

type UserAgreementItem struct {
	ID            int64  `json:"id"`
	UserID        string `json:"user_id"`
	AgreementID   int64  `json:"agreement_id"`
	AgreementName string `json:"agreement_name"`
	VersionCode   int    `json:"version_code"`
	Agreed        bool   `json:"agreed"`
	CreatedAt     int64  `json:"created_at"`
}
