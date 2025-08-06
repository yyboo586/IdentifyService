package logics

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"IdentifyService/library/liberr"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	roleOnce     sync.Once
	roleInstance *role
)

type role struct {
	enforcer *casbin.SyncedEnforcer
}

func NewRole() service.IRole {
	roleOnce.Do(func() {
		enforcer, err := service.CasbinEnforcer(context.Background())
		if err != nil {
			panic(err)
		}
		roleInstance = &role{
			enforcer: enforcer,
		}
	})
	return roleInstance
}

func (r *role) Add(ctx context.Context, req *system.RoleAddReq) (id int64, err error) {
	operatorInfo := service.ContextService().Get(ctx)

	req.MenuIDs, err = authRuleInstance.FilterRuleIDsByUserID(ctx, req.MenuIDs, operatorInfo.User.ID)
	if err != nil {
		err = fmt.Errorf("过滤规则失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err = dao.Role.Ctx(ctx).TX(tx).InsertAndGetId(map[string]interface{}{
			dao.Role.Columns().OrgID:     req.OrgID,
			dao.Role.Columns().PID:       req.Pid,
			dao.Role.Columns().Name:      req.Name,
			dao.Role.Columns().Status:    model.RoleStatusEnabled,
			dao.Role.Columns().CreatorID: operatorInfo.User.ID,
		})
		if err != nil {
			err = fmt.Errorf("添加角色失败: %w", err)
			return
		}

		if len(req.MenuIDs) > 0 {
			err = r.addRoleRule(id, req.MenuIDs)
			if err != nil {
				err = fmt.Errorf("添加角色权限失败: %w", err)
				return
			}
		}
		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (r *role) DeleteByIDs(ctx context.Context, roleIDs []int64) (err error) {
	for _, id := range roleIDs {
		if !r.hasManageAccess(ctx, id, false) {
			err = errors.New("没有删除这个角色的权限")
			g.Log().Error(ctx, err)
			return
		}
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.Role.Ctx(ctx).TX(tx).WhereIn(dao.Role.Columns().ID, roleIDs).Delete()
		if err != nil {
			err = fmt.Errorf("删除角色失败: %w", err)
			return
		}

		err = r.delRoleRule(roleIDs)
		if err != nil {
			err = fmt.Errorf("删除角色权限失败: %w", err)
			return
		}
		return
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

// TODO List
// 修改角色权限后，需要撤销令牌
func (r *role) Edit(ctx context.Context, req *system.RoleEditReq) (err error) {
	if !r.hasManageAccess(ctx, req.ID, true) {
		err = errors.New("没有修改这个角色的权限")
		g.Log().Error(ctx, err)
		return
	}

	operatorInfo := service.ContextService().Get(ctx)
	req.MenuIDs, err = authRuleInstance.FilterRuleIDsByUserID(ctx, req.MenuIDs, operatorInfo.User.ID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.Role.Ctx(ctx).TX(tx).WherePri(req.ID).Data(map[string]interface{}{
			dao.Role.Columns().PID:  req.Pid,
			dao.Role.Columns().Name: req.Name,
		}).Update()
		if err != nil {
			err = fmt.Errorf("修改角色失败: %w", err)
			return
		}

		err = r.delRoleRule([]int64{req.ID})
		if err != nil {
			err = fmt.Errorf("删除角色权限失败: %w", err)
			return
		}

		if len(req.MenuIDs) > 0 {
			err = r.addRoleRule(req.ID, req.MenuIDs)
			if err != nil {
				err = fmt.Errorf("添加角色权限失败: %w", err)
				return
			}
		}

		return
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (r *role) EditStatus(ctx context.Context, id int64, enabled bool) (err error) {
	dataUpdate := map[string]interface{}{
		dao.Role.Columns().Status: model.RoleStatusDisabled,
	}
	if enabled {
		dataUpdate[dao.Role.Columns().Status] = model.RoleStatusEnabled
	}

	_, err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().ID, id).Data(dataUpdate).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (r *role) Get(ctx context.Context, id int64) (res *model.Role, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查是否有查看此角色的权限
		if !r.hasManageAccess(ctx, id, true) {
			liberr.ErrIsNil(ctx, errors.New("没有查看这个角色的权限"), "没有查看这个角色的权限")
		}

		var entityRole entity.Role
		err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().ID, id).Scan(&entityRole)
		liberr.ErrIsNil(ctx, err, "获取角色信息失败")
		res = r.convertEntityToLogics(&entityRole)

		policy, _ := r.enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))
		res.MenuIDs = make([]int64, len(policy))
		for k, v := range policy {
			res.MenuIDs[k] = gconv.Int64(v[1])
		}
	})
	return
}

func (r *role) ListByOrgID(ctx context.Context, orgID string) (out []*model.Role, err error) {
	var list []*entity.Role
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().OrgID, orgID).Scan(&list)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		err = fmt.Errorf("获取角色数据失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range list {
		var item *model.Role
		item, err = r.Get(ctx, v.ID)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		out = append(out, item)
	}
	return
}

// 1、获取根节点
// 2、获取根节点下的所有子节点(递归)
func (r *role) GetTreeByRoleID(ctx context.Context, id int64) (out *system.RoleNode, err error) {
	var roleEntity entity.Role
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().ID, id).Scan(&roleEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		err = fmt.Errorf("获取角色数据失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	out = &system.RoleNode{
		RoleInfo: r.convertModelToSystem(r.convertEntityToLogics(&roleEntity)),
	}

	var children []*entity.Role
	err = dao.Role.Ctx(ctx).Fields(dao.Role.Columns().ID).Where(dao.Role.Columns().PID, id).Scan(&children)
	if err != nil {
		err = fmt.Errorf("获取角色子节点失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	if len(children) > 0 {
		out.Children = make([]*system.RoleNode, 0, len(children))
		for _, v := range children {
			var childNode *system.RoleNode
			childNode, err = r.GetTreeByRoleID(ctx, v.ID)
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			out.Children = append(out.Children, childNode)
		}
	}

	return
}

func (r *role) ListTreesByOrgID(ctx context.Context, orgID string) (out []*system.RoleNode, err error) {
	/*
		operatorInfo := service.ContextService().Get(ctx)
		if !userInstance.IsSuperAdmin(ctx, operatorInfo.User.ID) && operatorInfo.User.OrgID != orgID {
			err = errors.New("没有权限获取组织角色列表")
			g.Log().Error(ctx, err)
			return
		}
	*/
	var rootRoles []*entity.Role
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().OrgID, orgID).Where(dao.Role.Columns().PID, 0).Scan(&rootRoles)
	if err != nil {
		err = fmt.Errorf("获取组织角色列表失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range rootRoles {
		var node *system.RoleNode
		node, err = r.GetTreeByRoleID(ctx, v.ID)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		out = append(out, node)
	}
	return
}

func (r *role) GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error) {
	groupPolicy, _ := r.enforcer.GetFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", userInstance.casBinUserPrefix, userID))
	if len(groupPolicy) > 0 {
		roleIDs = make([]int64, len(groupPolicy))
		for k, v := range groupPolicy {
			roleIDs[k] = gconv.Int64(v[1])
		}

		roleIDs = libUtils.SliceUnique(roleIDs)
	}

	return
}

func (r *role) FilterRoleIDs(ctx context.Context, roleIds []int64, userID string, includeChildren bool) (out []int64, err error) {
	//directAccessRoleIDs, err := r.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		err = fmt.Errorf("获取用户角色ID失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	// 这里该怎么做呢？看起来需要根据孩子找到父亲
	if includeChildren {

	}

	return
}

// 私有方法

// hasManageAccess 检查是否有管理指定角色的权限
// 功能：判断当前用户是否有权限管理指定角色
// 权限逻辑：
// 1. 超级管理员可以管理所有角色
// 2. 普通用户只能管理自己创建的角色
// 3. 普通用户只能管理自己有权限的角色
func (r *role) hasManageAccess(ctx context.Context, roleId int64, includeChildren bool) bool {
	operatorInfo := service.ContextService().Get(ctx)

	// 超级管理员拥有所有权限
	if !service.User().IsSuperAdmin(ctx, operatorInfo.User.ID) {
		var (
			roleIds   []int64
			hasAccess bool
			err       error
			list      []*model.Role
		)

		// 获取角色列表
		list, err = r.ListByOrgID(ctx, operatorInfo.User.OrgID)
		if err != nil {
			g.Log().Error(ctx, err)
			return false
		}
		// 检查是否是当前用户创建的角色
		for _, v := range list {
			// 判断是否当前用户所建角色
			if roleId == v.ID && v.CreatorID == operatorInfo.User.ID {
				return true
			}
		}

		// 获取当前用户的角色ID列表
		roleIds, err = r.GetRoleIDsByUserID(ctx, operatorInfo.User.ID)
		if err != nil {
			g.Log().Error(ctx, err)
			return false
		}

		// 检查是否有权限管理该角色
		if len(roleIds) > 0 {
			for _, v := range roleIds {
				if v == int64(roleId) {
					hasAccess = true
					break
				}
			}
		}
		return hasAccess
	}
	return true
}

func (r *role) addRoleRule(roleID int64, ruleIDs []int64) (err error) {
	ruleIdsStr := gconv.Strings(ruleIDs)
	rules := make([][]string, len(ruleIdsStr))

	for k, v := range ruleIdsStr {
		rules[k] = []string{gconv.String(roleID), v, "All"}
	}

	_, err = r.enforcer.AddNamedPolicies("p", rules)
	if err != nil {
		return
	}

	return
}

func (r *role) delRoleRule(roleIDs []int64) (err error) {
	for _, v := range roleIDs {
		_, err = r.enforcer.RemoveFilteredPolicy(0, gconv.String(v))
		if err != nil {
			return
		}
	}
	return
}

func (r *role) convertEntityToLogics(in *entity.Role) (out *model.Role) {
	out = &model.Role{
		ID:        in.ID,
		PID:       in.Pid,
		OrgID:     in.OrgID,
		Name:      in.Name,
		Status:    model.RoleStatus(in.Status),
		CreatorID: in.CreatorID,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}

func (r *role) convertModelToSystem(in *model.Role) (out *system.RoleInfo) {
	out = &system.RoleInfo{
		ID:        in.ID,
		PID:       in.PID,
		OrgID:     in.OrgID,
		Name:      in.Name,
		Enabled:   in.Status == model.RoleStatusEnabled,
		CreatorID: in.CreatorID,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
