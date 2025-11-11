package controller

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/service"
)

func (c *userController) UpdateUserType(ctx context.Context, req *system.UpdateUserTypeReq) (res *system.UpdateUserTypeRes, err error) {
	err = service.SysUser().UpdateUserType(ctx, req.UserID, req.UserType)
	return
}
