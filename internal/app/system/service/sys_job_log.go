package service

import (
	"context"

	"IdentifyService/api/v1/system"
)

type (
	ISysJobLog interface {
		Add(ctx context.Context, data map[string]interface{}) (err error)
		List(ctx context.Context, req *system.SysJobLogListReq) (listRes *system.SysJobLogListRes, err error)
		Delete(ctx context.Context, logIds []uint64) (err error)
	}
)

var (
	localSysJobLog ISysJobLog
)

func SysJobLog() ISysJobLog {
	if localSysJobLog == nil {
		panic("implement not found for interface ISysJobLog, forgot register?")
	}
	return localSysJobLog
}

func RegisterSysJobLog(i ISysJobLog) {
	localSysJobLog = i
}
