package logics

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"IdentifyService/library/liberr"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	authRuleOnce     sync.Once
	authRuleInstance *authRule
)

type authRule struct {
	enforcer *casbin.SyncedEnforcer
}

func NewAuthRule() service.IAuthRule {
	authRuleOnce.Do(func() {
		enforcer, err := service.CasbinEnforcer(context.Background())
		if err != nil {
			panic(err)
		}
		authRuleInstance = &authRule{
			enforcer: enforcer,
		}
	})
	return authRuleInstance
}

func (a *authRule) Add(ctx context.Context, req *system.RuleAddReq) (ruleID int64, err error) {
	data := map[string]interface{}{
		dao.AuthRule.Columns().Pid:       req.Pid,
		dao.AuthRule.Columns().Name:      req.Name,
		dao.AuthRule.Columns().Type:      req.Type,
		dao.AuthRule.Columns().Path:      req.Path,
		dao.AuthRule.Columns().Component: req.Component,

		dao.AuthRule.Columns().Title:                    req.Medata.Title,
		dao.AuthRule.Columns().Icon:                     req.Medata.Icon,
		dao.AuthRule.Columns().ActiveIcon:               req.Medata.ActiveIcon,
		dao.AuthRule.Columns().KeepAlive:                req.Medata.KeepAlive,
		dao.AuthRule.Columns().HideInMenu:               req.Medata.HideInMenu,
		dao.AuthRule.Columns().HideInTab:                req.Medata.HideInTab,
		dao.AuthRule.Columns().HideInBreadcrumb:         req.Medata.HideInBreadcrumb,
		dao.AuthRule.Columns().HideChildrenInMenu:       req.Medata.HideChildrenInMenu,
		dao.AuthRule.Columns().Authority:                req.Medata.Authority,
		dao.AuthRule.Columns().Badge:                    req.Medata.Badge,
		dao.AuthRule.Columns().BadgeType:                req.Medata.BadgeType,
		dao.AuthRule.Columns().BadgeVariants:            req.Medata.BadgeVariants,
		dao.AuthRule.Columns().FullPathKey:              req.Medata.FullPathKey,
		dao.AuthRule.Columns().ActivePath:               req.Medata.ActivePath,
		dao.AuthRule.Columns().AffixTab:                 req.Medata.AffixTab,
		dao.AuthRule.Columns().AffixTabOrder:            req.Medata.AffixTabOrder,
		dao.AuthRule.Columns().IframeSrc:                req.Medata.IframeSrc,
		dao.AuthRule.Columns().IgnoreAccess:             req.Medata.IgnoreAccess,
		dao.AuthRule.Columns().Link:                     req.Medata.Link,
		dao.AuthRule.Columns().MaxNumOfOpenTab:          req.Medata.MaxNumOfOpenTab,
		dao.AuthRule.Columns().MenuVisibleWithForbidden: req.Medata.MenuVisibleWithForbidden,
		dao.AuthRule.Columns().OpenInNewWindow:          req.Medata.OpenInNewWindow,
		dao.AuthRule.Columns().Order:                    req.Medata.Order,
		dao.AuthRule.Columns().Query:                    req.Medata.Query,
		dao.AuthRule.Columns().NoBasicLayout:            req.Medata.NoBasicLayout,
	}
	if req.Medata.KeepAlive {
		data[dao.AuthRule.Columns().KeepAlive] = 1
	} else {
		data[dao.AuthRule.Columns().KeepAlive] = 0
	}
	if req.Medata.HideInMenu {
		data[dao.AuthRule.Columns().HideInMenu] = 1
	} else {
		data[dao.AuthRule.Columns().HideInMenu] = 0
	}
	if req.Medata.HideInTab {
		data[dao.AuthRule.Columns().HideInTab] = 1
	} else {
		data[dao.AuthRule.Columns().HideInTab] = 0
	}
	if req.Medata.HideInBreadcrumb {
		data[dao.AuthRule.Columns().HideInBreadcrumb] = 1
	} else {
		data[dao.AuthRule.Columns().HideInBreadcrumb] = 0
	}
	if req.Medata.HideChildrenInMenu {
		data[dao.AuthRule.Columns().HideChildrenInMenu] = 1
	} else {
		data[dao.AuthRule.Columns().HideChildrenInMenu] = 0
	}
	if req.Medata.FullPathKey {
		data[dao.AuthRule.Columns().FullPathKey] = 1
	} else {
		data[dao.AuthRule.Columns().FullPathKey] = 0
	}
	if req.Medata.AffixTab {
		data[dao.AuthRule.Columns().AffixTab] = 1
	} else {
		data[dao.AuthRule.Columns().AffixTab] = 0
	}
	if req.Medata.IgnoreAccess {
		data[dao.AuthRule.Columns().IgnoreAccess] = 1
	} else {
		data[dao.AuthRule.Columns().IgnoreAccess] = 0
	}
	if req.Medata.MenuVisibleWithForbidden {
		data[dao.AuthRule.Columns().MenuVisibleWithForbidden] = 1
	} else {
		data[dao.AuthRule.Columns().MenuVisibleWithForbidden] = 0
	}
	if req.Medata.OpenInNewWindow {
		data[dao.AuthRule.Columns().OpenInNewWindow] = 1
	} else {
		data[dao.AuthRule.Columns().OpenInNewWindow] = 0
	}
	if req.Medata.NoBasicLayout {
		data[dao.AuthRule.Columns().NoBasicLayout] = 1
	} else {
		data[dao.AuthRule.Columns().NoBasicLayout] = 0
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		ruleID, err = dao.AuthRule.Ctx(ctx).TX(tx).InsertAndGetId(data)
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = gerror.New("菜单规则已经存在")
				return
			}
			err = fmt.Errorf("添加菜单失败: %w", err)
			return
		}
		return
	})
	return
}

// TODO Check
func (a *authRule) DeleteByIDs(ctx context.Context, ids []int64) (err error) {
	var list []*model.AuthRule
	list, err = a.GetAllAuthRules(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	childrenIds := make([]int64, 0, len(list))
	for _, id := range ids {
		rules := a.findChildrenByParentID(list, int64(id))
		for _, child := range rules {
			childrenIds = append(childrenIds, child.ID)
		}
	}
	ids = append(ids, childrenIds...)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		return g.Try(ctx, func(ctx context.Context) {
			_, err = dao.AuthRule.Ctx(ctx).WhereIn(dao.AuthRule.Columns().ID, ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除失败")
			// 删除权限
			for _, v := range ids {
				_, err = a.enforcer.RemoveFilteredNamedPolicy("p", 1, gconv.String(v))
				liberr.ErrIsNil(ctx, err)
			}
		})
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (a *authRule) Update(ctx context.Context, req *system.RuleUpdateReq) (err error) {
	data := map[string]interface{}{
		dao.AuthRule.Columns().Pid:       req.Pid,
		dao.AuthRule.Columns().Name:      req.Name,
		dao.AuthRule.Columns().Type:      req.Type,
		dao.AuthRule.Columns().Path:      req.Path,
		dao.AuthRule.Columns().Component: req.Component,

		dao.AuthRule.Columns().Title:                    req.Medata.Title,
		dao.AuthRule.Columns().Icon:                     req.Medata.Icon,
		dao.AuthRule.Columns().ActiveIcon:               req.Medata.ActiveIcon,
		dao.AuthRule.Columns().KeepAlive:                req.Medata.KeepAlive,
		dao.AuthRule.Columns().HideInMenu:               req.Medata.HideInMenu,
		dao.AuthRule.Columns().HideInTab:                req.Medata.HideInTab,
		dao.AuthRule.Columns().HideInBreadcrumb:         req.Medata.HideInBreadcrumb,
		dao.AuthRule.Columns().HideChildrenInMenu:       req.Medata.HideChildrenInMenu,
		dao.AuthRule.Columns().Authority:                req.Medata.Authority,
		dao.AuthRule.Columns().Badge:                    req.Medata.Badge,
		dao.AuthRule.Columns().BadgeType:                req.Medata.BadgeType,
		dao.AuthRule.Columns().BadgeVariants:            req.Medata.BadgeVariants,
		dao.AuthRule.Columns().FullPathKey:              req.Medata.FullPathKey,
		dao.AuthRule.Columns().ActivePath:               req.Medata.ActivePath,
		dao.AuthRule.Columns().AffixTab:                 req.Medata.AffixTab,
		dao.AuthRule.Columns().AffixTabOrder:            req.Medata.AffixTabOrder,
		dao.AuthRule.Columns().IframeSrc:                req.Medata.IframeSrc,
		dao.AuthRule.Columns().IgnoreAccess:             req.Medata.IgnoreAccess,
		dao.AuthRule.Columns().Link:                     req.Medata.Link,
		dao.AuthRule.Columns().MaxNumOfOpenTab:          req.Medata.MaxNumOfOpenTab,
		dao.AuthRule.Columns().MenuVisibleWithForbidden: req.Medata.MenuVisibleWithForbidden,
		dao.AuthRule.Columns().OpenInNewWindow:          req.Medata.OpenInNewWindow,
		dao.AuthRule.Columns().Order:                    req.Medata.Order,
		dao.AuthRule.Columns().Query:                    req.Medata.Query,
		dao.AuthRule.Columns().NoBasicLayout:            req.Medata.NoBasicLayout,
	}
	if req.Medata.KeepAlive {
		data[dao.AuthRule.Columns().KeepAlive] = 1
	} else {
		data[dao.AuthRule.Columns().KeepAlive] = 0
	}
	if req.Medata.HideInMenu {
		data[dao.AuthRule.Columns().HideInMenu] = 1
	} else {
		data[dao.AuthRule.Columns().HideInMenu] = 0
	}
	if req.Medata.HideInTab {
		data[dao.AuthRule.Columns().HideInTab] = 1
	} else {
		data[dao.AuthRule.Columns().HideInTab] = 0
	}
	if req.Medata.HideInBreadcrumb {
		data[dao.AuthRule.Columns().HideInBreadcrumb] = 1
	} else {
		data[dao.AuthRule.Columns().HideInBreadcrumb] = 0
	}
	if req.Medata.HideChildrenInMenu {
		data[dao.AuthRule.Columns().HideChildrenInMenu] = 1
	} else {
		data[dao.AuthRule.Columns().HideChildrenInMenu] = 0
	}
	if req.Medata.FullPathKey {
		data[dao.AuthRule.Columns().FullPathKey] = 1
	} else {
		data[dao.AuthRule.Columns().FullPathKey] = 0
	}
	if req.Medata.AffixTab {
		data[dao.AuthRule.Columns().AffixTab] = 1
	} else {
		data[dao.AuthRule.Columns().AffixTab] = 0
	}
	if req.Medata.IgnoreAccess {
		data[dao.AuthRule.Columns().IgnoreAccess] = 1
	} else {
		data[dao.AuthRule.Columns().IgnoreAccess] = 0
	}
	if req.Medata.MenuVisibleWithForbidden {
		data[dao.AuthRule.Columns().MenuVisibleWithForbidden] = 1
	} else {
		data[dao.AuthRule.Columns().MenuVisibleWithForbidden] = 0
	}
	if req.Medata.OpenInNewWindow {
		data[dao.AuthRule.Columns().OpenInNewWindow] = 1
	} else {
		data[dao.AuthRule.Columns().OpenInNewWindow] = 0
	}
	if req.Medata.NoBasicLayout {
		data[dao.AuthRule.Columns().NoBasicLayout] = 1
	} else {
		data[dao.AuthRule.Columns().NoBasicLayout] = 0
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.AuthRule.Ctx(ctx).TX(tx).Where(dao.AuthRule.Columns().ID, req.ID).Update(data)
		if err != nil {
			err = fmt.Errorf("更新菜单失败: %w", err)
			return
		}
		return
	})
	return
}

/*
	func (a *authRule) GetDetailsByID(ctx context.Context, id int64) (out *model.AuthRule, err error) {
		var authRule entity.AuthRule
		err = dao.AuthRule.Ctx(ctx).Where(dao.AuthRule.Columns().ID, id).Scan(&authRule)
		if err != nil {
			return
		}

		out = a.convertEntityToModel(&authRule)

		return
	}
*/

func (a *authRule) GetMenuTreesByUserID(ctx context.Context, userID string, includeButton bool) (out []*model.AuthRuleNode, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		out, err = a.listTrees(ctx, includeButton)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		return
	}

	// 根据用户ID获取用户角色ID列表
	roleIDs, err := roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	g.Log().Info(ctx, "[DEBUG] roleIDs: ", "roleIDs", roleIDs)

	out, err = a.getMenuTreesByRoleIDs(ctx, roleIDs, includeButton)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (a *authRule) GetButtonListByUserID(ctx context.Context, userID string) (out []*model.AuthRule, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		out, err = a.listButtons(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		return
	}

	// 根据用户ID获取用户角色ID列表
	roleIDs, err := roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	out, err = a.getButtonListByRoleIDs(ctx, roleIDs)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

/*
func (a *authRule) GetMenuListByRoleID(ctx context.Context, roleID int64) (out []*model.AuthRule, err error) {
	ruleIDs, err := a.getRuleIDsByRoleIDs(ctx, []int64{roleID})
	if err != nil {
		return
	}
	var authRuleEntities []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).WhereIn(dao.AuthRule.Columns().ID, ruleIDs).Scan(&authRuleEntities)
	if err != nil {
		return
	}

	for _, v := range authRuleEntities {
		rule := a.convertEntityToModel(v)
		if rule.Type == model.MenuTypeButton {
			continue
		}
		out = append(out, rule)
	}
	return
}
*/

func (a *authRule) GetAllAuthRules(ctx context.Context) (list []*model.AuthRule, err error) {
	var entitys []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).Scan(&entitys)
	if err != nil {
		return
	}

	for _, entity := range entitys {
		list = append(list, a.convertEntityToModel(entity))
	}
	return
}

// 私有方法
func (a *authRule) listTrees(ctx context.Context, includeButton bool) (list []*model.AuthRuleNode, err error) {
	// 获取顶层目录ID
	var rootNodes []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).Fields(dao.AuthRule.Columns().ID).Where(dao.AuthRule.Columns().Pid, 0).Scan(&rootNodes)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 构建子树
	for _, v := range rootNodes {
		var treeNode *model.AuthRuleNode
		treeNode, err = a.getTreeByID(ctx, v.ID, includeButton)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		list = append(list, treeNode)
	}
	return
}

func (a *authRule) getMenuTreesByRoleIDs(ctx context.Context, roleIDs []int64, includeButton bool) (trees []*model.AuthRuleNode, err error) {
	ruleIDs, err := a.getRuleIDsByRoleIDs(ctx, roleIDs)
	if err != nil {
		return
	}

	var authRuleEntities []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).WhereIn(dao.AuthRule.Columns().ID, ruleIDs).Scan(&authRuleEntities)
	if err != nil {
		return
	}

	var rootNodeIDs []int64
	var authRules []*model.AuthRule
	for _, v := range authRuleEntities {
		rule := a.convertEntityToModel(v)

		if !includeButton && rule.Type == model.MenuTypeButton {
			continue
		}

		g.Log().Info(ctx, "[DEBUG] authRuleEntitiy: ", "rule", *rule)
		authRules = append(authRules, rule)
		if rule.Pid == 0 {
			rootNodeIDs = append(rootNodeIDs, rule.ID)
		}
	}

	for _, id := range rootNodeIDs {
		g.Log().Info(ctx, "[DEBUG] buildTree: rootNodeIDs", "id", id)
		trees = append(trees, a.buildTree(id, authRules))
	}

	return
}

func (a *authRule) listButtons(ctx context.Context) (out []*model.AuthRule, err error) {
	fields := []interface{}{
		dao.AuthRule.Columns().ID,
		dao.AuthRule.Columns().Pid,
		dao.AuthRule.Columns().Name,
		dao.AuthRule.Columns().Type,
	}
	var entitys []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).Fields(fields...).Where(dao.AuthRule.Columns().Type, model.MenuTypeButton).Scan(&entitys)
	if err != nil {
		return
	}

	for _, entity := range entitys {
		out = append(out, a.convertEntityToModel(entity))
	}
	return
}

func (a *authRule) getButtonListByRoleIDs(ctx context.Context, roleIDs []int64) (out []*model.AuthRule, err error) {
	ruleIDs, err := a.getRuleIDsByRoleIDs(ctx, roleIDs)
	if err != nil {
		return
	}

	var authRuleEntities []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).WhereIn(dao.AuthRule.Columns().ID, ruleIDs).Scan(&authRuleEntities)
	if err != nil {
		return
	}

	for _, v := range authRuleEntities {
		rule := a.convertEntityToModel(v)
		if rule.Type == model.MenuTypeButton {
			out = append(out, rule)
			continue
		}
	}

	return
}

// bindRoleRule 角色绑定菜单权限
func (a *authRule) bindRoleRule(ctx context.Context, ruleId interface{}, roleIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		for _, roleId := range roleIds {
			_, err = a.enforcer.AddNamedPolicy("p", fmt.Sprintf("%d", roleId), fmt.Sprintf("%d", ruleId), "All")
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (a *authRule) updateRoleRule(ctx context.Context, ruleId int64, roleIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧权限
		_, err = a.enforcer.RemoveFilteredNamedPolicy("p", 1, gconv.String(ruleId))
		liberr.ErrIsNil(ctx, err)
		// 添加新权限
		err = a.bindRoleRule(ctx, ruleId, roleIds)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (a *authRule) getTreeByID(ctx context.Context, id int64, includeButton bool) (out *model.AuthRuleNode, err error) {
	var authRuleEntity *entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).Where(dao.AuthRule.Columns().ID, id).Scan(&authRuleEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("菜单不存在")
		}
		g.Log().Error(ctx, err)
		return
	}

	out = &model.AuthRuleNode{
		AuthRule: a.convertEntityToModel(authRuleEntity),
	}

	var children []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).Fields(dao.AuthRule.Columns().ID).Where(dao.AuthRule.Columns().Pid, id).Scan(&children)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	if len(children) > 0 {
		for _, v := range children {
			var child *model.AuthRuleNode
			child, err = a.getTreeByID(ctx, v.ID, includeButton)
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			out.Children = append(out.Children, child)
		}
	}

	return
}

func (a *authRule) findChildrenByParentID(list []*model.AuthRule, pid int64) []*model.AuthRule {
	children := make([]*model.AuthRule, 0, len(list))
	for _, v := range list {
		if v.Pid == pid {
			children = append(children, v)
			fChildren := a.findChildrenByParentID(list, v.ID)
			children = append(children, fChildren...)
		}
	}
	return children
}

func (a *authRule) getRuleIDsByRoleIDs(ctx context.Context, roleIDs []int64) (ruleIDs []int64, err error) {
	for _, roleID := range roleIDs {
		var policies [][]string
		policies, err = a.enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(roleID))
		if err != nil {
			return
		}
		for _, policy := range policies {
			// policy ["1","1","ALL"]
			g.Log().Info(ctx, "[DEBUG] GetFilteredNamedPolicy: ", "policy", policy)
			ruleIDs = append(ruleIDs, gconv.Int64(policy[1]))
		}
	}

	ruleIDs = libUtils.SliceUnique(ruleIDs)

	return
}

func (a *authRule) buildTree(id int64, rules []*model.AuthRule) (root *model.AuthRuleNode) {
	// 1、找到根节点
	for _, rule := range rules {
		if rule.ID == id {
			root = &model.AuthRuleNode{
				AuthRule: rule,
			}
			break
		}
	}

	if root == nil {
		return
	}

	// 2、递归构建子树
	for _, v := range rules {
		if v.Pid == id {
			root.Children = append(root.Children, a.buildTree(v.ID, rules))
		}
	}
	return
}

func (a *authRule) filterRuleIDsByRoleIDs(ctx context.Context, ruleIDs []int64, roleIDs []int64) (out []int64, err error) {
	accessRuleIDs, err := a.getRuleIDsByRoleIDs(ctx, roleIDs)
	if err != nil {
		return
	}
	accessRulesMap := make(map[int64]bool)
	for _, ruleID := range accessRuleIDs {
		accessRulesMap[ruleID] = true
	}

	for _, ruleID := range ruleIDs {
		if _, ok := accessRulesMap[ruleID]; ok {
			out = append(out, ruleID)
		}
	}
	return
}

func (a *authRule) convertEntityToModel(in *entity.AuthRule) (out *model.AuthRule) {
	out = &model.AuthRule{
		ID:        in.ID,
		Pid:       in.Pid,
		Name:      in.Name,
		Type:      model.AuthRuleType(in.Type),
		Path:      in.Path,
		Component: in.Component,
		MMeta: &model.AuthRuleMeta{
			Title:                    in.Title,
			Icon:                     in.Icon,
			ActiveIcon:               in.ActiveIcon,
			KeepAlive:                in.KeepAlive == 1,
			HideInMenu:               in.HideInMenu == 1,
			HideInTab:                in.HideInTab == 1,
			HideInBreadcrumb:         in.HideInBreadcrumb == 1,
			HideChildrenInMenu:       in.HideChildrenInMenu == 1,
			Authority:                in.Authority,
			Badge:                    in.Badge,
			BadgeType:                in.BadgeType,
			BadgeVariants:            in.BadgeVariants,
			FullPathKey:              in.FullPathKey == 1,
			ActivePath:               in.ActivePath,
			AffixTab:                 in.AffixTab == 1,
			AffixTabOrder:            in.AffixTabOrder,
			IframeSrc:                in.IframeSrc,
			IgnoreAccess:             in.IgnoreAccess == 1,
			Link:                     in.Link,
			MaxNumOfOpenTab:          in.MaxNumOfOpenTab,
			MenuVisibleWithForbidden: in.MenuVisibleWithForbidden == 1,
			OpenInNewWindow:          in.OpenInNewWindow == 1,
			Order:                    in.Order,
			Query:                    in.Query,
			NoBasicLayout:            in.NoBasicLayout == 1,
		},

		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
