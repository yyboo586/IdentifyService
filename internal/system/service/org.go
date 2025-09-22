package service

import (
	"IdentifyService/api/v1/system"
	commonService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"
	"context"
	"fmt"

	"github.com/google/uuid"
)

var OrgService = orgService{}

type orgService struct {
	domainUser       interfaces.IUser
	domainOrg        interfaces.IOrg
	domainRole       interfaces.IRole
	domainPermission interfaces.IPermission
}

func NewOrgService(domainUser interfaces.IUser, domainOrg interfaces.IOrg, domainRole interfaces.IRole, domainPermission interfaces.IPermission) *orgService {
	return &orgService{
		domainUser:       domainUser,
		domainOrg:        domainOrg,
		domainRole:       domainRole,
		domainPermission: domainPermission,
	}
}

func (c *orgService) CreateOrg(ctx context.Context, req *system.OrgCreateReq) (res *system.OrgCreateRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPoint{model.OrgCreate})
	if err != nil {
		return nil, err
	}

	in := &model.Org{
		ID:          uuid.New().String(),
		PID:         req.PID,
		Name:        req.Name,
		ManagerID:   req.ManagerID,
		ManagerName: req.ManagerName,
	}
	err = c.domainOrg.Create(ctx, nil, in)
	if err != nil {
		return nil, err
	}

	return &system.OrgCreateRes{
		ID: in.ID,
	}, nil
}

func (c *orgService) DeleteOrg(ctx context.Context, req *system.OrgDeleteReq) (res *system.OrgDeleteRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPoint{model.OrgDelete})
	if err != nil {
		return nil, err
	}

	err = c.domainOrg.Delete(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	return &system.OrgDeleteRes{}, nil
}

func (c *orgService) EditOrg(ctx context.Context, req *system.OrgEdiReq) (res *system.OrgEditRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPoint{model.OrgEdit})
	if err != nil {
		return nil, err
	}

	in := &model.Org{
		ID:          req.OrgID,
		Name:        req.Name,
		ManagerID:   req.ManagerID,
		ManagerName: req.ManagerName,
		Status:      model.OrgStatus(req.Status),
	}
	err = c.domainOrg.Edit(ctx, in)
	if err != nil {
		return nil, err
	}

	return &system.OrgEditRes{}, nil
}

func (c *orgService) GetOrg(ctx context.Context, req *system.OrgGetReq) (res *system.OrgGetRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPoint{model.OrgView})
	if err != nil {
		return nil, err
	}

	out, err := c.domainOrg.Get(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	return &system.OrgGetRes{
		OrgInfo: c.convertOrg(out),
	}, nil
}

func (c *orgService) GetTreeOrg(ctx context.Context, req *system.OrgGetTreeReq) (res *system.OrgGetTreeRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPoint{model.OrgView})
	if err != nil {
		return nil, err
	}

	out, err := c.domainOrg.GetTree(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	return &system.OrgGetTreeRes{
		OrgTreeNode: out,
	}, nil
}

func (c *orgService) checkPermission(ctx context.Context, permissionPoints []model.PermissionPoint) (err error) {
	userID, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("权限检查失败: %w", err)
	}
	orgID, err := commonService.ContextService().GetOrgID(ctx)
	if err != nil {
		return fmt.Errorf("权限检查失败: %w", err)
	}
	orgInfo, err := c.domainOrg.Get(ctx, orgID)
	if err != nil {
		return fmt.Errorf("权限检查失败: %w", err)
	}

	for _, permission := range permissionPoints {
		hasPermission, err := c.domainPermission.HasPermission(ctx, permission, userID, orgInfo)
		if err != nil {
			return fmt.Errorf("权限检查失败: %w", err)
		}
		if !hasPermission {
			return fmt.Errorf("没有权限执行操作")
		}
	}
	return nil
}

func (c *orgService) convertOrg(in *model.Org) (out *system.OrgInfo) {
	out = &system.OrgInfo{
		ID:          in.ID,
		PID:         in.PID,
		Name:        in.Name,
		ManagerID:   in.ManagerID,
		ManagerName: in.ManagerName,
		Enabled:     in.Status == model.OrgStatusEnabled,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return
}
