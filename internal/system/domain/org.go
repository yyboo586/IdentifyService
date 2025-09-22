package domain

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/system/dao"
	"IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/model/entity"
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
)

var (
	orgOnce     sync.Once
	orgInstance interfaces.IOrg
)

func NewOrg(userDomain interfaces.IUser) interfaces.IOrg {
	orgOnce.Do(func() {
		orgInstance = &org{
			userDomain: userDomain,
		}
	})
	return orgInstance
}

type org struct {
	userDomain interfaces.IUser
}

func (s *org) Create(ctx context.Context, tx gdb.TX, in *model.Org) (err error) {
	orgInsertData := map[string]interface{}{
		dao.Org.Columns().ID:          in.ID,
		dao.Org.Columns().PID:         in.PID,
		dao.Org.Columns().Name:        in.Name,
		dao.Org.Columns().ManagerID:   in.ManagerID,
		dao.Org.Columns().ManagerName: in.ManagerName,
		dao.Org.Columns().Status:      model.OrgStatusEnabled,
	}

	if tx != nil {
		_, err = dao.Org.Ctx(ctx).TX(tx).Insert(orgInsertData)
	} else {
		_, err = dao.Org.Ctx(ctx).Insert(orgInsertData)
	}

	return
}

func (s *org) Delete(ctx context.Context, id string) (err error) {
	hasUser, err := s.userDomain.IsOrgHasUsers(ctx, id)
	if err != nil {
		return err
	}
	if hasUser {
		return fmt.Errorf("组织下还有用户，不能删除")
	}

	if id == model.DefaultOrgID {
		return fmt.Errorf("默认组织不能删除")
	}

	_, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Delete()
	if err != nil {
		return fmt.Errorf("删除组织失败: %w", err)
	}

	return nil
}

func (s *org) Edit(ctx context.Context, in *model.Org) (err error) {
	var data map[string]interface{} = make(map[string]interface{})
	if in.Name != "" {
		data[dao.Org.Columns().Name] = in.Name
	}
	if in.ManagerID != "" {
		data[dao.Org.Columns().ManagerID] = in.ManagerID
	}
	if in.ManagerName != "" {
		data[dao.Org.Columns().ManagerName] = in.ManagerName
	}
	if model.IsValidOrgStatus(in.Status) {
		data[dao.Org.Columns().Status] = in.Status
	}

	_, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, in.ID).Update(data)
	if err != nil {
		return fmt.Errorf("修改组织信息失败: %w", err)
	}

	return
}

func (s *org) Get(ctx context.Context, id string) (out *model.Org, err error) {
	var org entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Scan(&org)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("组织不存在: %s", id)
		}
		return nil, fmt.Errorf("获取组织详情失败: %w", err)
	}

	out = s.convertEntityToLogics(&org)
	return
}

func (s *org) GetTree(ctx context.Context, id string) (out *system.OrgTreeNode, err error) {
	var orgs *entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().ID, id).Scan(&orgs)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("组织不存在: %s", id)
		}
		return nil, fmt.Errorf("获取组织失败: %w", err)
	}

	out = &system.OrgTreeNode{
		OrgInfo: s.convertModelToSystem(s.convertEntityToLogics(orgs)),
	}

	var children []*entity.Org
	err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().PID, id).Scan(&children)
	if err != nil {
		return nil, fmt.Errorf("获取组织子节点失败: %w", err)
	}

	if len(children) > 0 {
		out.Children = make([]*system.OrgTreeNode, 0, len(children))
		for _, v := range children {
			var childNode *system.OrgTreeNode
			childNode, err = s.GetTree(ctx, v.ID)
			if err != nil {
				return nil, fmt.Errorf("获取组织子节点失败: %w", err)
			}
			out.Children = append(out.Children, childNode)
		}
	}

	return
}

func (s *org) GetRoleTrees(ctx context.Context, orgID string) (out []*system.RoleNode, err error) {
	return
}

func (s *org) IsEnabled(ctx context.Context, orgID string) (err error) {
	orgInfo, err := s.Get(ctx, orgID)
	if err != nil {
		return err
	}

	if orgInfo.Status != model.OrgStatusEnabled {
		return fmt.Errorf("组织已被禁用")
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
