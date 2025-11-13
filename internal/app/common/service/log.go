package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yyboo586/common/LogModule"
)

type ILog interface {
	WriteLog(ctx context.Context, logs []*LogModule.LogItem) (err error)
	ListLog(ctx context.Context, filter *LogModule.LogListFilter) (out []*LogModule.LogItem, err error)
}

func init() {
	logConfig := &LogModule.Config{
		DSN:         g.Cfg().MustGet(context.Background(), "database.default.link").String(),
		Group:       "default",
		TableName:   "t_log",
		EnableDebug: true,
		MaxBatch:    200,
	}
	logManager, err := LogModule.NewLogManager(logConfig)
	if err != nil {
		panic(err)
	}
	localLog = &logService{
		logManager: logManager,
	}
}

func Log() ILog {
	return localLog
}

var localLog ILog

type logService struct {
	logManager LogModule.ILogManager
}

func (s *logService) WriteLog(ctx context.Context, logs []*LogModule.LogItem) (err error) {
	return s.logManager.BatchWrite(ctx, logs)
}

func (s *logService) ListLog(ctx context.Context, filter *LogModule.LogListFilter) (logs []*LogModule.LogItem, err error) {
	return s.logManager.List(ctx, filter)
}
