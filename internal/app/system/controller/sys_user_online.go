package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/service"
	"context"
)

var UserOnline = new(SysUserOnlineController)

type SysUserOnlineController struct{}

func (c *SysUserOnlineController) List(ctx context.Context, req *system.SysUserOnlineSearchReq) (res *system.SysUserOnlineSearchRes, err error) {
	res, err = service.SysUserOnline().GetOnlineListPage(ctx, req)
	return
}

func (c *SysUserOnlineController) ForceLogout(ctx context.Context, req *system.SysUserOnlineForceLogoutReq) (res *system.SysUserOnlineForceLogoutRes, err error) {
	err = service.SysUserOnline().ForceLogout(ctx, req.Ids)
	return
}
