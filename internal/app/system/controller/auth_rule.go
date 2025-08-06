package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/service"
	"context"
)

var MenuController = menuController{}

type menuController struct {
}

func (c *menuController) Add(ctx context.Context, req *system.RuleAddReq) (res *system.RuleAddRes, err error) {
	ruleID, err := service.AuthRule().Add(ctx, req)
	if err != nil {
		return
	}

	res = &system.RuleAddRes{
		ID: ruleID,
	}
	return
}

func (c *menuController) Delete(ctx context.Context, req *system.RuleDeleteReq) (res *system.RuleDeleteRes, err error) {
	err = service.AuthRule().DeleteByIDs(ctx, req.Ids)
	return
}

func (c *menuController) Update(ctx context.Context, req *system.RuleUpdateReq) (res *system.RuleUpdateRes, err error) {
	err = service.AuthRule().Update(ctx, req)
	return
}

// 权限管理场景：获取完整权限树
func (c *menuController) GetFullAuthRuleTree(ctx context.Context, req *system.RuleListTreeReq) (res *system.RuleListTreeRes, err error) {
	res = &system.RuleListTreeRes{}
	res.List, err = service.AuthRule().GetFullAuthRuleTree(ctx, service.ContextService().Get(ctx).User.ID)
	if err != nil {
		return
	}
	return
}
