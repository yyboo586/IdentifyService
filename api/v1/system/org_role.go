package system

import (
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RoleCreateReq struct {
	g.Meta `path:"/orgs/{org_id}/roles" tags:"组织架构/角色管理" method:"post" summary:"创建角色"`
	model.Author
	OrgID              string  `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Pid                int64   `json:"pid" v:"required#父级角色ID不能为空" dc:"父级角色ID"`
	Name               string  `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	PermissionPointIDs []int64 `json:"permission_point_ids" dc:"权限点ID列表"`
}

type RoleCreateRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"角色ID"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/orgs/{org_id}/roles" tags:"组织架构/角色管理" method:"delete" summary:"删除角色"`
	model.Author
	OrgID   string  `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	RoleIDs []int64 `json:"role_ids" v:"required#角色IDs不能为空" dc:"角色ID列表"`
}

type RoleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type RoleEditReq struct {
	g.Meta `path:"/orgs/{org_id}/roles/{role_id}" tags:"组织架构/角色管理" method:"put" summary:"编辑角色"`
	model.Author
	OrgID              string  `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	RoleID             int64   `p:"role_id" v:"required#角色ID必须" dc:"角色ID"`
	PID                int64   `json:"pid" dc:"父级ID"`
	Name               string  `json:"name" dc:"角色名称"`
	PermissionPointIDs []int64 `json:"permission_point_ids" dc:"权限点ID列表"`
}

type RoleEditRes struct {
	g.Meta `mime:"application/json"`
}

type RoleEditStatusReq struct {
	g.Meta `path:"/orgs/{org_id}/roles/{role_id}/status" tags:"组织架构/角色管理" method:"put" summary:"编辑角色状态"`
	model.Author
	OrgID   string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	RoleID  int64  `p:"role_id" v:"required#角色ID不能为空" dc:"角色ID"`
	Enabled bool   `json:"enabled" v:"required#状态不能为空" dc:"状态(启用/禁用)"`
}

type RoleEditStatusRes struct {
	g.Meta `mime:"application/json"`
}

type RoleGetReq struct {
	g.Meta `path:"/orgs/{org_id}/roles/{role_id}" tags:"组织架构/角色管理" method:"get" summary:"获取角色详情"`
	model.Author
	OrgID  string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	RoleID int64  `p:"role_id" v:"required#角色ID不能为空" dc:"角色ID"`
}

type RoleGetRes struct {
	g.Meta           `mime:"application/json"`
	Role             *RoleInfo               `json:"role" dc:"角色详情"`
	PermissionPoints []model.PermissionPoint `json:"permission_points" dc:"资源ID列表"`
}

type RoleTreeReq struct {
	g.Meta `path:"/orgs/{org_id}/roles/trees" tags:"组织架构/角色管理" method:"get" summary:"角色列表(树形结构)"`
	model.Author
	OrgID string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	model.PageReq
}

type RoleTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*RoleNode `json:"list" dc:"角色列表"`
	model.PageRes
}

type RoleInfo struct {
	ID        int64       `json:"id" dc:"角色ID"`
	OrgID     string      `json:"org_id" dc:"组织ID"`
	PID       int64       `json:"pid" dc:"父级ID"`
	Name      string      `json:"name" dc:"角色名称"`
	CreatorID string      `json:"creator_id" dc:"创建人ID"`
	DeletorID string      `json:"deletor_id" dc:"删除人ID"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" dc:"删除时间"`
}

type RoleNode struct {
	*RoleInfo
	Children []*RoleNode `json:"children"`
}
