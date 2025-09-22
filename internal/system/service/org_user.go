package service

import (
	"context"
	"fmt"

	"IdentifyService/api/v1/system"

	commonService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

func (c *orgService) CreateUser(ctx context.Context, req *system.UserCreateReq) (res *system.UserCreateRes, err error) {
	checkPoints := []model.PermissionPointCode{model.UserCreate}
	if len(req.RoleIDs) > 0 {
		checkPoints = append(checkPoints, model.RoleAssign)
	}
	err = c.checkPermission(ctx, checkPoints)
	if err != nil {
		return nil, err
	}

	currentUserId, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	roleIDs, err := c.domainPermission.FilterRoleIDsByUserID(ctx, req.RoleIDs, currentUserId)
	if err != nil {
		return nil, err
	}

	userID := uuid.New().String()
	userInfo := &model.User{
		ID:       userID,
		Name:     req.UserName,
		Nickname: req.NickName,
		Password: req.Password,
		OrgID:    req.OrgID,
		Sex:      model.UserSex(req.Sex),
		Email:    req.Email,
		Avatar:   req.Avatar,
		Mobile:   req.Mobile,
		Address:  req.Address,
		Describe: req.Describe,
		IsAdmin:  req.IsAdmin == 1,
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if err = c.domainUser.Create(ctx, tx, userInfo); err != nil {
			return
		}

		if len(roleIDs) > 0 {
			if err = c.domainPermission.AssignUserRoles(ctx, userID, roleIDs); err != nil {
				return
			}
		}

		return
	})
	if err != nil {
		return nil, err
	}

	res = &system.UserCreateRes{
		ID: userID,
	}
	return
}

func (c *orgService) DeleteUser(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.UserDelete})
	if err != nil {
		return nil, err
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if err = c.domainUser.Delete(ctx, tx, req.IDs); err != nil {
			return
		}

		if err = c.domainPermission.RemoveUserAllRoles(ctx, req.IDs); err != nil {
			return
		}

		return
	})
	if err != nil {
		return nil, err
	}

	res = &system.UserDeleteRes{}
	return
}

func (c *orgService) EditUser(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.UserEdit})
	if err != nil {
		return nil, err
	}

	err = c.domainUser.Edit(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (c *orgService) EditUserRoles(ctx context.Context, req *system.EditUserRolesReq) (res *system.EditUserRolesRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.UserEdit, model.RoleAssign})
	if err != nil {
		return nil, err
	}

	currentUserId, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	roleIDs, err := c.domainPermission.FilterRoleIDsByUserID(ctx, req.RoleIDs, currentUserId)
	if err != nil {
		return nil, err
	}

	if len(roleIDs) <= 0 {
		err = fmt.Errorf("用户没有权限修改角色")
		return
	}

	err = c.domainPermission.RemoveUserAllRoles(ctx, []string{req.UserID})
	if err != nil {
		return nil, err
	}

	err = c.domainPermission.AssignUserRoles(ctx, req.UserID, roleIDs)
	if err != nil {
		return nil, err
	}

	return
}

func (c *orgService) ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.UserEdit})
	if err != nil {
		return nil, err
	}

	err = c.domainUser.ResetUserPwd(ctx, req)
	return
}

func (c *orgService) GetUserInfo(ctx context.Context, req *system.GetUserInfoReq) (res *system.GetUserInfoRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.UserView})
	if err != nil {
		return nil, err
	}

	user, err := c.domainUser.GetByID(ctx, req.UserID)
	if err != nil {
		return
	}

	roleIDs, err := c.domainPermission.GetRoleIDsByUserID(ctx, req.UserID)
	if err != nil {
		return
	}

	res = &system.GetUserInfoRes{
		User: c.format(user),
	}
	res.User.RoleIDs = roleIDs
	return
}

func (c *orgService) SearchUser(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error) {
	err = c.checkPermission(ctx, []model.PermissionPointCode{model.UserList})
	if err != nil {
		return nil, err
	}

	total, out, err := c.domainUser.Search(ctx, req)
	if err != nil {
		return
	}

	res = new(system.UserSearchRes)
	res.Total = gconv.Int(total)
	for _, v := range out {
		var roleIDs []int64
		roleIDs, err = c.domainPermission.GetRoleIDsByUserID(ctx, v.ID)
		if err != nil {
			return
		}

		item := c.format(v)
		item.RoleIDs = roleIDs
		res.List = append(res.List, item)
	}
	return
}

func (c *orgService) format(in *model.User) (out *system.User) {
	out = &system.User{
		ID:        in.ID,
		Name:      in.Name,
		Nickname:  in.Nickname,
		Mobile:    in.Mobile,
		Email:     in.Email,
		Enabled:   in.Status == model.UserStatusEnabled,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
