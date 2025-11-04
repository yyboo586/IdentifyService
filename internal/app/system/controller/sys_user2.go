package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func (c *userController) UpdateUserType(ctx context.Context, req *system.UpdateUserTypeReq) (res *system.UpdateUserTypeRes, err error) {
	err = service.SysUser().UpdateUserType(ctx, req.UserID, req.UserType)
	return
}
