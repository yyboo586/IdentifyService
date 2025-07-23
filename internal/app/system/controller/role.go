package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"context"
)

var RoleController = roleController{}

type roleController struct {
}

func (c *roleController) Add(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error) {
	id, err := service.Role().Add(ctx, req)
	if err != nil {
		return
	}

	res = &system.RoleAddRes{
		ID: id,
	}
	return
}

func (c *roleController) Delete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error) {
	err = service.Role().DeleteByIDs(ctx, req.IDs)
	return
}

func (c *roleController) Edit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = service.Role().Edit(ctx, req)
	return
}

func (c *roleController) Get(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	var roleInfo *model.Role
	roleInfo, err = service.Role().Get(ctx, req.ID)
	if err != nil {
		return
	}

	res = new(system.RoleGetRes)
	res.Role = c.formatRoleInfo(roleInfo)
	res.MenuIDs, err = service.Role().GetFilteredNamedPolicy(ctx, req.ID)
	return
}

func (c *roleController) ListByOrgID(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	list, err := service.Role().ListByOrgID(ctx, req.OrgID)
	if err != nil {
		return
	}

	res = new(system.RoleListRes)
	res.List = make([]*system.RoleInfo, 0, len(list))
	for _, v := range list {
		res.List = append(res.List, c.formatRoleInfo(v))
	}
	return
}

func (c *roleController) formatRoleInfo(in *model.Role) (out *system.RoleInfo) {
	out = &system.RoleInfo{
		ID:        in.ID,
		PID:       in.PID,
		OrgID:     in.OrgID,
		Name:      in.Name,
		Enabled:   in.Status == model.RoleStatusEnabled,
		CreatorID: in.CreatorID,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
