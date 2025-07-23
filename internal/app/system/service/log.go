package service

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ILog interface {
	InvokeLoginLog(in *model.LoginLog)
	InvokeOperLog(in *model.OperLog)

	ListLoginLog(ctx context.Context, req *system.ListLoginLogReq) (res *system.ListLoginLogRes, err error)
	ListOperLog(ctx context.Context, req *system.ListOperLogReq) (res *system.ListOperLogRes, err error)

	OperationLog(r *ghttp.Request)
}

var localLog ILog

func Log() ILog {
	if localLog == nil {
		panic("implement not found for interface ILog, forgot register?")
	}
	return localLog
}

func RegisterLog(i ILog) {
	localLog = i
}
