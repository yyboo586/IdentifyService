package service

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
)

type IOrg interface {
	// 添加组织
	Add(ctx context.Context, in *model.Org) (id string, err error)

	// 删除组织
	Delete(ctx context.Context, id string) (err error)

	// 编辑组织基本信息
	EditBasicInfo(ctx context.Context, in *model.Org) (err error)
	// 编辑组织状态
	EditStatus(ctx context.Context, id string, enabled bool) (err error)

	// 获取组织详情
	Get(ctx context.Context, id string) (out *system.OrgInfo, err error)
	// 获取组织树
	GetTree(ctx context.Context, id string) (out *system.OrgTreeNode, err error)
	// 获取组织列表
	ListTrees(ctx context.Context) (out []*system.OrgTreeNode, err error)
}

var localOrg IOrg

func Org() IOrg {
	if localOrg == nil {
		panic("implement not found for interface IOrg, forgot register?")
	}
	return localOrg
}

func RegisterOrg(i IOrg) {
	localOrg = i
}
