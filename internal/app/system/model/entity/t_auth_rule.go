package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRule is the golang structure for table t_auth_rule.
type AuthRule struct {
	ID        int64  `orm:"id"`
	Pid       int64  `orm:"pid"`
	Name      string `orm:"name"`
	Type      int    `orm:"type"`
	Path      string `orm:"path"`
	Component string `orm:"component"`

	Title                    string `orm:"title"`
	Icon                     string `orm:"icon"`
	ActiveIcon               string `orm:"active_icon"`
	KeepAlive                int    `orm:"keep_alive"`
	HideInMenu               int    `orm:"hide_in_menu"`
	HideInTab                int    `orm:"hide_in_tab"`
	HideInBreadcrumb         int    `orm:"hide_in_breadcrumb"`
	HideChildrenInMenu       int    `orm:"hide_children_in_menu"`
	Authority                string `orm:"authority"`
	Badge                    string `orm:"badge"`
	BadgeType                string `orm:"badge_type"`
	BadgeVariants            string `orm:"badge_variants"`
	FullPathKey              int    `orm:"full_path_key"`
	ActivePath               string `orm:"active_path"`
	AffixTab                 int    `orm:"affix_tab"`
	AffixTabOrder            int64  `orm:"affix_tab_order"`
	IframeSrc                string `orm:"iframe_src"`
	IgnoreAccess             int    `orm:"ignore_access"`
	Link                     string `orm:"link"`
	MaxNumOfOpenTab          int64  `orm:"max_num_of_open_tab"`
	MenuVisibleWithForbidden int    `orm:"menu_visible_with_forbidden"`
	OpenInNewWindow          int    `orm:"open_in_new_window"`
	Order                    int64  `orm:"order"`
	Query                    string `orm:"query"`
	NoBasicLayout            int    `orm:"no_basic_layout"`

	CreatedAt *gtime.Time `orm:"created_at"  description:"创建日期"`
	UpdatedAt *gtime.Time `orm:"updated_at"  description:"修改日期"`
}
