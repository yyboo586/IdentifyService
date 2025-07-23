package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"context"
)

var Dept = sysDeptController{}

type sysDeptController struct {
	BaseController
}

// List 部门列表
func (c *sysDeptController) List(ctx context.Context, req *system.DeptSearchReq) (res *system.DeptSearchRes, err error) {
	res = new(system.DeptSearchRes)
	req.UserId = service.Context().GetUserId(ctx)
	req.UserDeptId = service.Context().GetDeptId(ctx)
	res.DeptList, err = service.SysDept().GetList(ctx, req)
	return
}

// Add 添加部门
func (c *sysDeptController) Add(ctx context.Context, req *system.DeptAddReq) (res *system.DeptAddRes, err error) {
	err = service.SysDept().Add(ctx, req)
	return
}

// Edit 修改部门
func (c *sysDeptController) Edit(ctx context.Context, req *system.DeptEditReq) (res *system.DeptEditRes, err error) {
	err = service.SysDept().Edit(ctx, req)
	return
}

// Delete 删除部门
func (c *sysDeptController) Delete(ctx context.Context, req *system.DeptDeleteReq) (res *system.DeptDeleteRes, err error) {
	err = service.SysDept().Delete(ctx, req.Id)
	return
}

// TreeSelect 获取部门数据结构数据
func (c *sysDeptController) TreeSelect(ctx context.Context, req *system.DeptTreeSelectReq) (res *system.DeptTreeSelectRes, err error) {
	var deptList []*entity.SysDept
	currentUser := service.Context().GetLoginUser(ctx)
	deptList, err = service.SysDept().GetList(ctx, &system.DeptSearchReq{
		Status:     "1", //正常状态数据
		ShowAll:    !req.ShowOwner,
		UserId:     currentUser.Id,
		UserDeptId: currentUser.DeptId,
	})
	if err != nil {
		return
	}
	res = new(system.DeptTreeSelectRes)
	topIds := service.SysDept().GetTopIds(deptList)
	for _, v := range topIds {
		res.Deps = append(res.Deps, service.SysDept().GetListTree(v, deptList)...)
	}
	return
}
