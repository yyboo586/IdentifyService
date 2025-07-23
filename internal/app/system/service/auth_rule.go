package service

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
)

type (
	IAuthRule interface {
		// 添加菜单/按钮
		Add(ctx context.Context, req *system.RuleAddReq) (id int64, err error)
		// 删除菜单/按钮
		DeleteByIDs(ctx context.Context, ids []int64) (err error)
		// 更新菜单/按钮
		Update(ctx context.Context, req *system.RuleUpdateReq) (err error)

		// 根据用户ID获取用户目录/菜单树形结构
		GetMenuTreesByUserID(ctx context.Context, userID string, includeButton bool) (out []*model.AuthRuleNode, err error)
		// 根据用户ID获取用户按钮列表
		GetButtonListByUserID(ctx context.Context, userID string) (out []*model.AuthRule, err error)

		// 根据角色ID获取用户目录/菜单列表
		// GetMenuListByRoleID(ctx context.Context, roleID int64) (out []*model.AuthRule, err error)

		// 获取单个菜单详细信息，不包含子节点
		// GetDetailsByID(ctx context.Context, id int64) (out *model.AuthRule, err error)
		// 获取单个菜单的树形结构，包含子节点
		// GetTreeByID(ctx context.Context, id int64) (out *model.AuthRuleNode, err error)
		// GetAllAuthRules 获取所有AuthRule
		GetAllAuthRules(ctx context.Context) (list []*model.AuthRule, err error)
	}
)

var (
	localSysAuthRule IAuthRule
)

func AuthRule() IAuthRule {
	if localSysAuthRule == nil {
		panic("implement not found for interface IAuthRule, forgot register?")
	}
	return localSysAuthRule
}

func RegisterAuthRule(i IAuthRule) {
	localSysAuthRule = i
}
