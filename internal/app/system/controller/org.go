package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"context"
)

var OrgController = orgController{}

type orgController struct {
}

func (c *orgController) Add(ctx context.Context, req *system.OrgAddReq) (res *system.OrgAddRes, err error) {
	in := &model.Org{
		PID:         req.PID,
		Name:        req.Name,
		ManagerID:   req.ManagerID,
		ManagerName: req.ManagerName,
	}
	orgID, err := service.Org().Add(ctx, in)
	if err != nil {
		return
	}

	res = &system.OrgAddRes{
		ID: orgID,
	}
	return
}

func (c *orgController) Delete(ctx context.Context, req *system.OrgDeleteReq) (res *system.OrgDeleteRes, err error) {
	err = service.Org().Delete(ctx, req.ID)

	return
}

func (c *orgController) EditBasicInfo(ctx context.Context, req *system.OrgEditBasicInfoReq) (res *system.OrgEditBasicInfoRes, err error) {
	in := &model.Org{
		ID:          req.ID,
		Name:        req.Name,
		ManagerID:   req.ManagerID,
		ManagerName: req.ManagerName,
	}
	err = service.Org().EditBasicInfo(ctx, in)

	return
}

func (c *orgController) EditStatus(ctx context.Context, req *system.OrgStatusEditReq) (res *system.OrgStatusEditRes, err error) {
	err = service.Org().EditStatus(ctx, req.ID, req.Enabled)

	return
}

func (c *orgController) Get(ctx context.Context, req *system.OrgGetReq) (res *system.OrgGetRes, err error) {
	out, err := service.Org().Get(ctx, req.ID)
	if err != nil {
		return
	}

	res = &system.OrgGetRes{
		OrgInfo: out,
	}
	return
}

func (c *orgController) GetTree(ctx context.Context, req *system.OrgGetTreeReq) (res *system.OrgGetTreeRes, err error) {
	out, err := service.Org().GetTree(ctx, req.ID)
	if err != nil {
		return
	}

	res = &system.OrgGetTreeRes{
		OrgTreeNode: out,
	}
	return
}

func (c *orgController) ListTree(ctx context.Context, req *system.OrgListTreeReq) (res *system.OrgListTreeRes, err error) {
	list, err := service.Org().ListTrees(ctx)
	if err != nil {
		return
	}

	res = &system.OrgListTreeRes{
		List: list,
	}
	return
}
