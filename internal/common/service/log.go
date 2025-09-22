package service

import (
	"context"

	"IdentifyService/internal/common/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ILog interface {
	InvokeLog(in *model.Log)

	ListLoginLog(ctx context.Context, in *model.ListLogReq) (res *model.ListLogRes, err error)
	ListOperLog(ctx context.Context, in *model.ListLogReq) (res *model.ListLogRes, err error)

	RecordLog(r *ghttp.Request)
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
