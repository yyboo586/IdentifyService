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
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

/*
TODO List
1、角色支持多组织架构
*/

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
	if !userInstance.IsSuperAdmin(ctx, operatorInfo.User.ID) {
		// 过滤权限ID，确保只能分配自己有权限的菜单
		var roleIDs []int64
		roleIDs, err = r.GetRoleIDsByUserID(ctx, operatorInfo.User.ID)
		if err != nil {
			err = fmt.Errorf("获取角色ID失败: %w", err)
			g.Log().Error(ctx, err)
			return
		}
		req.MenuIDs, err = authRuleInstance.filterRuleIDsByRoleIDs(ctx, req.MenuIDs, roleIDs)
		if err != nil {
			err = fmt.Errorf("过滤规则失败: %w", err)
			g.Log().Error(ctx, err)
			return
		}
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		roleId, err := dao.Role.Ctx(ctx).TX(tx).InsertAndGetId(map[string]interface{}{
			dao.Role.Columns().OrgID:     req.OrgID,
			dao.Role.Columns().Pid:       req.Pid,
			dao.Role.Columns().Name:      req.Name,
			dao.Role.Columns().Status:    model.RoleStatusEnabled,
			dao.Role.Columns().CreatorID: operatorInfo.User.ID,
		})
		if err != nil {
			err = fmt.Errorf("添加角色失败: %w", err)
			return
		}

		if len(req.MenuIDs) > 0 {
			err = r.addRoleRule(roleId, req.MenuIDs)
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
	if !userInstance.IsSuperAdmin(ctx, operatorInfo.User.ID) {
		// 过滤权限ID，确保只能分配自己有权限的菜单
		var roleIDs []int64
		roleIDs, err = r.GetRoleIDsByUserID(ctx, operatorInfo.User.ID)
		if err != nil {
			err = fmt.Errorf("获取角色ID失败: %w", err)
			g.Log().Error(ctx, err)
			return
		}
		req.MenuIDs, err = authRuleInstance.filterRuleIDsByRoleIDs(ctx, req.MenuIDs, roleIDs)
		if err != nil {
			err = fmt.Errorf("过滤规则失败: %w", err)
			g.Log().Error(ctx, err)
			return
		}
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.Role.Ctx(ctx).TX(tx).WherePri(req.ID).Data(map[string]interface{}{
			dao.Role.Columns().Pid:    req.Pid,
			dao.Role.Columns().Status: req.Status,
			dao.Role.Columns().Name:   req.Name,
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

		err = r.addRoleRule(req.ID, req.MenuIDs)
		if err != nil {
			err = fmt.Errorf("添加角色权限失败: %w", err)
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

// Get 获取单个角色信息
func (r *role) Get(ctx context.Context, id int64) (res *model.Role, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查是否有查看此角色的权限
		if !r.hasManageAccess(ctx, id, true) {
			liberr.ErrIsNil(ctx, errors.New("没有查看这个角色的权限"), "没有查看这个角色的权限")
		}

		res = new(model.Role)
		err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().ID, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取角色信息失败")
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
		out = append(out, r.convertEntityToLogics(v))
	}
	return
}

// GetByName 根据角色名称获取角色信息
func (r *role) GetByName(ctx context.Context, roleName string) (out *model.Role, err error) {
	var entityRole entity.Role
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Name, roleName).Scan(&entityRole)
	if err != nil {
		err = fmt.Errorf("获取角色信息失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	out = r.convertEntityToLogics(&entityRole)
	return
}

// GetFilteredNamedPolicy 获取角色关联的菜单规则
// 功能：从Casbin获取角色拥有的所有菜单权限ID
//
// GetFilteredNamedPolicy 方法说明：
// - 参数1："p" 表示权限策略类型
// - 参数2：字段索引，0表示第一个字段（角色ID）
// - 参数3：字段值，要查询的角色ID
// - 返回：该角色的所有权限策略列表
func (r *role) GetFilteredNamedPolicy(ctx context.Context, id int64) (gpSlice []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 获取Casbin权限管理器
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)

		// GetFilteredNamedPolicy 获取过滤的命名策略
		// 参数说明：
		// - "p": 策略类型，表示权限策略
		// - 0: 字段索引，0表示第一个字段（角色ID）
		// - gconv.String(id): 角色ID（字符串格式）
		//
		// 返回格式：[[角色ID, 资源ID, 操作], [角色ID, 资源ID, 操作], ...]
		// 示例：[["1", "100", "All"], ["1", "200", "All"]]
		gp, _ := enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))
		gpSlice = make([]int64, len(gp))

		// 提取权限ID（第二个字段）
		for k, v := range gp {
			gpSlice[k] = gconv.Int64(v[1]) // v[1]是权限ID（资源ID）
		}
	})
	return
}

func (r *role) GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error) {
	// 查询用户关联的角色规则
	// 使用GetFilteredGroupingPolicy查询用户-角色关系
	// 格式：v0=用户key, v1=角色ID
	groupPolicy, _ := r.enforcer.GetFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", userInstance.casBinUserPrefix, userID))

	// 如果找到角色策略，则解析角色ID
	if len(groupPolicy) > 0 {
		roleIDs = make([]int64, len(groupPolicy))
		// 遍历每个策略，提取角色ID
		for k, v := range groupPolicy {
			// policy ["u_00000000-0000-0000-0000-000000000001","1"]
			roleIDs[k] = gconv.Int64(v[1])
		}
	}

	// 对角色ID列表进行去重处理
	roleIDs = libUtils.SliceUnique(roleIDs)

	return
}

// FindSonByParentId 根据父角色ID查找所有子角色
// 功能：递归查找角色的所有子角色（树形结构）
func (r *role) FindSonByParentId(roleList []*model.Role, id int64) []*model.Role {
	children := make([]*model.Role, 0, len(roleList))
	for _, v := range roleList {
		if v.PID == id {
			children = append(children, v)
			// 递归查找子角色的子角色
			fChildren := r.FindSonByParentId(roleList, v.ID)
			children = append(children, fChildren...)
		}
	}
	return children
}

// FindSonIDsByParentID 根据父角色ID查找所有子角色ID
// 功能：递归查找角色的所有子角色ID（树形结构）
func (r *role) FindSonIDsByParentID(roleList []*model.Role, id int64) []int64 {
	children := make([]int64, 0, len(roleList))
	for _, v := range roleList {
		if v.PID == id {
			children = append(children, v.ID)
			// 递归查找子角色的子角色ID
			fChildren := r.FindSonIDsByParentID(roleList, v.ID)
			children = append(children, fChildren...)
		}
	}
	return children
}

func (r *role) FilterRoleIDs(ctx context.Context, roleIds []int64, userID string, includeChildren bool) (out []int64, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		out = roleIds
		return
	}

	operatorInfo := service.ContextService().Get(ctx)
	var (
		accessRoleIDsList []int64
		accessRolesList   []*model.Role
	)
	accessRoleIDsList, err = r.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	accessRolesList, err = r.ListByOrgID(ctx, operatorInfo.User.OrgID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 子角色也要能够被授权
	sonIds := make([]int64, 0, 10)
	for _, v := range accessRoleIDsList {
		sonIds = append(sonIds, r.FindSonIDsByParentID(accessRolesList, v)...)
	}
	accessRoleIDsList = append(accessRoleIDsList, sonIds...)
	// 自己创建的角色可以被授权
	for _, v := range accessRolesList {
		if v.CreatorID == userID {
			accessRoleIDsList = append(accessRoleIDsList, v.ID)
		}
	}
	//去重accessRoleList
	accessRoleIDsList = gconv.Int64s(garray.NewArrayFrom(gconv.Interfaces(accessRoleIDsList)).Unique().Slice())
	for _, r := range roleIds {
		for _, a := range accessRoleIDsList {
			if r == a {
				out = append(out, r)
				break
			}
		}
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
		g.Log().Info(ctx, "[DEBUG] rolelist: ", "list", list)

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
