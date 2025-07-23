package controller

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/service"
)

type sysNoticeReadController struct {
	BaseController
}

var SysNoticeRead = new(sysNoticeReadController)

// List 列表
func (c *sysNoticeReadController) List(ctx context.Context, req *system.SysNoticeReadSearchReq) (res *system.SysNoticeReadSearchRes, err error) {
	res = new(system.SysNoticeReadSearchRes)
	res.SysNoticeReadSearchRes, err = service.SysNoticeRead().List(ctx, &req.SysNoticeReadSearchReq)
	return
}

// Get 获取已读记录
func (c *sysNoticeReadController) Get(ctx context.Context, req *system.SysNoticeReadGetReq) (res *system.SysNoticeReadGetRes, err error) {
	res = new(system.SysNoticeReadGetRes)
	res.SysNoticeReadInfoRes, err = service.SysNoticeRead().GetById(ctx, req.Id)
	return
}

// Add 添加已读记录
func (c *sysNoticeReadController) Add(ctx context.Context, req *system.SysNoticeReadAddReq) (res *system.SysNoticeReadAddRes, err error) {
	err = service.SysNoticeRead().Add(ctx, req.SysNoticeReadAddReq)
	return
}

// Add 添加已读记录
func (c *sysNoticeReadController) ReadNotice(ctx context.Context, req *system.SysNoticeReadNoticeAddReq) (res *system.SysNoticeReadNoticeAddRes, err error) {
	err = service.SysNoticeRead().ReadNotice(ctx, req.SysNoticeReadNoticeReq)
	return
}

// Edit 修改已读记录
func (c *sysNoticeReadController) Edit(ctx context.Context, req *system.SysNoticeReadEditReq) (res *system.SysNoticeReadEditRes, err error) {
	err = service.SysNoticeRead().Edit(ctx, req.SysNoticeReadEditReq)
	return
}

// Delete 删除已读记录
func (c *sysNoticeReadController) Delete(ctx context.Context, req *system.SysNoticeReadDeleteReq) (res *system.SysNoticeReadDeleteRes, err error) {
	err = service.SysNoticeRead().Delete(ctx, req.Ids)
	return
}
