package interfaces

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/database/gdb"
)

type IOrg interface {
	// 创建组织
	Create(ctx context.Context, tx gdb.TX, in *model.Org) (err error)
	// 删除组织
	Delete(ctx context.Context, id string) (err error)
	// 修改组织信息
	Edit(ctx context.Context, in *model.Org) (err error)
	// 获取单个组织详情
	Get(ctx context.Context, id string) (out *model.Org, err error)
	// 获取组织树
	GetTree(ctx context.Context, id string) (out *system.OrgTreeNode, err error)

	// 组织状态检查
	IsEnabled(ctx context.Context, orgID string) (orgInfo *model.Org, err error)

	// 根据组织ID获取角色列表(树形结构)
	GetRoleTrees(ctx context.Context, orgID string) (out []*system.RoleNode, err error)
}
