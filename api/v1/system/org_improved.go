package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ==================== 组织基础操作 ====================

// OrgCreateReq 创建组织
type OrgCreateReq struct {
	g.Meta `path:"/orgs" tags:"组织管理" method:"post" summary:"创建组织"`
	model.Author
	Name        string `json:"name" v:"required#组织名称不能为空" dc:"组织名称"`
	PID         string `json:"pid" v:"required#父级ID不能为空" dc:"父级ID"`
	Code        string `json:"code" dc:"组织编码"`
	ManagerID   string `json:"manager_id" dc:"负责人ID"`
	ManagerName string `json:"manager_name" dc:"负责人名称"`
	Description string `json:"description" dc:"组织描述"`
	Sort        int    `json:"sort" dc:"排序"`
}

type OrgCreateRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id" dc:"组织ID"`
}

// OrgGetReq 获取组织详情
type OrgGetReq2 struct {
	g.Meta `path:"/orgs/{id}" tags:"组织管理" method:"get" summary:"获取组织详情"`
	model.Author
	ID string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgGetRes2 struct {
	g.Meta `mime:"application/json"`
	*OrgDetailInfo
}

// OrgUpdateReq 更新组织信息
type OrgUpdateReq struct {
	g.Meta `path:"/orgs/{id}" tags:"组织管理" method:"put" summary:"更新组织信息"`
	model.Author
	ID          string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name        string `json:"name" dc:"组织名称"`
	Code        string `json:"code" dc:"组织编码"`
	ManagerID   string `json:"manager_id" dc:"负责人ID"`
	ManagerName string `json:"manager_name" dc:"负责人名称"`
	Description string `json:"description" dc:"组织描述"`
	Sort        int    `json:"sort" dc:"排序"`
}

type OrgUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// OrgDeleteReq 删除组织
type OrgDeleteReq2 struct {
	g.Meta `path:"/orgs/{id}" tags:"组织管理" method:"delete" summary:"删除组织"`
	model.Author
	ID string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgDeleteRes2 struct {
	g.Meta `mime:"application/json"`
}

// ==================== 组织状态管理 ====================

// OrgStatusUpdateReq 更新组织状态
type OrgStatusUpdateReq struct {
	g.Meta `path:"/orgs/{id}/status" tags:"组织管理" method:"patch" summary:"更新组织状态"`
	model.Author
	ID      string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	Enabled bool   `json:"enabled" v:"required#状态不能为空" dc:"组织状态(禁用/启用)"`
}

type OrgStatusUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 组织列表和搜索 ====================

// OrgListReq 获取组织列表
type OrgListReq struct {
	g.Meta `path:"/orgs" tags:"组织管理" method:"get" summary:"获取组织列表"`
	model.Author
	Page     int    `p:"page" d:"1" dc:"页码"`
	PageSize int    `p:"page_size" d:"20" dc:"每页数量"`
	Keyword  string `p:"keyword" dc:"搜索关键词"`
	PID      string `p:"pid" dc:"父级ID过滤"`
	Enabled  *bool  `p:"enabled" dc:"状态过滤"`
	SortBy   string `p:"sort_by" d:"created_at" dc:"排序字段"`
	SortDesc bool   `p:"sort_desc" d:"true" dc:"是否降序"`
}

type OrgListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*OrgInfo2 `json:"list" dc:"组织列表"`
	Total    int         `json:"total" dc:"总数"`
	Page     int         `json:"page" dc:"当前页"`
	PageSize int         `json:"page_size" dc:"每页数量"`
}

// ==================== 组织树形结构 ====================

// OrgTreeReq 获取组织树
type OrgTreeReq struct {
	g.Meta `path:"/orgs/tree" tags:"组织管理" method:"get" summary:"获取组织树形结构"`
	model.Author
	RootID   string `p:"root_id" dc:"根节点ID，不传则获取全部"`
	MaxDepth int    `p:"max_depth" d:"5" dc:"最大深度"`
	Enabled  *bool  `p:"enabled" dc:"状态过滤"`
}

type OrgTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*OrgTreeNode `json:"list" dc:"组织树列表"`
}

// OrgSubTreeReq 获取子组织树
type OrgSubTreeReq struct {
	g.Meta `path:"/orgs/{id}/subtree" tags:"组织管理" method:"get" summary:"获取子组织树"`
	model.Author
	ID       string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	MaxDepth int    `p:"max_depth" d:"3" dc:"最大深度"`
	Enabled  *bool  `p:"enabled" dc:"状态过滤"`
}

type OrgSubTreeRes struct {
	g.Meta `mime:"application/json"`
	*OrgTreeNode
}

// ==================== 组织移动和复制 ====================

// OrgMoveReq 移动组织
type OrgMoveReq struct {
	g.Meta `path:"/orgs/{id}/move" tags:"组织管理" method:"patch" summary:"移动组织"`
	model.Author
	ID     string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	NewPID string `json:"new_pid" v:"required#新父级ID不能为空" dc:"新父级ID"`
}

type OrgMoveRes struct {
	g.Meta `mime:"application/json"`
}

// OrgCopyReq 复制组织
type OrgCopyReq struct {
	g.Meta `path:"/orgs/{id}/copy" tags:"组织管理" method:"post" summary:"复制组织"`
	model.Author
	ID        string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
	NewPID    string `json:"new_pid" v:"required#新父级ID不能为空" dc:"新父级ID"`
	NewName   string `json:"new_name" v:"required#新组织名称不能为空" dc:"新组织名称"`
	CopyUsers bool   `json:"copy_users" dc:"是否复制用户"`
	CopyRoles bool   `json:"copy_roles" dc:"是否复制角色"`
	CopyPerms bool   `json:"copy_perms" dc:"是否复制权限"`
}

type OrgCopyRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id" dc:"新组织ID"`
}

// ==================== 组织统计信息 ====================

// OrgStatsReq 获取组织统计信息
type OrgStatsReq struct {
	g.Meta `path:"/orgs/{id}/stats" tags:"组织管理" method:"get" summary:"获取组织统计信息"`
	model.Author
	ID string `p:"id" v:"required#组织ID不能为空" dc:"组织ID"`
}

type OrgStatsRes struct {
	g.Meta `mime:"application/json"`
	*OrgStatsInfo
}

// ==================== 批量操作 ====================

// OrgBatchDeleteReq 批量删除组织
type OrgBatchDeleteReq struct {
	g.Meta `path:"/orgs/batch" tags:"组织管理" method:"delete" summary:"批量删除组织"`
	model.Author
	IDs []string `json:"ids" v:"required#组织ID列表不能为空" dc:"组织ID列表"`
}

type OrgBatchDeleteRes struct {
	g.Meta       `mime:"application/json"`
	SuccessCount int      `json:"success_count" dc:"成功删除数量"`
	FailedIDs    []string `json:"failed_ids" dc:"删除失败的组织ID"`
}

// OrgBatchStatusUpdateReq 批量更新组织状态
type OrgBatchStatusUpdateReq struct {
	g.Meta `path:"/orgs/batch/status" tags:"组织管理" method:"patch" summary:"批量更新组织状态"`
	model.Author
	IDs     []string `json:"ids" v:"required#组织ID列表不能为空" dc:"组织ID列表"`
	Enabled bool     `json:"enabled" v:"required#状态不能为空" dc:"组织状态"`
}

type OrgBatchStatusUpdateRes struct {
	g.Meta       `mime:"application/json"`
	SuccessCount int      `json:"success_count" dc:"成功更新数量"`
	FailedIDs    []string `json:"failed_ids" dc:"更新失败的组织ID"`
}

// ==================== 数据模型 ====================

// OrgInfo2 组织基本信息
type OrgInfo2 struct {
	ID          string      `json:"id" dc:"组织ID"`
	PID         string      `json:"pid" dc:"父级ID"`
	Name        string      `json:"name" dc:"组织名称"`
	Code        string      `json:"code" dc:"组织编码"`
	ManagerID   string      `json:"manager_id" dc:"负责人ID"`
	ManagerName string      `json:"manager_name" dc:"负责人名称"`
	Description string      `json:"description" dc:"组织描述"`
	Enabled     bool        `json:"enabled" dc:"组织状态"`
	Sort        int         `json:"sort" dc:"排序"`
	CreatedAt   *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at" dc:"修改时间"`
}

// OrgDetailInfo 组织详细信息
type OrgDetailInfo struct {
	*OrgInfo2
	UserCount     int    `json:"user_count" dc:"用户数量"`
	RoleCount     int    `json:"role_count" dc:"角色数量"`
	ChildrenCount int    `json:"children_count" dc:"子组织数量"`
	Level         int    `json:"level" dc:"组织层级"`
	Path          string `json:"path" dc:"组织路径"`
}

// OrgTreeNode 组织树节点
type OrgTreeNode2 struct {
	*OrgInfo2
	Children []*OrgTreeNode2 `json:"children" dc:"子组织"`
	Level    int             `json:"level" dc:"层级"`
	Path     string          `json:"path" dc:"路径"`
}

// OrgStatsInfo 组织统计信息
type OrgStatsInfo struct {
	TotalUsers     int `json:"total_users" dc:"总用户数"`
	ActiveUsers    int `json:"active_users" dc:"活跃用户数"`
	TotalRoles     int `json:"total_roles" dc:"总角色数"`
	TotalChildren  int `json:"total_children" dc:"子组织数"`
	DirectChildren int `json:"direct_children" dc:"直接子组织数"`
	MaxDepth       int `json:"max_depth" dc:"最大深度"`
}
