package service

import (
	"IdentifyService/api/v1/system"
	commonModel "IdentifyService/internal/common/model"
	commonService "IdentifyService/internal/common/service"
	"context"
)

var LogService = logService{}

type logService struct {
}

func (c *logService) ListLoginLog(ctx context.Context, req *system.ListLoginLogReq) (res *system.ListLoginLogRes, err error) {
	in := &commonModel.ListLogReq{
		OrgID:     req.OrgID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageReq:   req.PageReq,
	}

	result, err := commonService.Log().ListLoginLog(ctx, in)

	res = &system.ListLoginLogRes{
		List:    result.List,
		PageRes: result.PageRes,
	}
	return
}

func (c *logService) ListOperLog(ctx context.Context, req *system.ListOperLogReq) (res *system.ListOperLogRes, err error) {
	in := &commonModel.ListLogReq{
		OrgID:     req.OrgID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageReq:   req.PageReq,
	}
	result, err := commonService.Log().ListOperLog(ctx, in)

	res = &system.ListOperLogRes{
		List:    result.List,
		PageRes: result.PageRes,
	}
	return
}
