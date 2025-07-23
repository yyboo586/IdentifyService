package logics

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"context"
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

// ========== 基础CRUD操作 ==========

func (a *authRule) Add(ctx context.Context, req *system.RuleAddReq) (ruleID int64, err error) {
	data := a.buildAuthRuleData(req)

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

func (a *authRule) DeleteByIDs(ctx context.Context, ids []int64) (err error) {
	var list []*model.AuthRule
	list, err = a.getAllAuthRules(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 获取所有子节点ID
	childrenIds := make([]int64, 0, len(list))
	for _, id := range ids {
		rules := a.findChildrenByParentID(list, id)
		for _, child := range rules {
			childrenIds = append(childrenIds, child.ID)
		}
	}
	ids = append(ids, childrenIds...)

	_, err = dao.AuthRule.Ctx(ctx).WhereIn(dao.AuthRule.Columns().ID, ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 异步删除权限策略
	go func(ruleIds []int64) {
		for _, v := range ruleIds {
			_, err = a.enforcer.RemoveFilteredNamedPolicy("p", 1, gconv.String(v))
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}
	}(ids)

	return
}

func (a *authRule) Update(ctx context.Context, req *system.RuleUpdateReq) (err error) {
	data := a.buildAuthRuleData(req)

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

func (a *authRule) Get(ctx context.Context, id int64) (out *model.AuthRule, err error) {
	return
}

func (a *authRule) List(ctx context.Context) (out []*model.AuthRule, err error) {
	return
}

func (a *authRule) GetTree(ctx context.Context, id int64) (out *model.AuthRuleNode, err error) {
	return
}

func (a *authRule) ListTree(ctx context.Context) (out []*model.AuthRuleNode, err error) {
	return
}

// ========== 登录场景接口 ==========
func (a *authRule) GetMenuTreeByUserID(ctx context.Context, userID string) (out []*model.AuthRuleNode, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		return a.getFullMenuTree(ctx, false) // 不包含按钮
	}

	if userInstance.IsSuperAdmin(ctx, userID) {
		return a.getFullMenuTree(ctx, false) // 不包含按钮
	}

	roleIDs, err := roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return a.getMenuTreeByRoleIDs(ctx, roleIDs, false) // 不包含按钮
}

func (a *authRule) GetButtonListByUserID(ctx context.Context, userID string) (out []*model.AuthRule, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		return a.getAllButtons(ctx)
	}

	roleIDs, err := roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return a.getButtonListByRoleIDs(ctx, roleIDs)
}

// ========== 权限管理场景接口 ==========

func (a *authRule) GetFullAuthRuleTree(ctx context.Context, userID string) (out []*model.AuthRuleNode, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		return a.getFullMenuTree(ctx, true) // 包含按钮
	}

	roleIDs, err := roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return a.getMenuTreeByRoleIDs(ctx, roleIDs, true) // 包含按钮
}

// ========== 权限过滤工具 ==========

func (a *authRule) FilterRuleIDsByUserID(ctx context.Context, ruleIDs []int64, userID string) (out []int64, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		return ruleIDs, nil
	}

	accessRuleIDs, err := a.getRuleIDsByUserID(ctx, userID)
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

func (a *authRule) HasPermission(ctx context.Context, userID string, ruleID int64) (hasPermission bool, err error) {
	if userInstance.IsSuperAdmin(ctx, userID) {
		return true, nil
	}

	userRuleIDs, err := a.getRuleIDsByUserID(ctx, userID)
	if err != nil {
		return false, err
	}

	for _, id := range userRuleIDs {
		if id == ruleID {
			return true, nil
		}
	}
	return false, nil
}

// ========== 私有方法 ==========

// TODO List: 这里怎么才能不申请新的内存
// getFullMenuTree 获取完整的菜单树
func (a *authRule) getFullMenuTree(ctx context.Context, includeButton bool) (list []*model.AuthRuleNode, err error) {
	allAuthRules, err := a.getAllAuthRules(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var result []*model.AuthRule
	var rootNodeIDs []int64
	for _, v := range allAuthRules {
		// 过滤按钮
		if !includeButton && v.Type == model.MenuTypeButton {
			continue
		}

		result = append(result, v)
		if v.Pid == 0 {
			rootNodeIDs = append(rootNodeIDs, v.ID)
		}
	}

	// 构建树形结构
	for _, id := range rootNodeIDs {
		list = append(list, a.buildTree(id, result))
	}
	return
}

// getMenuTreeByRoleIDs 根据角色ID获取菜单树
func (a *authRule) getMenuTreeByRoleIDs(ctx context.Context, roleIDs []int64, includeButton bool) (trees []*model.AuthRuleNode, err error) {
	ruleIDs, err := a.getRuleIDsByRoleIDs(ctx, roleIDs)
	if err != nil {
		return
	}

	return a.buildTreeByRuleIDs(ctx, ruleIDs, includeButton)
}

// getAllButtons 获取所有按钮
func (a *authRule) getAllButtons(ctx context.Context) (out []*model.AuthRule, err error) {
	var entitys []*entity.AuthRule
	err = dao.AuthRule.Ctx(ctx).Where(dao.AuthRule.Columns().Type, model.MenuTypeButton).Scan(&entitys)
	if err != nil {
		return
	}

	for _, entity := range entitys {
		out = append(out, a.convertEntityToModel(entity))
	}
	return
}

// getButtonListByRoleIDs 根据角色ID获取按钮列表
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
		}
	}
	return
}

// buildTreeByRuleIDs 根据规则ID列表构建树
func (a *authRule) buildTreeByRuleIDs(ctx context.Context, ruleIDs []int64, includeButton bool) (trees []*model.AuthRuleNode, err error) {
	rulesMap := make(map[int64]bool)
	for _, v := range ruleIDs {
		rulesMap[v] = true
	}

	allAuthRules, err := a.getAllAuthRules(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var result []*model.AuthRule
	var rootNodeIDs []int64
	for _, v := range allAuthRules {
		if _, ok := rulesMap[v.ID]; !ok {
			continue
		}

		if !includeButton && v.Type == model.MenuTypeButton {
			continue
		}
		result = append(result, v)
		if v.Pid == 0 {
			rootNodeIDs = append(rootNodeIDs, v.ID)
		}
	}

	for _, id := range rootNodeIDs {
		trees = append(trees, a.buildTree(id, result))
	}

	return
}

// findChildrenByParentID 根据父ID查找子节点
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

// getRuleIDsByUserID 根据用户ID获取规则ID列表
// 1、根据用户ID获取角色ID列表
// 2、根据角色ID列表获取规则ID列表
func (a *authRule) getRuleIDsByUserID(ctx context.Context, userID string) (ruleIDs []int64, err error) {
	roleIDs, err := roleInstance.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		return
	}
	return a.getRuleIDsByRoleIDs(ctx, roleIDs)
}

// getRuleIDsByRoleIDs 根据角色ID列表获取规则ID列表
func (a *authRule) getRuleIDsByRoleIDs(ctx context.Context, roleIDs []int64) (ruleIDs []int64, err error) {
	for _, roleID := range roleIDs {
		var policies [][]string
		policies, err = a.enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(roleID))
		if err != nil {
			return
		}
		for _, policy := range policies {
			ruleIDs = append(ruleIDs, gconv.Int64(policy[1]))
		}
	}

	ruleIDs = libUtils.SliceUnique(ruleIDs)
	return
}

// buildTree 构建树形结构
// 统一的树形构建方法，在内存中递归构建
func (a *authRule) buildTree(rootID int64, rules []*model.AuthRule) (root *model.AuthRuleNode) {
	// 查找根节点
	for _, rule := range rules {
		if rule.ID == rootID {
			root = &model.AuthRuleNode{
				AuthRule: rule,
			}
			break
		}
	}

	if root == nil {
		return
	}

	// 递归构建子树
	for _, v := range rules {
		if v.Pid == rootID {
			root.Children = append(root.Children, a.buildTree(v.ID, rules))
		}
	}
	return
}

// getAllAuthRules 获取所有权限规则
func (a *authRule) getAllAuthRules(ctx context.Context) (list []*model.AuthRule, err error) {
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

// convertEntityToModel 实体转模型
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

// buildAuthRuleData 构建AuthRule数据
func (a *authRule) buildAuthRuleData(req interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	switch r := req.(type) {
	case *system.RuleAddReq:
		data[dao.AuthRule.Columns().Pid] = r.Pid
		data[dao.AuthRule.Columns().Name] = r.Name
		data[dao.AuthRule.Columns().Type] = r.Type
		data[dao.AuthRule.Columns().Path] = r.Path
		data[dao.AuthRule.Columns().Component] = r.Component
		a.setMetaData(data, r)
	case *system.RuleUpdateReq:
		data[dao.AuthRule.Columns().Pid] = r.Pid
		data[dao.AuthRule.Columns().Name] = r.Name
		data[dao.AuthRule.Columns().Type] = r.Type
		data[dao.AuthRule.Columns().Path] = r.Path
		data[dao.AuthRule.Columns().Component] = r.Component
		a.setMetaData(data, r)
	}

	return data
}

// setMetaData 设置元数据
func (a *authRule) setMetaData(data map[string]interface{}, meta interface{}) {
	switch m := meta.(type) {
	case *system.RuleAddReq:
		data[dao.AuthRule.Columns().Title] = m.Medata.Title
		data[dao.AuthRule.Columns().Icon] = m.Medata.Icon
		data[dao.AuthRule.Columns().ActiveIcon] = m.Medata.ActiveIcon
		data[dao.AuthRule.Columns().Authority] = m.Medata.Authority
		data[dao.AuthRule.Columns().Badge] = m.Medata.Badge
		data[dao.AuthRule.Columns().BadgeType] = m.Medata.BadgeType
		data[dao.AuthRule.Columns().BadgeVariants] = m.Medata.BadgeVariants
		data[dao.AuthRule.Columns().ActivePath] = m.Medata.ActivePath
		data[dao.AuthRule.Columns().AffixTabOrder] = m.Medata.AffixTabOrder
		data[dao.AuthRule.Columns().IframeSrc] = m.Medata.IframeSrc
		data[dao.AuthRule.Columns().Link] = m.Medata.Link
		data[dao.AuthRule.Columns().MaxNumOfOpenTab] = m.Medata.MaxNumOfOpenTab
		data[dao.AuthRule.Columns().Order] = m.Medata.Order
		data[dao.AuthRule.Columns().Query] = m.Medata.Query

		// 布尔值转换
		data[dao.AuthRule.Columns().KeepAlive] = a.boolToInt(m.Medata.KeepAlive)
		data[dao.AuthRule.Columns().HideInMenu] = a.boolToInt(m.Medata.HideInMenu)
		data[dao.AuthRule.Columns().HideInTab] = a.boolToInt(m.Medata.HideInTab)
		data[dao.AuthRule.Columns().HideInBreadcrumb] = a.boolToInt(m.Medata.HideInBreadcrumb)
		data[dao.AuthRule.Columns().HideChildrenInMenu] = a.boolToInt(m.Medata.HideChildrenInMenu)
		data[dao.AuthRule.Columns().FullPathKey] = a.boolToInt(m.Medata.FullPathKey)
		data[dao.AuthRule.Columns().AffixTab] = a.boolToInt(m.Medata.AffixTab)
		data[dao.AuthRule.Columns().IgnoreAccess] = a.boolToInt(m.Medata.IgnoreAccess)
		data[dao.AuthRule.Columns().MenuVisibleWithForbidden] = a.boolToInt(m.Medata.MenuVisibleWithForbidden)
		data[dao.AuthRule.Columns().OpenInNewWindow] = a.boolToInt(m.Medata.OpenInNewWindow)
		data[dao.AuthRule.Columns().NoBasicLayout] = a.boolToInt(m.Medata.NoBasicLayout)
	case *system.RuleUpdateReq:
		data[dao.AuthRule.Columns().Title] = m.Medata.Title
		data[dao.AuthRule.Columns().Icon] = m.Medata.Icon
		data[dao.AuthRule.Columns().ActiveIcon] = m.Medata.ActiveIcon
		data[dao.AuthRule.Columns().Authority] = m.Medata.Authority
		data[dao.AuthRule.Columns().Badge] = m.Medata.Badge
		data[dao.AuthRule.Columns().BadgeType] = m.Medata.BadgeType
		data[dao.AuthRule.Columns().BadgeVariants] = m.Medata.BadgeVariants
		data[dao.AuthRule.Columns().ActivePath] = m.Medata.ActivePath
		data[dao.AuthRule.Columns().AffixTabOrder] = m.Medata.AffixTabOrder
		data[dao.AuthRule.Columns().IframeSrc] = m.Medata.IframeSrc
		data[dao.AuthRule.Columns().Link] = m.Medata.Link
		data[dao.AuthRule.Columns().MaxNumOfOpenTab] = m.Medata.MaxNumOfOpenTab
		data[dao.AuthRule.Columns().Order] = m.Medata.Order
		data[dao.AuthRule.Columns().Query] = m.Medata.Query

		// 布尔值转换
		data[dao.AuthRule.Columns().KeepAlive] = a.boolToInt(m.Medata.KeepAlive)
		data[dao.AuthRule.Columns().HideInMenu] = a.boolToInt(m.Medata.HideInMenu)
		data[dao.AuthRule.Columns().HideInTab] = a.boolToInt(m.Medata.HideInTab)
		data[dao.AuthRule.Columns().HideInBreadcrumb] = a.boolToInt(m.Medata.HideInBreadcrumb)
		data[dao.AuthRule.Columns().HideChildrenInMenu] = a.boolToInt(m.Medata.HideChildrenInMenu)
		data[dao.AuthRule.Columns().FullPathKey] = a.boolToInt(m.Medata.FullPathKey)
		data[dao.AuthRule.Columns().AffixTab] = a.boolToInt(m.Medata.AffixTab)
		data[dao.AuthRule.Columns().IgnoreAccess] = a.boolToInt(m.Medata.IgnoreAccess)
		data[dao.AuthRule.Columns().MenuVisibleWithForbidden] = a.boolToInt(m.Medata.MenuVisibleWithForbidden)
		data[dao.AuthRule.Columns().OpenInNewWindow] = a.boolToInt(m.Medata.OpenInNewWindow)
		data[dao.AuthRule.Columns().NoBasicLayout] = a.boolToInt(m.Medata.NoBasicLayout)
	}
}

// boolToInt 布尔值转整数
func (a *authRule) boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
