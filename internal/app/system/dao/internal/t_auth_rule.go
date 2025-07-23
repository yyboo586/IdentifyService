package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRuleDao is the data access object for table t_auth_rule.
type AuthRuleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AuthRuleColumns // columns contains all the column names of Table for convenient usage.
}

// AuthRuleColumns defines and stores column names for table t_auth_rule.
type AuthRuleColumns struct {
	ID        string //
	Pid       string // 父ID
	Name      string // 规则名称
	Type      string
	Path      string // 路由地址
	Component string // 组件路径

	Title                    string // 规则名称
	Icon                     string // 图标
	ActiveIcon               string // 激活图标
	KeepAlive                string // 是否缓存
	HideInMenu               string // 是否在菜单中隐藏
	HideInTab                string // 是否在标签中隐藏
	HideInBreadcrumb         string // 是否在面包屑中隐藏
	HideChildrenInMenu       string // 是否在菜单中隐藏子菜单
	Authority                string // 作者
	Badge                    string // 徽章
	BadgeType                string // 徽章类型
	BadgeVariants            string // 页面的徽标颜色
	FullPathKey              string // 是否将路由的完整路径作为tab key(默认true)
	ActivePath               string // 用于配置当前激活的菜单，有时候页面没有显示在菜单内，需要激活父级菜单时使用
	AffixTab                 string // 用于配置页面是否固定标签页，固定后页面不可关闭
	AffixTabOrder            string // 用于配置页面固定标签页的顺序
	IframeSrc                string // 用于配置内嵌页面的 iframe 地址，设置后会在当前页面内嵌对应的页面
	IgnoreAccess             string // 用于配置页面是否忽略权限，直接可以访问
	Link                     string // 用于配置外链跳转路径，会在新窗口打开
	MaxNumOfOpenTab          string // 用于配置标签页最大打开数量，设置后会在打开新标签页时自动关闭最早打开的标签页(仅在打开同名标签页时生效)
	MenuVisibleWithForbidden string // 用于配置页面在菜单可以看到,但是访问会被重定向到403
	OpenInNewWindow          string // 设置为 true 时，会在新窗口打开页面
	Order                    string // 用于配置页面的排序，用于路由到菜单排序。
	Query                    string // 用于配置页面的菜单参数，会在菜单中传递给页面
	NoBasicLayout            string // 用于配置页面是否不使用基础布局，设置为 true 时，页面不会使用基础布局

	CreatedAt string // 创建日期
	UpdatedAt string // 修改日期
}

// authRuleColumns holds the columns for table t_auth_rule.
var authRuleColumns = AuthRuleColumns{
	ID:        "id",
	Pid:       "pid",
	Name:      "name",
	Type:      "type",
	Path:      "path",
	Component: "component",

	Title:                    "title",
	Icon:                     "icon",
	ActiveIcon:               "active_icon",
	KeepAlive:                "keep_alive",
	HideInMenu:               "hide_in_menu",
	HideInTab:                "hide_in_tab",
	HideInBreadcrumb:         "hide_in_breadcrumb",
	HideChildrenInMenu:       "hide_children_in_menu",
	Authority:                "authority",
	Badge:                    "badge",
	BadgeType:                "badge_type",
	BadgeVariants:            "badge_variants",
	FullPathKey:              "full_path_key",
	ActivePath:               "active_path",
	AffixTab:                 "affix_tab",
	AffixTabOrder:            "affix_tab_order",
	IframeSrc:                "iframe_src",
	IgnoreAccess:             "ignore_access",
	Link:                     "link",
	MaxNumOfOpenTab:          "max_num_of_open_tab",
	MenuVisibleWithForbidden: "menu_visible_with_forbidden",
	OpenInNewWindow:          "open_in_new_window",
	Order:                    "order",
	Query:                    "query",
	NoBasicLayout:            "no_basic_layout",

	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAuthRuleDao creates and returns a new DAO object for table data access.
func NewAuthRuleDao() *AuthRuleDao {
	return &AuthRuleDao{
		group:   "default",
		table:   "t_auth_rule",
		columns: authRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthRuleDao) Columns() AuthRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
