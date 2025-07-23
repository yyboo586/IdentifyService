package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/service"
	"context"
)

var LogController = logController{}

type logController struct {
}

func (c *logController) ListLoginLog(ctx context.Context, req *system.ListLoginLogReq) (res *system.ListLoginLogRes, err error) {
	res, err = service.Log().ListLoginLog(ctx, req)
	return
}

func (c *logController) ListOperLog(ctx context.Context, req *system.ListOperLogReq) (res *system.ListOperLogRes, err error) {
	res, err = service.Log().ListOperLog(ctx, req)
	return
}
