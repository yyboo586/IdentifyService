package logics

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

var (
	orgOnce     sync.Once
	orgInstance service.IOrg
)

func NewOrg() service.IOrg {
	orgOnce.Do(func() {
		orgInstance = &org{}
	})
	return orgInstance
}

type org struct {
}

func (s *org) Add(ctx context.Context, in *model.Org) (orgID string, err error) {
	orgID = uuid.New().String()
	orgInsertData := map[string]interface{}{
		dao.Org.Columns().ID:          orgID,
		dao.Org.Columns().PID:         in.PID,
		dao.Org.Columns().Name:        in.Name,
		dao.Org.Columns().ManagerID:   in.ManagerID,
		dao.Org.Columns().ManagerName: in.ManagerName,
		dao.Org.Columns().Status:      model.OrgStatusEnabled,
	}

	_, err = dao.Org.Ctx(ctx).Insert(orgInsertData)
	if err != nil {
		err = fmt.Errorf("添加组织失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	return
}

// TODO
// 1、删除组织
// 2、删除组织下的所有用户
// 3、删除组织下的所有角色
// 4、删除组织下的所有权限
func (s *org) Delete(ctx context.Context, id string) (err error) {
	if id == consts.DefaultBackgroundOrgID || id == consts.DefaultFrontOrgID {
		err = errors.New("默认组织不能删除")
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.Org.Ctx(ctx).TX(tx).Where(dao.Org.Columns().ID, id).Delete()
		if err != nil {
			err = fmt.Errorf("删除组织失败: %w", err)
			return
		}

		_, err = dao.User.Ctx(ctx).TX(tx).Where(dao.User.Columns().OrgID, id).Delete()
		if err != nil {
			err = fmt.Errorf("删除用户失败: %w", err)
		}

		return
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (s *org) EditBasicInfo(ctx context.Context, in *model.Org) (err error) {
	updateData := map[string]interface{}{
		dao.Org.Columns().Name:        in.Name,
		dao.Org.Columns().ManagerID:   in.ManagerID,
		dao.Org.Columns().ManagerName: in.ManagerName,
	}

	_, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, in.ID).Update(updateData)
	if err != nil {
		err = fmt.Errorf("修改组织基本信息失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (s *org) EditStatus(ctx context.Context, id string, enabled bool) (err error) {
	updateData := map[string]interface{}{}
	if enabled {
		updateData[dao.Org.Columns().Status] = model.OrgStatusEnabled
	} else {
		updateData[dao.Org.Columns().Status] = model.OrgStatusDisabled
	}

	_, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Update(updateData)
	if err != nil {
		err = fmt.Errorf("修改组织状态失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (s *org) Get(ctx context.Context, id string) (out *system.OrgInfo, err error) {
	var org entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Scan(&org)
	if err != nil {
		err = fmt.Errorf("获取组织详情失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	out = s.convertModelToSystem(s.convertEntityToLogics(&org))
	return
}

func (s *org) GetTree(ctx context.Context, id string) (out *system.OrgTreeNode, err error) {
	var orgs *entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Scan(&orgs)
	if err != nil {
		err = fmt.Errorf("获取组织失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	out = &system.OrgTreeNode{
		OrgInfo: s.convertModelToSystem(s.convertEntityToLogics(orgs)),
	}

	var children []*entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().PID, id).Scan(&children)
	if err != nil {
		err = fmt.Errorf("获取组织子节点失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	if len(children) > 0 {
		out.Children = make([]*system.OrgTreeNode, 0, len(children))
		for _, v := range children {
			var childNode *system.OrgTreeNode
			childNode, err = s.GetTree(ctx, v.ID)
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			out.Children = append(out.Children, childNode)
		}
	}

	return
}

// 1、获取根组织节点
// 2、获取根组织节点下的所有子节点
func (s *org) ListTrees(ctx context.Context) (out []*system.OrgTreeNode, err error) {
	var rootNodes []*entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().PID, "0").Scan(&rootNodes)
	if err != nil {
		err = fmt.Errorf("获取根组织列表失败: %w", err)
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range rootNodes {
		var rootNode *system.OrgTreeNode
		rootNode, err = s.GetTree(ctx, v.ID)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		out = append(out, rootNode)
	}

	return
}

func (s *org) AssertExistsByID(ctx context.Context, id string) (exists bool, err error) {
	exists, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Exist()
	if err != nil {
		err = fmt.Errorf("检查组织是否存在失败: %w", err)
		return
	}

	if !exists {
		err = fmt.Errorf("组织不存在: %s", id)
		return
	}

	return
}

func (s *org) convertEntityToLogics(in *entity.Org) (out *model.Org) {
	out = &model.Org{
		ID:          in.ID,
		PID:         in.PID,
		Name:        in.Name,
		ManagerID:   in.ManagerID,
		ManagerName: in.ManagerName,
		Status:      model.OrgStatus(in.Status),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return
}

func (s *org) convertModelToSystem(in *model.Org) (out *system.OrgInfo) {
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
