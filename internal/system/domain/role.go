package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"

	"IdentifyService/api/v1/system"
	commonService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/dao"
	service "IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/model/entity"
	"IdentifyService/library/libUtils"

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

func NewRole(enforcer *casbin.SyncedEnforcer) service.IRole {
	roleOnce.Do(func() {
		roleInstance = &role{
			enforcer: enforcer,
		}
	})
	return roleInstance
}

// Create 添加角色
func (r *role) Create(ctx context.Context, tx gdb.TX, req *system.RoleCreateReq) (res *system.RoleCreateRes, err error) {
	operID, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	// 1. 验证父角色是否存在（如果指定了父角色） && 验证父角色是否属于同一组织
	if req.Pid > 0 {
		parentRole, err := r.Get(ctx, req.Pid)
		if err != nil {
			return nil, fmt.Errorf("获取父角色信息失败: %w", err)
		}
		if parentRole.OrgID != req.OrgID {
			return nil, errors.New("父角色必须和子角色属于同一组织")
		}
	}

	// 2. 执行添加操作
	id, err := dao.Role.Ctx(ctx).TX(tx).InsertAndGetId(map[string]interface{}{
		dao.Role.Columns().OrgID:     req.OrgID,
		dao.Role.Columns().PID:       req.Pid,
		dao.Role.Columns().Name:      req.Name,
		dao.Role.Columns().CreatorID: operID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, fmt.Errorf("角色名称已存在: %w", err)
		}
		return nil, fmt.Errorf("添加角色失败: %w", err)
	}

	return &system.RoleCreateRes{
		ID: id,
	}, nil
}

func (r *role) DeleteByIDs(ctx context.Context, tx gdb.TX, roleIDs []int64) (err error) {
	// 检查每个角色是否可以删除
	if err = r.canDelete(ctx, roleIDs); err != nil {
		return fmt.Errorf("检查角色删除条件失败: %w", err)
	}

	// 执行删除操作
	if _, err = dao.Role.Ctx(ctx).TX(tx).WhereIn(dao.Role.Columns().ID, roleIDs).Delete(); err != nil {
		return fmt.Errorf("删除角色失败: %w", err)
	}

	return nil
}

// Edit 编辑角色
//
// TODO List
// 修改角色权限后，需要撤销令牌
func (r *role) Edit(ctx context.Context, tx gdb.TX, req *system.RoleEditReq) (err error) {
	// 验证父子角色关系
	err = r.validateRoleHierarchy(ctx, req.RoleID, req.PID)
	if err != nil {
		return
	}

	// 4. 执行更新操作
	_, err = dao.Role.Ctx(ctx).TX(tx).WherePri(req.RoleID).Data(map[string]interface{}{
		dao.Role.Columns().PID:  req.PID,
		dao.Role.Columns().Name: req.Name,
	}).Update()
	if err != nil {
		err = fmt.Errorf("修改角色失败: %w", err)
		return
	}

	return

}

func (r *role) Get(ctx context.Context, id int64) (res *model.Role, err error) {
	var entityRole entity.Role
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().ID, id).Scan(&entityRole)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("角色不存在: %w", err)
		}
		return nil, fmt.Errorf("获取角色信息失败: %w", err)
	}
	res = r.convertEntityToLogics(&entityRole)

	policy, _ := r.enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))

	res.PermissionPoints = make([]model.PermissionPoint, len(policy))
	for k, v := range policy {
		res.PermissionPoints[k] = model.PermissionPoint(v[1])
	}

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

/*
func (r *role) ListTreesByOrgID(ctx context.Context, orgID string) (out []*system.RoleNode, err error) {

			operatorInfo := service.ContextService().Get(ctx)
			if !userInstance.IsSuperAdmin(ctx, operatorInfo.User.ID) && operatorInfo.User.OrgID != orgID {
				err = errors.New("没有权限获取组织角色列表")
				g.Log().Error(ctx, err)
				return
			}

		var rootRoleIDs []*entity.Role
		err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().OrgID, orgID).Where(dao.Role.Columns().PID, 0).Scan(&rootRoleIDs)
		if err != nil {
			err = fmt.Errorf("获取组织角色列表失败: %w", err)
			g.Log().Error(ctx, err)
			return
		}

		for _, v := range rootRoleIDs {
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
*/
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

func (r *role) GetRolesByUserID(ctx context.Context, userID string) (out []*model.Role, err error) {
	roleIDs, err := r.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		return
	}
	for _, v := range roleIDs {
		item, err := r.Get(ctx, v)
		if err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return
}

func (r *role) FilterRoleIDs(ctx context.Context, roleIds []int64, userID string, includeChildren bool) (out []int64, err error) {
	currentUserId, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	if userInstance.IsSuperAdmin(ctx, currentUserId) {
		out = roleIds
		return
	}

	out = roleIds

	// 这里该怎么做呢？看起来需要根据孩子找到父亲
	if includeChildren {

	}

	return
}

// GetChildrenByRoleIDs 获取给定角色ID列表的所有后代（包括直接子角色、孙子角色等）
func (r *role) GetChildrenByRoleIDs(ctx context.Context, roleIDs []int64) (children []*model.Role, err error) {
	if len(roleIDs) == 0 {
		return
	}

	var entities []*entity.Role
	if err = dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().PID, roleIDs).Scan(&entities); err != nil {
		return nil, fmt.Errorf("获取子角色失败: %w", err)
	}
	if len(entities) == 0 {
		return
	}

	ids := make([]int64, 0, len(entities))
	for _, v := range entities {
		ids = append(ids, v.ID)
		children = append(children, r.convertEntityToLogics(v))
	}

	result, err := r.GetChildrenByRoleIDs(ctx, ids)
	if err != nil {
		return
	}
	children = append(children, result...)
	return
}

// HasUsers 检查角色是否有关联用户
func (r *role) HasUsers(ctx context.Context, roleID int64) (hasUsers bool, err error) {
	// 通过Casbin的分组策略检查用户角色关联
	// 格式：g, user_id, role_id
	groupPolicy, _ := r.enforcer.GetFilteredNamedGroupingPolicy("g", 1, gconv.String(roleID))

	return len(groupPolicy) > 0, nil
}

// validateRoleHierarchy 验证角色层次关系
// 防止创建循环引用和无效的父子关系
func (r *role) validateRoleHierarchy(ctx context.Context, roleID, parentID int64) error {
	// 如果父角色ID为0，表示是根角色，无需验证
	if parentID == 0 {
		return nil
	}

	// 不能将自己设为父角色
	if roleID == parentID {
		return errors.New("不能将角色设置为自己的父角色")
	}

	// 检查父角色是否存在
	parentRole, err := r.Get(ctx, parentID)
	if err != nil {
		return fmt.Errorf("获取父角色信息失败: %w", err)
	}
	if parentRole == nil {
		return errors.New("指定的父角色不存在")
	}

	// 获取当前角色信息
	currentRole, err := r.Get(ctx, roleID)
	if err != nil {
		return fmt.Errorf("获取当前角色信息失败: %w", err)
	}
	if currentRole == nil {
		return errors.New("当前角色不存在")
	}

	// 验证父角色是否属于同一组织
	if parentRole.OrgID != currentRole.OrgID {
		return errors.New("父角色必须属于同一组织")
	}

	// 检查是否会形成循环引用
	// 遍历父角色的所有祖先，确保当前角色不在其中
	return r.checkCircularReference(ctx, parentID, roleID)
}

// checkCircularReference 检查循环引用
// 递归检查父角色链，确保不会形成循环
func (r *role) checkCircularReference(ctx context.Context, parentID, targetRoleID int64) error {
	if parentID == 0 {
		return nil // 到达根节点，无循环
	}

	if parentID == targetRoleID {
		return errors.New("不能形成角色循环引用")
	}

	// 获取父角色的父角色
	parentRole, err := r.Get(ctx, parentID)
	if err != nil {
		return fmt.Errorf("检查循环引用时获取角色失败: %w", err)
	}
	if parentRole == nil {
		return nil // 父角色不存在，结束检查
	}

	// 递归检查上级
	return r.checkCircularReference(ctx, parentRole.PID, targetRoleID)
}

// canDelete 检查角色是否可以删除
// 删除规则：
// 1. 如果存在 角色的后代（包括直接子角色、孙子角色等）不在删除列表中，不允许删除该父角色
// 2. 如果存在 角色的关联用户，不允许删除
func (r *role) canDelete(ctx context.Context, roleIDs []int64) (err error) {
	// 获取所有要删除角色的后代（包括直接子角色、孙子角色等）
	allChildren, err := r.GetChildrenByRoleIDs(ctx, roleIDs)
	if err != nil {
		return fmt.Errorf("获取后代角色失败: %w", err)
	}

	// 创建删除角色的映射，用于快速查找
	deletedMap := make(map[int64]bool, len(roleIDs))
	for _, roleID := range roleIDs {
		deletedMap[roleID] = true
	}

	// 检查每个子角色是否也在删除列表中
	for _, child := range allChildren {
		if !deletedMap[child.ID] {
			// 如果子角色不在删除列表中，则不能删除其父角色
			parentRole, err := r.Get(ctx, child.PID)
			if err != nil {
				return fmt.Errorf("获取父角色失败: %w", err)
			}
			return fmt.Errorf("角色 %s 存在子角色 %s 不在删除列表中，请先删除子角色或将其一并删除", parentRole.Name, child.Name)
		}
	}

	// 检查每个角色是否有关联用户
	for _, roleID := range roleIDs {
		hasUsers, err := r.HasUsers(ctx, roleID)
		if err != nil {
			return fmt.Errorf("获取关联用户失败: %w", err)
		}
		if hasUsers {
			roleInfo, err := r.Get(ctx, roleID)
			if err != nil {
				return fmt.Errorf("获取角色失败: %w", err)
			}
			return fmt.Errorf("角色 %s 关联了用户，请先解除用户关联", roleInfo.Name)
		}
	}

	return nil
}

func (r *role) convertEntityToLogics(in *entity.Role) (out *model.Role) {
	out = &model.Role{
		ID:        in.ID,
		PID:       in.Pid,
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

func (r *role) convertModelToSystem(in *model.Role) (out *system.RoleInfo) {
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
