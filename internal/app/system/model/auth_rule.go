package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AuthRuleType int

const (
	MenuTypeDirectory AuthRuleType = iota // 目录
	MenuTypeMenu                          // 菜单
	MenuTypeButton                        // 按钮
)

type AuthRule struct {
	ID        int64        `json:"id"`        //
	Pid       int64        `json:"pid"`       // 父ID
	Name      string       `json:"name"`      // 规则名称
	Type      AuthRuleType `json:"type"`      // 类型 0目录 1菜单 2按钮
	Path      string       `json:"path"`      // 路由地址
	Component string       `json:"component"` // 组件路径

	MMeta *AuthRuleMeta `json:"meta"`

	CreatedAt *gtime.Time `json:"createdAt"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt"` // 更新时间
}

type AuthRuleMeta struct {
	Title                    string `json:"title"`                    // 规则名称
	Icon                     string `json:"icon"`                     // 图标
	ActiveIcon               string `json:"activeIcon"`               // 激活图标
	KeepAlive                bool   `json:"keepAlive"`                // 是否缓存
	HideInMenu               bool   `json:"hideInMenu"`               // 是否在菜单中隐藏
	HideInTab                bool   `json:"hideInTab"`                // 是否在标签中隐藏
	HideInBreadcrumb         bool   `json:"hideInBreadcrumb"`         // 是否在面包屑中隐藏
	HideChildrenInMenu       bool   `json:"hideChildrenInMenu"`       // 是否在菜单中隐藏子菜单
	Authority                string `json:"authority"`                // 作者
	Badge                    string `json:"badge"`                    // 徽章
	BadgeType                string `json:"badgeType"`                // 徽章类型
	BadgeVariants            string `json:"badgeVariants"`            // 页面的徽标颜色
	FullPathKey              bool   `json:"fullPathKey"`              // 是否将路由的完整路径作为tab key(默认true)
	ActivePath               string `json:"activePath"`               // 用于配置当前激活的菜单，有时候页面没有显示在菜单内，需要激活父级菜单时使用
	AffixTab                 bool   `json:"affixTab"`                 // 用于配置页面是否固定标签页，固定后页面不可关闭
	AffixTabOrder            int64  `json:"affixTabOrder"`            // 用于配置页面固定标签页的顺序
	IframeSrc                string `json:"iframeSrc"`                // 用于配置内嵌页面的 iframe 地址，设置后会在当前页面内嵌对应的页面
	IgnoreAccess             bool   `json:"ignoreAccess"`             // 用于配置页面是否忽略权限，直接可以访问
	Link                     string `json:"link"`                     // 用于配置外链跳转路径，会在新窗口打开
	MaxNumOfOpenTab          int64  `json:"maxNumOfOpenTab"`          // 用于配置标签页最大打开数量，设置后会在打开新标签页时自动关闭最早打开的标签页(仅在打开同名标签页时生效)
	MenuVisibleWithForbidden bool   `json:"menuVisibleWithForbidden"` // 用于配置页面在菜单可以看到,但是访问会被重定向到403
	OpenInNewWindow          bool   `json:"openInNewWindow"`          // 设置为 true 时，会在新窗口打开页面
	Order                    int64  `json:"order"`                    // 用于配置页面的排序，用于路由到菜单排序。
	Query                    string `json:"query"`                    // 用于配置页面的菜单参数，会在菜单中传递给页面
	NoBasicLayout            bool   `json:"noBasicLayout"`            // 用于配置页面是否不使用基础布局，设置为 true 时，页面不会使用基础布局
}

// AuthRuleTree 菜单树形结构
type AuthRuleNode struct {
	*AuthRule `json:""`
	Children  []*AuthRuleNode `json:"children"`
}
