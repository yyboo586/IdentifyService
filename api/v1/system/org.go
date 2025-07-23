package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrgAddReq struct {
	g.Meta `path:"/org" tags:"组织管理" method:"post" summary:"组织创建"`
	model.Author
	Name        string `json:"name" v:"required#组织名称不能为空" dc:"组织名称"`
	PID         string `json:"pid" v:"required#父级ID不能为空" dc:"父级ID"`
	ManagerID   string `json:"manager_id" dc:"负责人ID"`
	ManagerName string `json:"manager_name" dc:"负责人名称"`
}

type OrgAddRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id" dc:"组织ID"`
}

type OrgDeleteReq struct {
	g.Meta `path:"/org/{id}" tags:"组织管理" method:"delete" summary:"组织删除"`
	model.Author
	ID string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type OrgEditBasicInfoReq struct {
	g.Meta `path:"/org/{id}" tags:"组织管理" method:"put" summary:"组织基本信息更新(全量更新)"`
	model.Author
	ID          string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name        string `json:"name" dc:"组织名称"`
	ManagerID   string `json:"manager_id" dc:"负责人ID"`
	ManagerName string `json:"manager_name" dc:"负责人名称"`
}

type OrgEditBasicInfoRes struct {
	g.Meta `mime:"application/json"`
}

type OrgStatusEditReq struct {
	g.Meta `path:"/org/{id}/status" tags:"组织管理" method:"put" summary:"组织状态更新"`
	model.Author
	ID      string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	Enabled bool   `json:"enabled" v:"required#状态不能为空" dc:"组织状态(禁用/启用)"`
}

type OrgStatusEditRes struct {
	g.Meta `mime:"application/json"`
}

type OrgGetReq struct {
	g.Meta `path:"/org/{id}" tags:"组织管理" method:"get" summary:"组织详情"`
	model.Author
	ID string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgGetRes struct {
	g.Meta `mime:"application/json"`
	*OrgInfo
}

type OrgGetTreeReq struct {
	g.Meta `path:"/org/{id}/tree" tags:"组织管理" method:"get" summary:"组织详情(树形结构)"`
	model.Author
	ID string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgGetTreeRes struct {
	g.Meta `mime:"application/json"`
	*OrgTreeNode
}

type OrgListTreeReq struct {
	g.Meta `path:"/org/trees" tags:"组织管理" method:"get" summary:"组织列表(树形结构)"`
	model.Author
}

type OrgListTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*OrgTreeNode `json:"list"`
}

type OrgInfo struct {
	ID          string      `json:"id" dc:"组织ID"`
	PID         string      `json:"pid" dc:"父级ID"`
	Name        string      `json:"name" dc:"组织名称"`
	ManagerID   string      `json:"manager_id" dc:"负责人ID"`
	ManagerName string      `json:"manager_name" dc:"负责人名称"`
	Enabled     bool        `json:"enabled" dc:"组织状态(禁用/启用)"`
	CreatedAt   *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at" dc:"修改时间"`
}

type OrgTreeNode struct {
	*OrgInfo
	Children []*OrgTreeNode `json:"children"`
}
