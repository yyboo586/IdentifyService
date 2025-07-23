package logics

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"context"
	"slices"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	logOnce     sync.Once
	logInstance *logicsLog
)

func NewLog() service.ILog {
	logOnce.Do(func() {
		logInstance = &logicsLog{
			ctx:  context.Background(),
			Pool: grpool.New(100),
		}
	})
	return logInstance
}

type logicsLog struct {
	ctx  context.Context
	Pool *grpool.Pool
}

func (s *logicsLog) InvokeLoginLog(in *model.LoginLog) {
	s.Pool.Add(s.ctx, func(ctx context.Context) {
		s.addLoginLog(ctx, in)
	})
}

func (s *logicsLog) InvokeOperLog(in *model.OperLog) {
	s.Pool.Add(s.ctx, func(ctx context.Context) {
		s.addOperLog(ctx, in)
	})
}

func (s *logicsLog) addLoginLog(ctx context.Context, in *model.LoginLog) (err error) {
	dataInsert := g.Map{
		dao.LoginLog.Columns().OrgID:     in.OrgID,
		dao.LoginLog.Columns().LoginName: in.LoginName,
		dao.LoginLog.Columns().IP:        in.IP,
		dao.LoginLog.Columns().Browser:   in.Browser,
		dao.LoginLog.Columns().Message:   in.Message,
		dao.LoginLog.Columns().LoginTime: in.LoginTime,
		dao.LoginLog.Columns().CreatedAt: in.CreatedAt,
	}
	if in.Success {
		dataInsert[dao.LoginLog.Columns().Status] = model.LoginStatusSuccess
	} else {
		dataInsert[dao.LoginLog.Columns().Status] = model.LoginStatusFailed
	}
	_, err = dao.LoginLog.Ctx(ctx).Insert(dataInsert)
	return
}

func (s *logicsLog) addOperLog(ctx context.Context, in *model.OperLog) (err error) {
	insertData := map[string]interface{}{
		dao.OperLog.Columns().OrgID:      in.OrgID,
		dao.OperLog.Columns().OperName:   in.OperName,
		dao.OperLog.Columns().OperUrl:    in.OperUrl,
		dao.OperLog.Columns().OperMethod: in.OperMethod,
		dao.OperLog.Columns().OperIP:     in.OperIP,
		dao.OperLog.Columns().OperTime:   in.OperTime,
		dao.OperLog.Columns().CreatedAt:  in.CreatedAt,
	}
	_, err = dao.OperLog.Ctx(ctx).Insert(insertData)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (s *logicsLog) ListLoginLog(ctx context.Context, in *system.ListLoginLogReq) (out *system.ListLoginLogRes, err error) {
	out = new(system.ListLoginLogRes)
	if in.PageNum == 0 {
		in.PageNum = 1
	}
	out.CurrentPage = in.PageNum
	if in.PageSize == 0 {
		in.PageSize = consts.PageSize
	}

	m := dao.LoginLog.Ctx(ctx).Where(dao.LoginLog.Columns().OrgID, in.OrgID)
	if in.UserName != "" {
		m = m.Where(dao.LoginLog.Columns().LoginName, in.UserName)
	}
	if in.StartTime != nil {
		m = m.WhereGTE(dao.LoginLog.Columns().LoginTime, in.StartTime)
	}
	if in.EndTime != nil {
		m = m.WhereLTE(dao.LoginLog.Columns().LoginTime, in.EndTime)
	}

	out.Total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var list []*entity.TLoginLog
	err = m.Page(in.PageNum, in.PageSize).OrderDesc(dao.LoginLog.Columns().LoginTime).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var loginList = make([]*model.LoginLog, len(list))
	for k, v := range list {
		loginList[k] = s.convertEntityToModelLoginLog(v)
	}
	out.List = loginList

	return
}

func (s *logicsLog) ListOperLog(ctx context.Context, in *system.ListOperLogReq) (out *system.ListOperLogRes, err error) {
	out = new(system.ListOperLogRes)
	if in.PageNum == 0 {
		in.PageNum = 1
	}
	out.CurrentPage = in.PageNum
	if in.PageSize == 0 {
		in.PageSize = consts.PageSize
	}

	m := dao.OperLog.Ctx(ctx).Where(dao.OperLog.Columns().OrgID, in.OrgID)
	if in.OperName != "" {
		m = m.Where(dao.OperLog.Columns().OperName, in.OperName)
	}
	if in.StartTime != nil {
		m = m.WhereGTE(dao.OperLog.Columns().OperTime, in.StartTime)
	}
	if in.EndTime != nil {
		m = m.WhereLTE(dao.OperLog.Columns().OperTime, in.EndTime)
	}

	out.Total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var list []*entity.TOperLog
	err = m.Page(in.PageNum, in.PageSize).OrderDesc(dao.OperLog.Columns().OperTime).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var operList = make([]*model.OperLog, len(list))
	for k, v := range list {
		operList[k] = s.convertEntityToModelOperLog(v)
	}
	out.List = operList

	return
}

func (s *logicsLog) OperationLog(r *ghttp.Request) {
	vars, err := g.Cfg().Get(r.GetCtx(), "gfToken.excludePaths")
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	excludePaths := gconv.Strings(vars)
	if slices.Contains(excludePaths, r.RequestURI) {
		return
	}

	operatorInfo := service.ContextService().GetUser(r.GetCtx())
	data := &model.OperLog{
		OrgID:      operatorInfo.OrgID,
		OperName:   operatorInfo.Name,
		OperUrl:    r.Request.URL.Path,
		OperMethod: r.Method,
		OperIP:     libUtils.GetClientIp(r.GetCtx()),
		OperTime:   gtime.Now(),
	}
	service.Log().InvokeOperLog(data)
}

func (s *logicsLog) convertEntityToModelLoginLog(in *entity.TLoginLog) *model.LoginLog {
	return &model.LoginLog{
		ID:        in.ID,
		OrgID:     in.OrgID,
		LoginName: in.LoginName,
		IP:        in.IP,
		Browser:   in.Browser,
		Success:   model.LoginStatus(in.Status) == model.LoginStatusSuccess,
		Message:   in.Msg,
		LoginTime: in.LoginTime,
		CreatedAt: in.CreatedAt,
	}
}

func (s *logicsLog) convertEntityToModelOperLog(in *entity.TOperLog) *model.OperLog {
	return &model.OperLog{
		ID:         in.ID,
		OrgID:      in.OrgID,
		OperName:   in.OperName,
		OperUrl:    in.OperUrl,
		OperMethod: in.OperMethod,
		OperIP:     in.OperIP,
		OperTime:   in.OperTime,
		CreatedAt:  in.CreatedAt,
	}
}
