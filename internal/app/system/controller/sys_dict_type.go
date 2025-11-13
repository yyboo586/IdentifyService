package controller

import (
	"context"

	"IdentifyService/api/v1/system"
	commonService "IdentifyService/internal/app/common/service"

	"github.com/yyboo586/common/MiddleWare"
)

var DictType = &SysDictTypeController{}

type SysDictTypeController struct {
}

// List 字典类型列表
func (c *SysDictTypeController) List(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error) {
	res, err = commonService.SysDictType().List(ctx, req)
	return
}

// Add 添加字典类型
func (c *SysDictTypeController) Add(ctx context.Context, req *system.DictTypeAddReq) (res *system.DictTypeAddRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	err = commonService.SysDictType().Add(ctx, req, operator.UserID)
	return
}

// Get 获取字典类型
func (c *SysDictTypeController) Get(ctx context.Context, req *system.DictTypeGetReq) (res *system.DictTypeGetRes, err error) {
	res = new(system.DictTypeGetRes)
	res.DictType, err = commonService.SysDictType().Get(ctx, req)
	return
}

// Edit 修改字典数据
func (c *SysDictTypeController) Edit(ctx context.Context, req *system.DictTypeEditReq) (res *system.DictTypeEditRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	err = commonService.SysDictType().Edit(ctx, req, operator.UserID)
	return
}

func (c *SysDictTypeController) Delete(ctx context.Context, req *system.DictTypeDeleteReq) (res *system.DictTypeDeleteRes, err error) {
	err = commonService.SysDictType().Delete(ctx, req.DictIds)
	return
}

// OptionSelect 获取字典选择框列表
func (c *SysDictTypeController) OptionSelect(ctx context.Context, req *system.DictTypeAllReq) (res *system.DictTYpeAllRes, err error) {
	res = new(system.DictTYpeAllRes)
	res.DictType, err = commonService.SysDictType().GetAllDictType(ctx)
	return
}
