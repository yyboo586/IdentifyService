package domain

import (
	"IdentifyService/internal/system/dao"
	"IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/model/entity"
	"IdentifyService/library/libUtils"
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// TODO: 权限点的存储，应该还能优化
/*
| p     | r_10                                   | role:create | All |    |    |    |
| p     | r_10                                   | role:edit   | All |    |    |    |
| p     | r_10                                   | role:delete | All |    |    |    |
| p     | r_10                                   | role:view   | All |    |    |    |
| p     | r_10                                   | role:list   | All |    |    |    |
| p     | r_10                                   | role:assign | All |    |    |    |

*/
var (
	permissionOnce     sync.Once
	permissionInstance *permission
)

type permission struct {
	casbinEnforcer   *casbin.SyncedEnforcer
	casbinUserPrefix string // CasBin 用户ID前缀
	casbinRolePrefix string // CasBin 角色ID前缀
}

func NewPermission(casbinEnforcer *casbin.SyncedEnforcer) interfaces.IPermission {
	permissionOnce.Do(func() {
		permissionInstance = &permission{
			casbinEnforcer:   casbinEnforcer,
			casbinUserPrefix: "u_",
			casbinRolePrefix: "r_",
		}
	})
	return permissionInstance
}

// checked
// AssignUserRoles 给用户分配角色
func (p *permission) AssignUserRoles(ctx context.Context, userID string, roleIDs []int64) (err error) {
	for _, v := range roleIDs {
		_, err = p.casbinEnforcer.AddNamedGroupingPolicy("g", fmt.Sprintf("%s%s", p.casbinUserPrefix, userID), gconv.String(v))
		if err != nil {
			return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
		}
	}
	return nil
}

// checked
// RemoveUserRoles 删除用户指定角色
func (p *permission) RemoveUserRoles(ctx context.Context, userID string, roleIDs []int64) (err error) {
	_, err = p.casbinEnforcer.RemoveFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", p.casbinUserPrefix, userID))
	if err != nil {
		return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}
	return nil
}

// checked
// RemoveUserAllRoles 删除用户全部角色
func (p *permission) RemoveUserAllRoles(ctx context.Context, userIDs []string) (err error) {
	for _, v := range userIDs {
		_, err = p.casbinEnforcer.RemoveFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", p.casbinUserPrefix, v))
		if err != nil {
			return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
		}
	}
	return nil
}

func (p *permission) AssignRolePermissions(ctx context.Context, roleID int64, permissionPointCodes []model.PermissionPointCode) (err error) {
	rules := make([][]string, len(permissionPointCodes))
	for k, v := range permissionPointCodes {
		rules[k] = []string{gconv.String(roleID), v.String(), "All"}
	}

	_, err = p.casbinEnforcer.AddNamedPolicies("p", rules)
	if err != nil {
		return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	return nil
}

// RemoveRolePermissionPoints 删除角色指定权限
func (p *permission) RemoveRolePermissions(ctx context.Context, roleID int64, permissionPointCodes []model.PermissionPointCode) (err error) {
	for _, v := range permissionPointCodes {
		// 同时匹配角色ID(v0)和权限ID(v1)
		_, err = p.casbinEnforcer.RemoveFilteredNamedPolicy("p", 1, gconv.String(roleID), v.String())
		if err != nil {
			return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
		}
	}
	return nil
}

// RemoveRoleAllPermissionPoints 删除角色全部权限
func (p *permission) RemoveRoleAllPermissions(ctx context.Context, roleIDs []int64) (err error) {
	for _, v := range roleIDs {
		_, err = p.casbinEnforcer.RemoveFilteredNamedPolicy("p", 0, gconv.String(v))
		if err != nil {
			return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
		}
	}
	return nil
}

func (p *permission) FilterUserRoleIDs(ctx context.Context, userID string, roleIDs []int64) (out []int64, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		out = roleIDs
		return
	}

	result, err := p.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		return
	}
	m := make(map[int64]bool, len(result))
	for _, v := range result {
		m[v] = true
	}

	for _, v := range roleIDs {
		if m[v] {
			out = append(out, v)
		}
	}
	return
}

// checked
func (p *permission) GetRoleIDsByUserID(ctx context.Context, userID string) (roleIDs []int64, err error) {
	groupPolicy, err := p.casbinEnforcer.GetFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", p.casbinUserPrefix, userID))
	if err != nil {
		return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	for _, v := range groupPolicy {
		roleIDs = append(roleIDs, gconv.Int64(v[1]))
	}

	roleIDs = libUtils.SliceUnique(roleIDs)

	return roleIDs, nil
}

func (p *permission) GetPermissionCodesByRoleID(ctx context.Context, roleID int64) (out []string, err error) {
	groupPolicy, err := p.casbinEnforcer.GetFilteredNamedPolicy("p", 0, fmt.Sprintf("%s%d", p.casbinRolePrefix, roleID))
	if err != nil {
		return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	for _, v := range groupPolicy {
		out = append(out, gconv.String(v[1]))
	}

	out = libUtils.SliceUnique(out)

	return out, nil
}

// TODO: 数据权限是否要在这里校验？该怎么校验
func (p *permission) HasPermission(ctx context.Context, permission model.PermissionPointCode, userID string, orgInfo *model.Org) (hasPermission bool, err error) {
	// 超级管理员拥有所有权限
	if p.IsSuperAdmin(ctx, userID) {
		return true, nil
	}

	// 组织管理员拥有组织下所有权限
	if p.IsOrgAdmin(ctx, userID, orgInfo) {
		// TODO: 检查操作的资源ID是否属于该组织
		return true, nil
	}

	permissions, err := p.GetPermissionPointsByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	g.Log().Info(ctx, "[DEBUG] HasPermission", permissions)
	for _, v := range permissions {
		if v.Code == permission {
			return true, nil
		}
	}

	return false, nil
}

func (p *permission) IsSuperAdmin(ctx context.Context, userID string) bool {
	return userID == model.DefaultSuperAdminID
}

func (p *permission) IsOrgAdmin(ctx context.Context, userID string, orgInfo *model.Org) bool {
	return orgInfo.ManagerID == userID
}

func (p *permission) GetPermissionPointsByUserID(ctx context.Context, userID string) (out []*model.PermissionPoint, err error) {
	if p.IsSuperAdmin(ctx, userID) {
		return nil, nil
	}

	// 获取用户所有角色ID
	roleIDs, err := p.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return p.GetPermissionPointsByRoleID(ctx, roleIDs)
}

func (p *permission) GetPermissionPointCodesByUserID(ctx context.Context, userID string) (out []model.PermissionPointCode, err error) {
	if p.IsSuperAdmin(ctx, userID) {
		return []model.PermissionPointCode{"*/*"}, nil
	}

	permissionPoints, err := p.GetPermissionPointsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.GetFrontPermissionPointCodes(permissionPoints), nil
}

func (p *permission) GetPermissionPointsByRoleID(ctx context.Context, roleIDs []int64) (out []*model.PermissionPoint, err error) {
	// 获取角色对应的权限点
	var codes []string
	for _, roleID := range roleIDs {
		result, err := p.GetPermissionCodesByRoleID(ctx, roleID)
		if err != nil {
			return nil, err
		}
		codes = append(codes, result...)
	}

	codes = libUtils.SliceUnique(codes)

	for _, code := range codes {
		permission, err := p.getPermissionByCode(ctx, code)
		if err != nil {
			return nil, err
		}
		out = append(out, permission)
	}

	return out, nil
}

func (p *permission) FilterPermissionPointsByUserID(ctx context.Context, permissionPoints []model.PermissionPointCode, userID string) (out []model.PermissionPointCode, err error) {
	return permissionPoints, nil
}

func (p *permission) FilterRoleIDsByUserID(ctx context.Context, roleIDs []int64, userID string) (out []int64, err error) {
	return
}

// 私有方法
// getAllPermissionPointsMap 获取所有权限点
func (p *permission) getAllPermissionPointsMap(ctx context.Context) (out map[string]*model.PermissionPoint, err error) {
	var entities []*entity.PermissionPoint
	err = dao.PermissionPoint.Ctx(ctx).Scan(&entities)
	if err != nil {
		return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	out = make(map[string]*model.PermissionPoint, len(entities))
	for _, v := range entities {
		out[v.Code] = p.convertPermissionEntityToModel(v)
	}

	return out, nil
}

// getPermissionByCode 根据权限点Code获取权限点
func (p *permission) getPermissionByCode(ctx context.Context, code string) (out *model.PermissionPoint, err error) {
	var pps []*entity.PermissionPoint
	err = dao.PermissionPoint.Ctx(ctx).Where(dao.PermissionPoint.Columns().Code, code).Scan(&pps)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w, %v", model.ErrRecordNotFound, fmt.Errorf("permission_code: %s", code))
		}
		return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	if len(pps) == 0 {
		return nil, nil
	}

	out = p.convertPermissionEntityToModel(pps[0])

	// 一个权限点对应多个资源
	for _, v := range pps {
		var resourceEntity entity.Resource
		err = dao.Resource.Ctx(ctx).Where(dao.Resource.Columns().ID, v.ResourceID).Scan(&resourceEntity)
		if err != nil {
			return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
		}
		out.Resources = append(out.Resources, p.convertResourceEntityToModel(&resourceEntity))
	}

	return out, nil
}

// convertEntityToModel 转换实体到模型
func (p *permission) convertPermissionEntityToModel(entity *entity.PermissionPoint) (out *model.PermissionPoint) {
	out = &model.PermissionPoint{
		ID:          entity.ID,
		Code:        model.PermissionPointCode(entity.Code),
		CodeName:    entity.CodeName,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}

	return
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
