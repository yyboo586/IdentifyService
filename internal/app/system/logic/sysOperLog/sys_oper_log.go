package sysOperLog

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"IdentifyService/library/liberr"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOperateLog struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterOperateLog(New())
}

func New() service.IOperateLog {
	return &sOperateLog{
		Pool: grpool.New(100),
	}
}

// OperationLog 操作日志写入
func (s *sOperateLog) OperationLog(r *ghttp.Request) {
	userInfo := service.Context().GetLoginUser(r.GetCtx())
	if userInfo == nil {
		return
	}
	url := r.Request.URL //请求地址
	//获取菜单
	//获取地址对应的菜单id
	menuList, err := service.SysAuthRule().GetMenuList(r.GetCtx())
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	var menu *model.SysAuthRuleInfoRes
	path := gstr.TrimLeft(url.Path, "/")
	for _, m := range menuList {
		if gstr.Equal(m.Name, path) {
			menu = m
			break
		}
	}
	data := &model.SysOperLogAdd{
		User:         userInfo,
		Menu:         menu,
		Url:          url,
		Params:       r.GetMap(),
		Method:       r.Method,
		ClientIp:     libUtils.GetClientIp(r.GetCtx()),
		OperatorType: 1,
	}
	s.Invoke(gctx.New(), data)
}

func (s *sOperateLog) Invoke(ctx context.Context, data *model.SysOperLogAdd) {
	s.Pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		s.operationLogAdd(ctx, data)
	})
}

// OperationLogAdd 添加操作日志
func (s *sOperateLog) operationLogAdd(ctx context.Context, data *model.SysOperLogAdd) {
	menuTitle := ""
	if data.Menu != nil {
		menuTitle = data.Menu.Title
	}
	dept, err := service.SysDept().GetByDeptId(ctx, data.User.DeptId)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if dept == nil {
		dept = &entity.SysDept{}
	}
	insertData := map[string]interface{}{
		dao.SysOperLog.Columns().Title:         menuTitle,
		dao.SysOperLog.Columns().Method:        data.Url.Path,
		dao.SysOperLog.Columns().RequestMethod: data.Method,
		dao.SysOperLog.Columns().OperatorType:  data.OperatorType,
		dao.SysOperLog.Columns().OperName:      data.User.UserName,
		dao.SysOperLog.Columns().DeptName:      dept.DeptName,
		dao.SysOperLog.Columns().OperIp:        data.ClientIp,
		dao.SysOperLog.Columns().OperLocation:  libUtils.GetCityByIp(data.ClientIp),
		dao.SysOperLog.Columns().OperTime:      gtime.Now(),
		dao.SysOperLog.Columns().OperParam:     gstr.SubStrRune(gconv.String(data.Params), 0, 60000),
	}
	rawQuery := data.Url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	insertData[dao.SysOperLog.Columns().OperUrl] = data.Url.Path + rawQuery
	_, err = dao.SysOperLog.Ctx(ctx).Insert(insertData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *sOperateLog) List(ctx context.Context, req *system.SysOperLogSearchReq) (listRes *system.SysOperLogSearchRes, err error) {
	listRes = new(system.SysOperLogSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysOperLog.Ctx(ctx)
		if req.Title != "" {
			m = m.Where(dao.SysOperLog.Columns().Title+" = ?", req.Title)
		}
		if req.RequestMethod != "" {
			m = m.Where(dao.SysOperLog.Columns().RequestMethod+" = ?", req.RequestMethod)
		}
		if req.OperName != "" {
			m = m.Where(dao.SysOperLog.Columns().OperName+" like ?", "%"+req.OperName+"%")
		}
		if len(req.DateRange) != 0 {
			m = m.Where("oper_time >=? AND oper_time <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "oper_id DESC"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysOperLogInfoRes
		err = m.Fields(system.SysOperLogSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysOperLogListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysOperLogListRes{
				OperId:         v.OperId,
				Title:          v.Title,
				RequestMethod:  v.RequestMethod,
				OperName:       v.OperName,
				DeptName:       v.DeptName,
				LinkedDeptName: v.LinkedDeptName,
				OperUrl:        v.OperUrl,
				OperIp:         v.OperIp,
				OperLocation:   v.OperLocation,
				OperParam:      v.OperParam,
				OperTime:       v.OperTime,
			}
		}
	})
	return
}

func (s *sOperateLog) GetByOperId(ctx context.Context, operId uint64) (res *model.SysOperLogInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysOperLog.Ctx(ctx).WithAll().Where(dao.SysOperLog.Columns().OperId, operId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sOperateLog) DeleteByIds(ctx context.Context, ids []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysOperLog.Ctx(ctx).Delete("oper_id in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sOperateLog) ClearLog(ctx context.Context) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = g.DB().Ctx(ctx).Exec(ctx, "truncate "+dao.SysOperLog.Table())
		liberr.ErrIsNil(ctx, err, "清除失败")
	})
	return
}
