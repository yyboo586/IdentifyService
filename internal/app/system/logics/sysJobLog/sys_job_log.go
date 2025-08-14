package sysJobLog

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/liberr"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysJobLog(New())
}

func New() service.ISysJobLog {
	return new(sSysJobLog)
}

type sSysJobLog struct {
}

func (s *sSysJobLog) Add(ctx context.Context, data map[string]interface{}) (err error) {
	_, err = dao.SysJobLog.Ctx(ctx).Insert(data)
	return
}

func (s *sSysJobLog) List(ctx context.Context, req *system.SysJobLogListReq) (listRes *system.SysJobLogListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		listRes = new(system.SysJobLogListRes)
		m := dao.SysJobLog.Ctx(ctx).Where(dao.SysJobLog.Columns().TargetName, req.TargetName)
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		err = m.Fields(system.SysJobLogListRes{}).Page(req.PageNum, req.PageSize).Order("id desc").Scan(&listRes.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sSysJobLog) Delete(ctx context.Context, logIds []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysJobLog.Ctx(ctx).Delete(dao.SysJobLog.Columns().ID+" in (?)", logIds)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
