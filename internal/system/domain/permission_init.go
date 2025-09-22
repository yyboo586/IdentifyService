package domain

import (
	"IdentifyService/internal/system/dao"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/model/entity"
	"context"
	"database/sql"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 目前不需要考虑并发，因为服务单节点部署。
// TODO: 目前的代码逻辑有问题
// 1、新增资源、旧权限点无法关联到新的资源
func (p *permission) Init(ctx context.Context) (err error) {
	// 初始化资源
	err = p.initResource(ctx)
	if err != nil {
		return err
	}

	// 初始化权限点
	err = p.initPermissionPoint(ctx)
	if err != nil {
		return err
	}

	// 初始化角色
	err = p.initDefaultRoles(ctx)
	if err != nil {
		g.Log().Error(ctx, "初始化默认角色失败:", err)
		return err
	}

	return nil
}

/*
initResource: 初始化资源列表
1、如果资源不存在，则创建资源
2、如果资源存在，则获取资源
3、构建资源缓存列表，便于下一步构建权限点
*/
func (p *permission) initResource(ctx context.Context) (err error) {
	g.Log().Info(ctx, "开始初始化资源列表...")

	allResources := append([]*model.Resource{}, model.UserResources...)
	allResources = append(allResources, model.RoleResources...)
	allResources = append(allResources, model.OrgResources...)

	for _, v := range allResources {
		var resourceEntity entity.Resource
		err = dao.Resource.Ctx(ctx).Where(dao.Resource.Columns().Type, v.Type).Where(dao.Resource.Columns().Code, v.Code).Scan(&resourceEntity)
		if err != nil {
			if err != sql.ErrNoRows {
				return fmt.Errorf("检查资源是否存在失败: %w", err)
			}
		}

		// 资源不存在，则创建资源
		if err == sql.ErrNoRows {
			_, err := dao.Resource.Ctx(ctx).Data(map[string]any{
				dao.Resource.Columns().Type: v.Type,
				dao.Resource.Columns().Code: v.Code,
			}).Insert()
			if err != nil {
				return fmt.Errorf("创建资源失败: %w", err)
			}
		} else {
			// 资源存在，则跳过
			continue
		}
	}

	g.Log().Info(ctx, "资源列表初始化完成...")
	return
}

func (p *permission) getAllResources(ctx context.Context) (resourcesMap map[string]*model.Resource, err error) {
	var resourceEntities []*entity.Resource
	err = dao.Resource.Ctx(ctx).Scan(&resourceEntities)
	if err != nil {
		return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	resourcesMap = make(map[string]*model.Resource, len(resourceEntities))
	for _, v := range resourceEntities {
		resourcesMap[v.Code] = p.convertResourceEntityToModel(v)
	}

	return resourcesMap, nil
}

func (p *permission) initPermissionPoint(ctx context.Context) (err error) {
	g.Log().Info(ctx, "开始初始化权限点...")

	resourcesMap, err := p.getAllResources(ctx)
	if err != nil {
		return fmt.Errorf("获取资源列表失败: %w", err)
	}

	allPermissionPoints := append([]*model.PermissionPoint{}, model.UserPermissions...)
	allPermissionPoints = append(allPermissionPoints, model.RolePermissions...)
	allPermissionPoints = append(allPermissionPoints, model.OrgPermissions...)

	// 初始化权限点
	for _, pp := range allPermissionPoints {
		// 1、检查权限点是否已存在
		var permissionEntity entity.PermissionPoint
		err = dao.PermissionPoint.Ctx(ctx).Where(dao.PermissionPoint.Columns().Code, pp.Code).Scan(&permissionEntity)
		if err != nil {
			if err != sql.ErrNoRows {
				return fmt.Errorf("检查权限点是否存在失败: %w", err)
			}
		}

		// 权限点不存在，则创建权限点
		if err == sql.ErrNoRows {
			var data []map[string]any
			for _, v := range pp.Resources {
				resource, ok := resourcesMap[v.Code]
				if !ok {
					return fmt.Errorf("数据异常,资源不存在: %s", v.Code)
				}
				data = append(data, map[string]any{
					dao.PermissionPoint.Columns().Code:        pp.Code,
					dao.PermissionPoint.Columns().CodeName:    pp.CodeName,
					dao.PermissionPoint.Columns().ResourceID:  resource.ID,
					dao.PermissionPoint.Columns().Description: pp.Description,
				})
			}
			_, err = dao.PermissionPoint.Ctx(ctx).Data(data).Insert()
			if err != nil {
				return fmt.Errorf("创建权限点失败: %w", err)
			}
		} else {
			// 权限点存在，则跳过
			continue
		}
	}

	g.Log().Info(ctx, "权限点初始化完成...")
	return nil
}

/*
p role_id permission_point_id
g user_id role_id
*/
// initDefaultRoles 初始化默认角色
func (p *permission) initDefaultRoles(ctx context.Context) (err error) {
	g.Log().Info(ctx, "开始初始化默认角色...")

	permissionPointsMap, err := p.getAllPermissionPointsMap(ctx)
	if err != nil {
		return err
	}

	allRoles := append([]*model.Role{}, model.DefaultRoles...)

	for _, v := range allRoles {
		var roleEntity entity.Role
		err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Name, v.Name).Scan(&roleEntity)
		if err != nil {
			if err != sql.ErrNoRows {
				return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
			}
		}

		var roleID int64
		if err == sql.ErrNoRows {
			roleID, err = dao.Role.Ctx(ctx).Data(map[string]interface{}{
				dao.Role.Columns().OrgID:     v.OrgID,
				dao.Role.Columns().PID:       v.PID,
				dao.Role.Columns().Name:      v.Name,
				dao.Role.Columns().CreatorID: v.CreatorID,
			}).InsertAndGetId()
			if err != nil {
				return fmt.Errorf("数据插入失败: %w", err)
			}
		} else {
			roleID = roleEntity.ID
		}

		for _, pp := range v.PermissionPoints {
			_, ok := permissionPointsMap[pp.String()]
			if !ok {
				return fmt.Errorf("数据异常,权限点不存在: %s", pp.String())
			}

			exists, err := dao.CasbinRule.Ctx(ctx).
				Where(dao.CasbinRule.Columns().Ptype, "p").
				Where(dao.CasbinRule.Columns().V0, fmt.Sprintf("%s%s", p.casbinRolePrefix, gconv.String(roleID))).
				Where(dao.CasbinRule.Columns().V1, pp.String()).Exist()
			if err != nil {
				return fmt.Errorf("检查权限点是否已经分配给角色失败: %w", err)
			}
			if !exists {
				_, err = p.casbinEnforcer.AddNamedPolicy("p", []string{fmt.Sprintf("%s%s", p.casbinRolePrefix, gconv.String(roleID)), pp.String(), "All"})
				if err != nil {
					return fmt.Errorf("添加权限点失败: %w", err)
				}
			}
		}
	}

	g.Log().Info(ctx, "默认角色初始化完成")
	return nil
}
