package controller

/*


import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"context"
	"fmt"
)

var OrgImprovedController = orgImprovedController{}

type orgImprovedController struct {
}

// Create 创建组织
func (c *orgImprovedController) Create(ctx context.Context, req *system.OrgCreateReq) (res *system.OrgCreateRes, err error) {
	// 参数验证
	if err = c.validateCreateParams(req); err != nil {
		return nil, err
	}

	in := &model.Org{
		PID:         req.PID,
		Name:        req.Name,
		Code:        req.Code,
		ManagerID:   req.ManagerID,
		ManagerName: req.ManagerName,
		Description: req.Description,
		Sort:        req.Sort,
	}

	orgID, err := service.Org().Create(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("创建组织失败: %w", err)
	}

	res = &system.OrgCreateRes{
		ID: orgID,
	}
	return
}

// Get 获取组织详情
func (c *orgImprovedController) Get(ctx context.Context, req *system.OrgGetReq) (res *system.OrgGetRes, err error) {
	out, err := service.Org().GetDetail(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("获取组织详情失败: %w", err)
	}

	res = &system.OrgGetRes{
		OrgDetailInfo: out,
	}
	return
}

// Update 更新组织信息
func (c *orgImprovedController) Update(ctx context.Context, req *system.OrgUpdateReq) (res *system.OrgUpdateRes, err error) {
	// 参数验证
	if err = c.validateUpdateParams(req); err != nil {
		return nil, err
	}

	in := &model.Org{
		ID:          req.ID,
		Name:        req.Name,
		Code:        req.Code,
		ManagerID:   req.ManagerID,
		ManagerName: req.ManagerName,
		Description: req.Description,
		Sort:        req.Sort,
	}

	err = service.Org().Update(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("更新组织信息失败: %w", err)
	}

	res = &system.OrgUpdateRes{}
	return
}

// Delete 删除组织
func (c *orgImprovedController) Delete(ctx context.Context, req *system.OrgDeleteReq) (res *system.OrgDeleteRes, err error) {
	err = service.Org().Delete(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("删除组织失败: %w", err)
	}

	res = &system.OrgDeleteRes{}
	return
}

// UpdateStatus 更新组织状态
func (c *orgImprovedController) UpdateStatus(ctx context.Context, req *system.OrgStatusUpdateReq) (res *system.OrgStatusUpdateRes, err error) {
	err = service.Org().UpdateStatus(ctx, req.ID, req.Enabled)
	if err != nil {
		return nil, fmt.Errorf("更新组织状态失败: %w", err)
	}

	res = &system.OrgStatusUpdateRes{}
	return
}

// List 获取组织列表
func (c *orgImprovedController) List(ctx context.Context, req *system.OrgListReq) (res *system.OrgListRes, err error) {
	// 参数验证和默认值设置
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 || req.PageSize > 100 {
		req.PageSize = 20
	}

	list, total, err := service.Org().List(ctx, &model.OrgListParams{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		PID:      req.PID,
		Enabled:  req.Enabled,
		SortBy:   req.SortBy,
		SortDesc: req.SortDesc,
	})
	if err != nil {
		return nil, fmt.Errorf("获取组织列表失败: %w", err)
	}

	res = &system.OrgListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return
}

// GetTree 获取组织树
func (c *orgImprovedController) GetTree(ctx context.Context, req *system.OrgTreeReq) (res *system.OrgTreeRes, err error) {
	list, err := service.Org().GetTree(ctx, &model.OrgTreeParams{
		RootID:   req.RootID,
		MaxDepth: req.MaxDepth,
		Enabled:  req.Enabled,
	})
	if err != nil {
		return nil, fmt.Errorf("获取组织树失败: %w", err)
	}

	res = &system.OrgTreeRes{
		List: list,
	}
	return
}

// GetSubTree 获取子组织树
func (c *orgImprovedController) GetSubTree(ctx context.Context, req *system.OrgSubTreeReq) (res *system.OrgSubTreeRes, err error) {
	tree, err := service.Org().GetSubTree(ctx, req.ID, &model.OrgTreeParams{
		MaxDepth: req.MaxDepth,
		Enabled:  req.Enabled,
	})
	if err != nil {
		return nil, fmt.Errorf("获取子组织树失败: %w", err)
	}

	res = &system.OrgSubTreeRes{
		OrgTreeNode: tree,
	}
	return
}

// Move 移动组织
func (c *orgImprovedController) Move(ctx context.Context, req *system.OrgMoveReq) (res *system.OrgMoveRes, err error) {
	// 参数验证
	if req.ID == req.NewPID {
		return nil, fmt.Errorf("不能将组织移动到自身")
	}

	err = service.Org().Move(ctx, req.ID, req.NewPID)
	if err != nil {
		return nil, fmt.Errorf("移动组织失败: %w", err)
	}

	res = &system.OrgMoveRes{}
	return
}

// Copy 复制组织
func (c *orgImprovedController) Copy(ctx context.Context, req *system.OrgCopyReq) (res *system.OrgCopyRes, err error) {
	// 参数验证
	if err = c.validateCopyParams(req); err != nil {
		return nil, err
	}

	newOrgID, err := service.Org().Copy(ctx, &model.OrgCopyParams{
		ID:        req.ID,
		NewPID:    req.NewPID,
		NewName:   req.NewName,
		CopyUsers: req.CopyUsers,
		CopyRoles: req.CopyRoles,
		CopyPerms: req.CopyPerms,
	})
	if err != nil {
		return nil, fmt.Errorf("复制组织失败: %w", err)
	}

	res = &system.OrgCopyRes{
		ID: newOrgID,
	}
	return
}

// GetStats 获取组织统计信息
func (c *orgImprovedController) GetStats(ctx context.Context, req *system.OrgStatsReq) (res *system.OrgStatsRes, err error) {
	stats, err := service.Org().GetStats(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("获取组织统计信息失败: %w", err)
	}

	res = &system.OrgStatsRes{
		OrgStatsInfo: stats,
	}
	return
}

// BatchDelete 批量删除组织
func (c *orgImprovedController) BatchDelete(ctx context.Context, req *system.OrgBatchDeleteReq) (res *system.OrgBatchDeleteRes, err error) {
	// 参数验证
	if len(req.IDs) == 0 {
		return nil, fmt.Errorf("组织ID列表不能为空")
	}
	if len(req.IDs) > 100 {
		return nil, fmt.Errorf("批量删除数量不能超过100个")
	}

	result, err := service.Org().BatchDelete(ctx, req.IDs)
	if err != nil {
		return nil, fmt.Errorf("批量删除组织失败: %w", err)
	}

	res = &system.OrgBatchDeleteRes{
		SuccessCount: result.SuccessCount,
		FailedIDs:    result.FailedIDs,
	}
	return
}

// BatchUpdateStatus 批量更新组织状态
func (c *orgImprovedController) BatchUpdateStatus(ctx context.Context, req *system.OrgBatchStatusUpdateReq) (res *system.OrgBatchStatusUpdateRes, err error) {
	// 参数验证
	if len(req.IDs) == 0 {
		return nil, fmt.Errorf("组织ID列表不能为空")
	}
	if len(req.IDs) > 100 {
		return nil, fmt.Errorf("批量更新数量不能超过100个")
	}

	result, err := service.Org().BatchUpdateStatus(ctx, req.IDs, req.Enabled)
	if err != nil {
		return nil, fmt.Errorf("批量更新组织状态失败: %w", err)
	}

	res = &system.OrgBatchStatusUpdateRes{
		SuccessCount: result.SuccessCount,
		FailedIDs:    result.FailedIDs,
	}
	return
}

// ==================== 私有方法 ====================

// validateCreateParams 验证创建参数
func (c *orgImprovedController) validateCreateParams(req *system.OrgCreateReq) error {
	if req.Name == "" {
		return fmt.Errorf("组织名称不能为空")
	}
	if req.PID == "" {
		return fmt.Errorf("父级ID不能为空")
	}
	if req.Code != "" {
		// 验证组织编码唯一性
		exists, err := service.Org().CheckCodeExists(ctx, req.Code)
		if err != nil {
			return fmt.Errorf("验证组织编码失败: %w", err)
		}
		if exists {
			return fmt.Errorf("组织编码已存在")
		}
	}
	return nil
}

// validateUpdateParams 验证更新参数
func (c *orgImprovedController) validateUpdateParams(req *system.OrgUpdateReq) error {
	if req.ID == "" {
		return fmt.Errorf("组织ID不能为空")
	}
	if req.Code != "" {
		// 验证组织编码唯一性（排除自身）
		exists, err := service.Org().CheckCodeExistsExcludeSelf(ctx, req.Code, req.ID)
		if err != nil {
			return fmt.Errorf("验证组织编码失败: %w", err)
		}
		if exists {
			return fmt.Errorf("组织编码已存在")
		}
	}
	return nil
}

// validateCopyParams 验证复制参数
func (c *orgImprovedController) validateCopyParams(req *system.OrgCopyReq) error {
	if req.ID == "" {
		return fmt.Errorf("源组织ID不能为空")
	}
	if req.NewPID == "" {
		return fmt.Errorf("目标父级ID不能为空")
	}
	if req.NewName == "" {
		return fmt.Errorf("新组织名称不能为空")
	}
	if req.ID == req.NewPID {
		return fmt.Errorf("不能将组织复制到自身")
	}
	return nil
}

*/
