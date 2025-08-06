package service

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
)

type (
	IAuthRule interface {
		// ========== 基础CRUD操作 ==========
		// 添加目录/菜单/按钮
		Add(ctx context.Context, req *system.RuleAddReq) (id int64, err error)
		// 删除目录/菜单/按钮
		DeleteByIDs(ctx context.Context, ids []int64) (err error)
		// 更新目录/菜单/按钮
		Update(ctx context.Context, req *system.RuleUpdateReq) (err error)

		// ========== 登录场景接口 ==========
		// 获取用户登录后的菜单树（目录+菜单，不包含按钮）
		GetUserMenuTree(ctx context.Context, userID string) (out []*model.AuthRuleNode, err error)
		// 获取用户登录后的按钮权限列表
		GetUserButtonList(ctx context.Context, userID string) (out []*model.AuthRule, err error)

		// ========== 权限管理场景接口 ==========
		// 获取完整的权限树（目录+菜单+按钮），用于权限管理界面
		GetFullAuthRuleTree(ctx context.Context, userID string) (out []*model.AuthRuleNode, err error)

		// ========== 权限过滤工具 ==========
		// 根据用户角色过滤权限ID列表
		FilterRuleIDsByUserID(ctx context.Context, ruleIDs []int64, userID string) (out []int64, err error)
		// 检查用户是否有指定权限
		HasPermission(ctx context.Context, userID string, ruleID int64) (hasPermission bool, err error)
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
