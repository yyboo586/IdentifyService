package system

import (
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrgCreateReq struct {
	g.Meta `path:"/orgs" tags:"组织架构/组织管理" method:"post" summary:"组织创建"`
	model.Author
	Name        string `json:"name" v:"required#组织名称不能为空" dc:"组织名称"`
	PID         string `json:"pid" v:"required#父级ID不能为空" dc:"父级ID"`
	ManagerID   string `json:"manager_id" dc:"负责人ID"`
	ManagerName string `json:"manager_name" dc:"负责人名称"`
}

type OrgCreateRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id" dc:"组织ID"`
}

type OrgDeleteReq struct {
	g.Meta `path:"/orgs/{org_id}" tags:"组织架构/组织管理" method:"delete" summary:"组织删除"`
	model.Author
	OrgID string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type OrgEdiReq struct {
	g.Meta `path:"/orgs/{org_id}" tags:"组织架构/组织管理" method:"put" summary:"组织更新"`
	model.Author
	OrgID       string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name        string `json:"name" dc:"组织名称"`
	ManagerID   string `json:"manager_id" dc:"负责人ID"`
	ManagerName string `json:"manager_name" dc:"负责人名称"`
	Status      int64  `json:"status" dc:"组织状态(1:启用,2:禁用)"`
}

type OrgEditRes struct {
	g.Meta `mime:"application/json"`
}

type OrgGetReq struct {
	g.Meta `path:"/orgs/{org_id}" tags:"组织架构/组织管理" method:"get" summary:"组织详情"`
	model.Author
	OrgID string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgGetRes struct {
	g.Meta `mime:"application/json"`
	*OrgInfo
}

type OrgGetTreeReq struct {
	g.Meta `path:"/orgs/{org_id}/tree" tags:"组织架构/组织管理" method:"get" summary:"组织树"`
	model.Author
	OrgID string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgGetTreeRes struct {
	g.Meta `mime:"application/json"`
	*OrgTreeNode
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
