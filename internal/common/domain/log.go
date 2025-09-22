package domain

import (
	"IdentifyService/internal/common/dao"
	"IdentifyService/internal/common/model"
	"IdentifyService/internal/common/model/entity"
	"IdentifyService/internal/common/service"
	"IdentifyService/library/libUtils"
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
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

func (s *logicsLog) InvokeLog(in *model.Log) {
	s.Pool.Add(s.ctx, func(ctx context.Context) {
		s.addLog(ctx, in)
	})
}

func (s *logicsLog) addLog(ctx context.Context, in *model.Log) (err error) {
	content, err := json.Marshal(in.Content)
	if err != nil {
		return fmt.Errorf("failed to marshal log content: %w", err)
	}
	dataInsert := g.Map{
		dao.Log.Columns().OrgID:     in.OrgID,
		dao.Log.Columns().UserID:    in.UserID,
		dao.Log.Columns().UserName:  in.UserName,
		dao.Log.Columns().IP:        in.IP,
		dao.Log.Columns().Type:      in.Type,
		dao.Log.Columns().Content:   string(content),
		dao.Log.Columns().CreatedAt: in.CreatedAt,
	}
	_, err = dao.Log.Ctx(ctx).Insert(dataInsert)
	return
}

func (s *logicsLog) ListLoginLog(ctx context.Context, in *model.ListLogReq) (out *model.ListLogRes, err error) {
	out = new(model.ListLogRes)
	if in.PageNum == 0 {
		in.PageNum = 1
	}
	out.CurrentPage = in.PageNum
	if in.PageSize == 0 {
		in.PageSize = model.PageSize
	}

	m := dao.Log.Ctx(ctx).Where(dao.Log.Columns().OrgID, in.OrgID).Where(dao.Log.Columns().Type, model.LogTypeAuth)
	if in.StartTime != nil {
		m = m.WhereGTE(dao.Log.Columns().CreatedAt, in.StartTime)
	}
	if in.EndTime != nil {
		m = m.WhereLTE(dao.Log.Columns().CreatedAt, in.EndTime)
	}

	out.Total, err = m.Count()
	if err != nil {
		return nil, fmt.Errorf("failed to count log: %w", err)
	}

	var list []*entity.TLog
	err = m.Page(in.PageNum, in.PageSize).OrderDesc(dao.Log.Columns().CreatedAt).Scan(&list)
	if err != nil {
		return nil, fmt.Errorf("failed to get log list: %w", err)
	}

	var loginList = make([]*model.Log, len(list))
	for k, v := range list {
		loginList[k] = s.convertAuthLog(v)
	}
	out.List = loginList

	return
}

func (s *logicsLog) ListOperLog(ctx context.Context, in *model.ListLogReq) (out *model.ListLogRes, err error) {
	out = new(model.ListLogRes)
	if in.PageNum == 0 {
		in.PageNum = 1
	}
	out.CurrentPage = in.PageNum
	if in.PageSize == 0 {
		in.PageSize = model.PageSize
	}

	m := dao.Log.Ctx(ctx).Where(dao.Log.Columns().OrgID, in.OrgID).Where(dao.Log.Columns().Type, model.LogTypeOper)
	if in.StartTime != nil {
		m = m.WhereGTE(dao.Log.Columns().CreatedAt, in.StartTime)
	}
	if in.EndTime != nil {
		m = m.WhereLTE(dao.Log.Columns().CreatedAt, in.EndTime)
	}

	out.Total, err = m.Count()
	if err != nil {
		return nil, fmt.Errorf("failed to count log: %w", err)
	}

	var list []*entity.TLog
	err = m.Page(in.PageNum, in.PageSize).OrderDesc(dao.Log.Columns().CreatedAt).Scan(&list)
	if err != nil {
		return nil, fmt.Errorf("failed to get log list: %w", err)
	}

	var operList = make([]*model.Log, len(list))
	for k, v := range list {
		operList[k] = s.convertOperLog(v)
	}
	out.List = operList

	return
}

func (s *logicsLog) RecordLog(r *ghttp.Request) {
	vars, err := g.Cfg().Get(r.GetCtx(), "gfToken.excludePaths")
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	excludePaths := gconv.Strings(vars)
	if slices.Contains(excludePaths, r.RequestURI) {
		return
	}

	orgID, err := service.ContextService().GetOrgID(r.GetCtx())
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	userID, err := service.ContextService().GetUserID(r.GetCtx())
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	userName, err := service.ContextService().GetUserName(r.GetCtx())
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}

	var typ model.LogType
	var content interface{}
	if strings.HasSuffix(r.RequestURI, "login") || strings.HasSuffix(r.RequestURI, "loout") {
		typ = model.LogTypeAuth
		err := r.GetError()
		if err != nil {
			content = map[string]interface{}{"result": fmt.Sprintf("登录失败: %s", err.Error())}
		} else {
			content = map[string]interface{}{"result": "登录成功"}
		}
	} else if strings.HasSuffix(r.RequestURI, "registration") {
		typ = model.LogTypeAuth
		err := r.GetError()
		if err != nil {
			content = map[string]interface{}{"result": fmt.Sprintf("注册失败: %s", err.Error())}
		} else {
			content = map[string]interface{}{"result": "注册成功"}
		}
	} else {
		typ = model.LogTypeOper
		err := r.GetError()
		if err != nil {
			method := r.Method
			component := r.RequestURI
			result := fmt.Sprintf("操作失败：%s", err.Error())
			content = map[string]interface{}{"result": result, "method": method, "component": component}
		} else {
			method := r.Method
			component := r.RequestURI
			result := fmt.Sprintf("操作成功：%s", r.RequestURI)
			content = map[string]interface{}{"result": result, "method": method, "component": component}
		}
	}
	data := &model.Log{
		OrgID:    orgID,
		UserID:   userID,
		UserName: userName,
		IP:       libUtils.GetClientIp(r.GetCtx()),
		Type:     typ,
		Content:  content,
	}

	service.Log().InvokeLog(data)
}

func (s *logicsLog) convertAuthLog(in *entity.TLog) (out *model.Log) {
	out = &model.Log{
		ID:        in.ID,
		OrgID:     in.OrgID,
		UserID:    in.UserID,
		UserName:  in.UserName,
		IP:        in.IP,
		Type:      model.LogType(in.Type),
		CreatedAt: in.CreatedAt,
	}
	err := json.Unmarshal([]byte(in.Content), &out.Content)
	if err != nil {
		g.Log().Error(context.Background(), err)
	}

	return
}

func (s *logicsLog) convertOperLog(in *entity.TLog) (out *model.Log) {
	out = &model.Log{
		ID:        in.ID,
		OrgID:     in.OrgID,
		UserID:    in.UserID,
		UserName:  in.UserName,
		IP:        in.IP,
		Type:      model.LogType(in.Type),
		CreatedAt: in.CreatedAt,
	}
	err := json.Unmarshal([]byte(in.Content), &out.Content)
	if err != nil {
		g.Log().Error(context.Background(), err)
	}
	return
}
