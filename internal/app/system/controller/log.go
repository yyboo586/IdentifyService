package controller

import (
	"IdentifyService/api/v1/system"
	commonService "IdentifyService/internal/app/common/service"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"context"
	"errors"

	"github.com/yyboo586/common/LogModule"
)

var Log = logController{}

type logController struct {
	BaseController
}

func (c *logController) ListUserLog(ctx context.Context, req *system.ListUserLogReq) (res *system.ListUserLogRes, err error) {
	operatorInfo := service.Context().GetLoginUser(ctx)
	if operatorInfo == nil {
		return nil, errors.New("用户未登录")
	}

	res = new(system.ListUserLogRes)
	for _, action := range []LogModule.LogAction{
		model.LogActionUserLogin,
		model.LogActionUserLogout,
		model.LogActionUserUnRegister,
	} {
		logs, err := commonService.Log().ListLog(ctx, &LogModule.LogListFilter{
			Module:     model.LogModuleUser,
			Action:     action,
			OperatorID: operatorInfo.Id,
			Page:       req.PageReq.Page,
			Size:       req.PageReq.Size,
		})
		if err != nil {
			return nil, err
		}

		for _, log := range logs {
			res.List = append(res.List, convertLogItem(log))
		}
	}
	return res, nil
}

func convertLogItem(log *LogModule.LogItem) *system.LogItem {
	return &system.LogItem{
		Action:     model.GetLogActionName(log.Action),
		Message:    log.Message,
		Detail:     log.Detail,
		OperatorID: log.OperatorID,
		IP:         log.IP,
		CreateTime: log.CreateTime,
	}
}
