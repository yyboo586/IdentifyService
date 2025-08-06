package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RoleAddReq struct {
	g.Meta `path:"/role" tags:"角色管理" method:"post" summary:"创建角色"`
	model.Author
	OrgID   string  `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Pid     int64   `json:"pid" v:"required#父级角色ID不能为空" dc:"父级角色ID"`
	Name    string  `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	MenuIDs []int64 `json:"menu_ids" dc:"资源ID列表"`
}

type RoleAddRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"角色ID"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role" tags:"角色管理" method:"delete" summary:"删除角色"`
	model.Author
	IDs []int64 `json:"ids" v:"required#角色IDs不能为空" dc:"角色ID列表"`
}

type RoleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type RoleEditReq struct {
	g.Meta `path:"/role/{id}" tags:"角色管理" method:"put" summary:"编辑角色(全量更新)"`
	model.Author
	ID      int64   `p:"id" v:"required#角色ID必须" dc:"角色ID"`
	Pid     int64   `json:"pid" dc:"父级ID"`
	Name    string  `json:"name" dc:"角色名称"`
	MenuIDs []int64 `json:"menu_ids" dc:"资源ID列表"`
}

type RoleEditRes struct {
	g.Meta `mime:"application/json"`
}

type RoleEditStatusReq struct {
	g.Meta `path:"/role/{id}/status" tags:"角色管理" method:"put" summary:"编辑角色状态"`
	model.Author
	ID      int64 `p:"id" v:"required#角色ID不能为空" dc:"角色ID"`
	Enabled bool  `json:"enabled" v:"required#状态不能为空" dc:"状态(启用/禁用)"`
}

type RoleEditStatusRes struct {
	g.Meta `mime:"application/json"`
}

type RoleGetReq struct {
	g.Meta `path:"/role/{id}" tags:"角色管理" method:"get" summary:"获取角色详情"`
	model.Author
	ID int64 `p:"id" v:"required#角色ID不能为空" dc:"角色ID"`
}

type RoleGetRes struct {
	g.Meta  `mime:"application/json"`
	Role    *RoleInfo `json:"role" dc:"角色详情"`
	MenuIDs []int64   `json:"menu_ids" dc:"资源ID列表"`
}

type RoleListReq struct {
	g.Meta `path:"/role/trees" tags:"角色管理" method:"get" summary:"角色列表(树形结构)"`
	model.Author
	OrgID string `json:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	model.PageReq
}

type RoleListRes struct {
	g.Meta `mime:"application/json"`
	List   []*RoleNode `json:"list" dc:"角色列表"`
	model.PageRes
}

type RoleInfo struct {
	ID        int64       `json:"id" dc:"角色ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	PID       int64       `json:"pid" dc:"父级ID"`
	Name      string      `json:"name" dc:"角色名称"`
	Enabled   bool        `json:"enabled" dc:"是否启用"`
	CreatorID string      `json:"creator_id" dc:"创建人ID"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}

type RoleNode struct {
	*RoleInfo
	Children []*RoleNode `json:"children"`
}
