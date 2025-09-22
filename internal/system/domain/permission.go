package domain

import (
	"IdentifyService/internal/system/dao"
	"IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/model/entity"
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	permissionOnce     sync.Once
	permissionInstance *permission
)

type permission struct {
	enforcer         *casbin.SyncedEnforcer
	casbinUserPrefix string // CasBin 用户ID前缀
	casbinRolePrefix string // CasBin 角色ID前缀

	permissionPoints []*model.Permission
}

func NewPermission(casbinEnforcer *casbin.SyncedEnforcer) interfaces.IPermission {
	permissionOnce.Do(func() {
		permissionInstance = &permission{
			enforcer:         casbinEnforcer,
			casbinUserPrefix: "u_",
			casbinRolePrefix: "r_",

			permissionPoints: make([]*model.Permission, 0),
		}
	})
	return permissionInstance
}

func (p *permission) IsSuperAdmin(ctx context.Context, userID string) bool {
	return userID == model.DefaultSuperAdminID
}

func (p *permission) IsOrgAdmin(ctx context.Context, userID string, orgInfo *model.Org) bool {
	return orgInfo.ManagerID == userID
}

func (p *permission) GetAllPermissionPonits(ctx context.Context) (permissions []*model.Permission, err error) {
	return
}

func (p *permission) GetPermissionPointsByUserID(ctx context.Context, userID string) (permissionPointList []*model.Permission, err error) {
	return permissionPointList, nil
}

func (p *permission) GetPermissionPointsByRoleID(ctx context.Context, roleID int64) (permissionPointList []*model.Permission, err error) {
	return permissionPointList, nil
}

func (p *permission) FilterPermissionPointIDsByUserID(ctx context.Context, permissionPointIDs []int64, userID string) (out []int64, err error) {
	return permissionPointIDs, nil
}

func (p *permission) FilterRoleIDsByUserID(ctx context.Context, roleIDs []int64, userID string) (out []int64, err error) {
	return
}

func (p *permission) AssignUserRoles(ctx context.Context, userID string, roleIDs []int64) (err error) {
	for _, v := range roleIDs {
		_, err = p.enforcer.AddNamedGroupingPolicy("g", fmt.Sprintf("%s%s", p.casbinUserPrefix, userID), gconv.String(v))
		if err != nil {
			return
		}
	}
	return
}

// RemoveUserRoles 删除用户指定角色
func (p *permission) RemoveUserRoles(ctx context.Context, userID string, roleIDs []int64) (err error) {
	_, err = p.enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", p.casbinUserPrefix, userID))
	if err != nil {
		return
	}
	return
}

// RemoveUserAllRoles 删除用户全部角色
func (p *permission) RemoveUserAllRoles(ctx context.Context, userIDs []string) (err error) {
	for _, v := range userIDs {
		_, err = p.enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", p.casbinUserPrefix, v))
		if err != nil {
			return
		}
	}
	return
}

func (p *permission) FilterUserRoleIDs(ctx context.Context, userID string, roleIDs []int64) (out []int64, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		out = roleIDs
		return
	}

	roleInfos, err := roleInstance.GetRolesByUserID(ctx, userID)
	if err != nil {
		return
	}
	m := make(map[int64]bool, len(roleInfos))
	for _, v := range roleInfos {
		m[v.ID] = true
	}

	for _, v := range roleIDs {
		if m[v] {
			out = append(out, v)
		}
	}
	return
}

func (p *permission) GetRolesByUserID(ctx context.Context, userID string) (out []*model.Role, err error) {
	out, err = roleInstance.GetRolesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (p *permission) GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error) {
	roleIDs, err = roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return roleIDs, nil
}

func (p *permission) AssignRolePermissions(ctx context.Context, roleID int64, permissionIDs []int64) (err error) {
	ruleIdsStr := gconv.Strings(permissionIDs)
	rules := make([][]string, len(ruleIdsStr))

	for k, v := range ruleIdsStr {
		rules[k] = []string{gconv.String(roleID), v, "All"}
	}

	_, err = p.enforcer.AddNamedPolicies("p", rules)
	if err != nil {
		return
	}

	return
}

// RemoveRolePermissions 删除角色指定权限
func (p *permission) RemoveRolePermissions(ctx context.Context, roleID int64, permissionIDs []int64) (err error) {
	for _, v := range permissionIDs {
		// 同时匹配角色ID(v0)和权限ID(v1)
		_, err = p.enforcer.RemoveFilteredNamedPolicy("p", 1, gconv.String(roleID), gconv.String(v))
		if err != nil {
			return
		}
	}
	return
}

// RemoveRoleAllPermissions 删除角色全部权限
func (p *permission) RemoveRoleAllPermissions(ctx context.Context, roleIDs []int64) (err error) {
	for _, v := range roleIDs {
		_, err = p.enforcer.RemoveFilteredNamedPolicy("p", 0, gconv.String(v))
		if err != nil {
			return
		}
	}
	return
}

// TODO: 数据权限是否要在这里校验？该怎么校验
func (p *permission) HasPermission(ctx context.Context, permission model.PermissionPoint, userID string, orgInfo *model.Org) (hasPermission bool, err error) {
	// 超级管理员拥有所有权限

	if p.IsSuperAdmin(ctx, userID) {
		return true, nil
	}

	// 组织管理员拥有组织下所有权限
	if p.IsOrgAdmin(ctx, userID, orgInfo) {
		// TODO: 检查操作的资源ID是否属于该组织
	}

	/*
		hasPermission, err := service.Permission().HasPermission(ctx, permission)
		if err != nil {
			return false, err
		}
	*/
	return hasPermission, nil
}

// GetAllPermissions 获取所有权限点
func (p *permission) GetAllPermissions(ctx context.Context) (permissions []*model.Permission, err error) {
	var entities []*entity.Permission
	err = dao.Permission.Ctx(ctx).Scan(&entities)
	if err != nil {
		g.Log().Error(ctx, "获取权限点列表失败:", err)
		return
	}

	for _, e := range entities {
		permissions = append(permissions, p.convertPermissionEntityToModel(e))
	}
	return
}

func (p *permission) Init(ctx context.Context) (err error) {
	g.Log().Info(ctx, "开始初始化权限点...")

	// 初始化权限点
	for _, permissionPoint := range p.permissionPoints {
		// 检查权限点是否已存在
		exists, err := dao.Permission.Ctx(ctx).Where(dao.Permission.Columns().Code, permissionPoint.Code).Exist()
		if err != nil {
			g.Log().Error(ctx, "检查权限点是否存在失败:", err)
			return err
		}
		// 权限点已存在，跳过
		if exists {
			continue
		}

		// 创建权限点及其关联资源
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 创建资源
			resourceIDs := make([]int64, 0, len(permissionPoint.Resources))
			for _, resource := range permissionPoint.Resources {
				// 检查资源是否已存在
				resourceExists, err := dao.Resource.Ctx(ctx).Where(dao.Resource.Columns().Type, resource.Type).Where(dao.Resource.Columns().Code, resource.Code).Exist()
				if err != nil {
					return err
				}

				var resourceID int64
				if resourceExists {
					// 资源已存在，获取ID
					var resourceEntity entity.Resource
					err = dao.Resource.Ctx(ctx).Where(dao.Resource.Columns().Type, resource.Type).Where(dao.Resource.Columns().Code, resource.Code).Scan(&resourceEntity)
					if err != nil {
						return err
					}
					resourceID = resourceEntity.ID
				} else {
					// 创建新资源
					resourceID, err = dao.Resource.Ctx(ctx).Data(map[string]any{
						dao.Resource.Columns().Type: resource.Type,
						dao.Resource.Columns().Code: resource.Code,
					}).InsertAndGetId()
					if err != nil {
						if strings.Contains(err.Error(), "Duplicate entry") {
							// 并发情况下可能重复创建，重新查询
							var resourceEntity entity.Resource
							err = dao.Resource.Ctx(ctx).Where(dao.Resource.Columns().Type, resource.Type).Where(dao.Resource.Columns().Code, resource.Code).Scan(&resourceEntity)
							if err != nil {
								return err
							}
							resourceID = resourceEntity.ID
						} else {
							return err
						}
					}
				}
				resourceIDs = append(resourceIDs, resourceID)
			}

			// 创建权限点
			permissionID, err := dao.Permission.Ctx(ctx).Data(map[string]any{
				dao.Permission.Columns().Code:        permissionPoint.Code,
				dao.Permission.Columns().CodeName:    permissionPoint.CodeName,
				dao.Permission.Columns().Description: permissionPoint.Description,
			}).InsertAndGetId()
			if err != nil {
				if strings.Contains(err.Error(), "Duplicate entry") {
					// 权限点已存在，跳过
					return nil
				}
				return err
			}

			// 创建权限点与资源的关联（这里假设一个权限点对应一个资源，如果需要一对多关系需要调整）
			if len(resourceIDs) > 0 {
				// 使用第一个资源ID作为主要关联
				_, err = dao.Permission.Ctx(ctx).WherePri(permissionID).Data(map[string]any{
					dao.Permission.Columns().ResourceID: resourceIDs[0],
				}).Update()
				if err != nil {
					return err
				}
			}

			g.Log().Info(ctx, "权限点创建成功:", permissionPoint.Code)
			return nil
		})
		if err != nil {
			g.Log().Error(ctx, "创建权限点失败:", permissionPoint.Code, err)
			return err
		}
	}

	// 清空权限点列表
	p.permissionPoints = nil
	g.Log().Info(ctx, "权限点初始化完成")

	// 初始化角色
	err = p.initDefaultRoles(ctx)
	if err != nil {
		g.Log().Error(ctx, "初始化默认角色失败:", err)
		return err
	}

	return nil
}

func (p *permission) RegisterPermissionPoints(ctx context.Context, in []*model.Permission) {
	p.permissionPoints = append(p.permissionPoints, in...)
}

// initDefaultRoles 初始化默认角色
func (p *permission) initDefaultRoles(ctx context.Context) (err error) {
	g.Log().Info(ctx, "开始初始化默认角色...")

	// 获取所有权限点
	allPermissions, err := p.GetAllPermissions(ctx)
	if err != nil {
		g.Log().Error(ctx, "获取权限点列表失败:", err)
		return err
	}

	// 创建权限点映射，便于查找
	permissionMap := make(map[model.PermissionPoint]int64)
	for _, perm := range allPermissions {
		permissionMap[perm.Code] = perm.ID
	}

	// 初始化默认角色
	for _, defaultRole := range model.DefaultRoles {
		// 检查角色是否已存在
		exists, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().OrgID, defaultRole.OrgID).Where(dao.Role.Columns().Name, defaultRole.Name).Exist()
		if err != nil {
			g.Log().Error(ctx, "检查角色是否存在失败:", err)
			return err
		}

		// 角色已存在，跳过
		if exists {
			g.Log().Debug(ctx, "角色已存在，跳过:", defaultRole.Name)
			continue
		}

		// 创建角色
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 创建角色
			roleID, err := dao.Role.Ctx(ctx).Data(map[string]any{
				dao.Role.Columns().OrgID:     defaultRole.OrgID,
				dao.Role.Columns().PID:       0, // 默认角色都是根角色
				dao.Role.Columns().Name:      defaultRole.Name,
				dao.Role.Columns().CreatorID: defaultRole.CreatorID,
			}).InsertAndGetId()
			if err != nil {
				if strings.Contains(err.Error(), "Duplicate entry") {
					// 角色已存在，跳过
					return nil
				}
				return err
			}

			// 分配权限点
			if len(defaultRole.PermissionPoints) > 0 {
				permissionIDs := make([]int64, 0, len(defaultRole.PermissionPoints))
				for _, permissionPoint := range defaultRole.PermissionPoints {
					if permissionID, exists := permissionMap[permissionPoint]; exists {
						permissionIDs = append(permissionIDs, permissionID)
					} else {
						g.Log().Warning(ctx, "权限点不存在，跳过:", permissionPoint)
					}
				}

				// 分配权限到角色
				if len(permissionIDs) > 0 {
					err = p.AssignRolePermissions(ctx, roleID, permissionIDs)
					if err != nil {
						g.Log().Error(ctx, "分配角色权限失败:", err)
						return err
					}
				}
			}

			g.Log().Info(ctx, "角色创建成功:", defaultRole.Name)
			return nil
		})
		if err != nil {
			g.Log().Error(ctx, "创建角色失败:", defaultRole.Name, err)
			return err
		}
	}

	g.Log().Info(ctx, "默认角色初始化完成")
	return nil
}

// convertEntityToModel 转换实体到模型
func (p *permission) convertPermissionEntityToModel(entity *entity.Permission) *model.Permission {
	return &model.Permission{
		ID:          entity.ID,
		Code:        model.PermissionPoint(entity.Code),
		CodeName:    entity.CodeName,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (p *permission) convertResourceEntityToModel(entity *entity.Resource) *model.Resource {
	return &model.Resource{
		ID:        entity.ID,
		Type:      model.ResourceType(entity.Type),
		Code:      entity.Code,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
