package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RuleAddReq struct {
	g.Meta `path:"/menu" tags:"菜单管理" method:"post" summary:"新增"`
	model.Author
	Pid       int64  `json:"pid"  v:"required|min:0" dc:"父级id(pid为0时, 代表顶级目录)"`
	Name      string `json:"name" v:"required#请填写规则名称" dc:"规则名称(唯一)"`
	Type      int64  `json:"type" v:"required|min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max" dc:"菜单类型(0:目录 1:菜单 2:按钮)"`
	Path      string `json:"path" dc:"菜单路径"`
	Component string `json:"component" dc:"组件路径"`

	Medata struct {
		Title                    string      `json:"title" dc:"用于配置页面的标题，会在菜单和标签页中显示"`
		Icon                     string      `json:"icon" dc:"用于配置页面的图标，会在菜单和标签页中显示"`
		ActiveIcon               string      `json:"activeIcon" dc:"用于配置页面的激活图标，会在菜单中显示(默认值:'')"`
		KeepAlive                bool        `json:"keepAlive" dc:"用于配置页面是否开启缓存，开启后页面会缓存，不会重新加载，仅在标签页启用时有效(默认值:false)"`
		HideInMenu               bool        `json:"hideInMenu" dc:"用于配置页面是否在菜单中隐藏，隐藏后页面不会在菜单中显示(默认值:false)"`
		HideInTab                bool        `json:"hideInTab" dc:"用于配置页面是否在标签页中隐藏，隐藏后页面不会在标签页中显示(默认值:false)"`
		HideInBreadcrumb         bool        `json:"hideInBreadcrumb" dc:"用于配置页面是否在面包屑中隐藏，隐藏后页面不会在面包屑中显示(默认值:false)"`
		HideChildrenInMenu       bool        `json:"hideChildrenInMenu" dc:"用于配置页面的子页面是否在菜单中隐藏，隐藏后子页面不会在菜单中显示(默认值:false)"`
		Authority                []string    `json:"authority" dc:"用于配置页面的权限，只有拥有对应权限的用户才能访问页面，不配置则不需要权限(默认值:'')"`
		Badge                    string      `json:"badge" dc:"用于配置页面的徽标，会在菜单显示(默认值:'')"`
		BadgeType                string      `json:"badgeType" dc:"用于配置页面的徽标类型(dot 为小红点, normal 为文本)"`
		BadgeVariants            string      `json:"badgeVariants" dc:"用于配置页面的徽标颜色('default' | 'destructive' | 'primary' | 'success' | 'warning' )"`
		FullPathKey              bool        `json:"fullPathKey" dc:"是否将路由的完整路径作为tab key(默认true)"`
		ActivePath               string      `json:"activePath" dc:"用于配置当前激活的菜单，有时候页面没有显示在菜单内，需要激活父级菜单时使用"`
		AffixTab                 bool        `json:"affixTab" dc:"用于配置页面是否固定标签页，固定后页面不可关闭(默认false)"`
		AffixTabOrder            int         `json:"affixTabOrder" dc:"用于配置页面固定标签页的排序, 采用升序排序(默认0)"`
		IframeSrc                string      `json:"iframeSrc" dc:"用于配置内嵌页面的 iframe 地址，设置后会在当前页面内嵌对应的页面(默认值:'')"`
		IgnoreAccess             bool        `json:"ignoreAccess" dc:"用于配置页面是否忽略权限，直接可以访问(默认false)"`
		Link                     string      `json:"link" dc:"用于配置外链跳转路径，会在新窗口打开(默认值:'')"`
		MaxNumOfOpenTab          int         `json:"maxNumOfOpenTab" dc:"用于配置标签页最大打开数量，设置后会在打开新标签页时自动关闭最早打开的标签页(默认值-1)"`
		MenuVisibleWithForbidden bool        `json:"menuVisibleWithForbidden" dc:"用于配置页面在菜单可以看到,但是访问会被重定向到403(默认值:false)"`
		OpenInNewWindow          bool        `json:"openInNewWindow" dc:"设置为 true 时，会在新窗口打开页面(默认值:false)"`
		Order                    int         `json:"order" dc:"排序(默认值:0)"`
		Query                    interface{} `json:"query" dc:"查询参数"`
		NoBasicLayout            bool        `json:"noBasicLayout" dc:"用于配置当前路由不使用基础布局，仅在顶级时生效(默认值:false)"`
	} `json:"meta"`
}

type RuleAddRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"菜单ID"`
}

type RuleDeleteReq struct {
	g.Meta `path:"/menu" tags:"菜单管理" method:"delete" summary:"删除"`
	model.Author
	Ids []int64 `json:"ids" v:"required#菜单ID列表必须" dc:"菜单ID列表"`
}

type RuleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type RuleUpdateReq struct {
	g.Meta `path:"/menu/{id}" tags:"菜单管理" method:"put" summary:"更新(全量更新)"`
	model.Author
	ID        int64  `p:"id" v:"required#id必须"`
	Pid       int64  `json:"pid"  v:"required|min:0" dc:"父级id"`
	Name      string `json:"name" v:"required#请填写规则名称" dc:"规则名称(唯一)"`
	Type      int64  `json:"type" v:"required|min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max" dc:"菜单类型(0:目录 1:菜单 2:按钮)"`
	Path      string `json:"path" dc:"菜单路径"`
	Component string `json:"component" dc:"组件路径"`

	Medata struct {
		Title                    string      `json:"title" dc:"用于配置页面的标题，会在菜单和标签页中显示"`
		Icon                     string      `json:"icon" dc:"用于配置页面的图标，会在菜单和标签页中显示"`
		ActiveIcon               string      `json:"activeIcon" dc:"用于配置页面的激活图标，会在菜单中显示(默认值:'')"`
		KeepAlive                bool        `json:"keepAlive" dc:"用于配置页面是否开启缓存，开启后页面会缓存，不会重新加载，仅在标签页启用时有效(默认值:false)"`
		HideInMenu               bool        `json:"hideInMenu" dc:"用于配置页面是否在菜单中隐藏，隐藏后页面不会在菜单中显示(默认值:false)"`
		HideInTab                bool        `json:"hideInTab" dc:"用于配置页面是否在标签页中隐藏，隐藏后页面不会在标签页中显示(默认值:false)"`
		HideInBreadcrumb         bool        `json:"hideInBreadcrumb" dc:"用于配置页面是否在面包屑中隐藏，隐藏后页面不会在面包屑中显示(默认值:false)"`
		HideChildrenInMenu       bool        `json:"hideChildrenInMenu" dc:"用于配置页面的子页面是否在菜单中隐藏，隐藏后子页面不会在菜单中显示(默认值:false)"`
		Authority                []string    `json:"authority" dc:"用于配置页面的权限，只有拥有对应权限的用户才能访问页面，不配置则不需要权限(默认值:'')"`
		Badge                    string      `json:"badge" dc:"用于配置页面的徽标，会在菜单显示(默认值:'')"`
		BadgeType                string      `json:"badgeType" dc:"用于配置页面的徽标类型(dot 为小红点, normal 为文本)"`
		BadgeVariants            string      `json:"badgeVariants" dc:"用于配置页面的徽标颜色('default' | 'destructive' | 'primary' | 'success' | 'warning' )"`
		FullPathKey              bool        `json:"fullPathKey" dc:"是否将路由的完整路径作为tab key(默认true)"`
		ActivePath               string      `json:"activePath" dc:"用于配置当前激活的菜单，有时候页面没有显示在菜单内，需要激活父级菜单时使用"`
		AffixTab                 bool        `json:"affixTab" dc:"用于配置页面是否固定标签页，固定后页面不可关闭(默认false)"`
		AffixTabOrder            int         `json:"affixTabOrder" dc:"用于配置页面固定标签页的排序, 采用升序排序(默认0)"`
		IframeSrc                string      `json:"iframeSrc" dc:"用于配置内嵌页面的 iframe 地址，设置后会在当前页面内嵌对应的页面(默认值:'')"`
		IgnoreAccess             bool        `json:"ignoreAccess" dc:"用于配置页面是否忽略权限，直接可以访问(默认false)"`
		Link                     string      `json:"link" dc:"用于配置外链跳转路径，会在新窗口打开(默认值:'')"`
		MaxNumOfOpenTab          int         `json:"maxNumOfOpenTab" dc:"用于配置标签页最大打开数量，设置后会在打开新标签页时自动关闭最早打开的标签页(默认值-1)"`
		MenuVisibleWithForbidden bool        `json:"menuVisibleWithForbidden" dc:"用于配置页面在菜单可以看到,但是访问会被重定向到403(默认值:false)"`
		OpenInNewWindow          bool        `json:"openInNewWindow" dc:"设置为 true 时，会在新窗口打开页面(默认值:false)"`
		Order                    int         `json:"order" dc:"排序(默认值:0)"`
		Query                    interface{} `json:"query" dc:"查询参数"`
		NoBasicLayout            bool        `json:"noBasicLayout" dc:"用于配置当前路由不使用基础布局，仅在顶级时生效(默认值:false)"`
	} `json:"meta"`
}

type RuleUpdateRes struct {
	g.Meta `mime:"application/json"`
}

type RuleGetReq struct {
	g.Meta `path:"/menu/{id}" tags:"菜单管理" method:"get" summary:"获取"`
	model.Author
	ID int64 `p:"id" v:"required#id必须"`
}

type RuleGetRes struct {
	g.Meta `mime:"application/json"`
	*model.AuthRule
}

// TODO fix
type RuleListReq struct {
	g.Meta `path:"/menu/list" tags:"菜单管理" method:"get" summary:"列表"`
	model.Author
}

type RuleListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.AuthRule `json:"list"`
}

type RuleGetTreeReq struct {
	g.Meta `path:"/menu/{id}/tree" tags:"菜单管理" method:"get" summary:"获取(树形结构)"`
	model.Author
}

type RuleGetTreeRes struct {
	g.Meta `mime:"application/json"`
	*model.AuthRuleNode
}

// TODO: need to modify /menu to /menu/tree
type RuleListTreeReq struct {
	g.Meta `path:"/menu" tags:"菜单管理" method:"get" summary:"列表(树形结构)"`
	model.Author
}

type RuleListTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.AuthRuleNode `json:"list"`
}
