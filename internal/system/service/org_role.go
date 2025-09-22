package service

import (
	"IdentifyService/api/v1/system"
	commonService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/model"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *orgService) CreateRole(ctx context.Context, req *system.RoleCreateReq) (res *system.RoleCreateRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.RoleCreate})
	if err != nil {
		return nil, err
	}

	// 检查组织是否启用
	_, err = c.domainOrg.IsEnabled(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	// 过滤权限点
	currentUserId, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	req.PermissionPointCodes, err = c.domainPermission.FilterPermissionPointsByUserID(ctx, req.PermissionPointCodes, currentUserId)
	if err != nil {
		return nil, err
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		res, err = c.domainRole.Create(ctx, tx, req)
		if err != nil {
			return err
		}

		if len(req.PermissionPointCodes) > 0 {
			err = c.domainPermission.AssignRolePermissions(ctx, res.ID, req.PermissionPointCodes)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *orgService) DeleteRole(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.RoleDelete})
	if err != nil {
		return nil, err
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = c.domainRole.DeleteByIDs(ctx, tx, req.RoleIDs)
		if err != nil {
			return err
		}

		err = c.domainPermission.RemoveRoleAllPermissions(ctx, req.RoleIDs)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *orgService) EditRole(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.RoleEdit})
	if err != nil {
		return nil, err
	}

	// 过滤权限点
	currentUserId, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	req.PermissionPointCodes, err = c.domainPermission.FilterPermissionPointsByUserID(ctx, req.PermissionPointCodes, currentUserId)
	if err != nil {
		return nil, err
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = c.domainRole.Edit(ctx, tx, req)
		if err != nil {
			return err
		}

		err = c.domainPermission.RemoveRoleAllPermissions(ctx, []int64{req.RoleID})
		if err != nil {
			return err
		}

		if len(req.PermissionPointCodes) > 0 {
			err = c.domainPermission.AssignRolePermissions(ctx, req.RoleID, req.PermissionPointCodes)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *orgService) GetRole(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.RoleView, model.RoleAssign})
	if err != nil {
		return nil, err
	}

	out, err := c.domainRole.Get(ctx, req.RoleID)
	if err != nil {
		return nil, err
	}

	permissions, err := c.domainPermission.GetPermissionPointsByRoleID(ctx, []int64{out.ID})
	if err != nil {
		return nil, err
	}

	res = new(system.RoleGetRes)
	res.Role = c.formatRoleInfo(out)
	res.Permissions = permissions

	return res, nil
}

func (c *orgService) GetRoleTrees(ctx context.Context, req *system.RoleTreeReq) (res *system.RoleTreeRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.RoleList})
	if err != nil {
		return nil, err
	}

	out, err := c.domainOrg.GetRoleTrees(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}

	return &system.RoleTreeRes{
		List: out,
	}, nil
}

func (c *orgService) formatRoleInfo(in *model.Role) (out *system.RoleInfo) {
	out = &system.RoleInfo{
		ID:        in.ID,
		PID:       in.PID,
		OrgID:     in.OrgID,
		Name:      in.Name,
		CreatorID: in.CreatorID,
		DeletorID: in.DeletorID,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		DeletedAt: in.DeletedAt,
	}
	return
}
