package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"认证管理" method:"post" summary:"账户登入"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	g.Meta     `mime:"application/json"`
	UserInfo   *model.UserLoginRes   `json:"user_info"`
	Token      string                `json:"token"`
	MenuList   []*model.AuthRuleNode `json:"menu_list"`
	ButtonList []*Button             `json:"button_list"`
}

type Button struct {
	ID   int64  `json:"id"`
	Pid  int64  `json:"pid"`
	Name string `json:"name"`
	Type int64  `json:"type"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/user/logout" tags:"认证管理" method:"get" summary:"账户登出"`
	model.Author
}

type UserLoginOutRes struct {
	g.Meta `mime:"application/json"`
}

type IntrospectReq struct {
	g.Meta `path:"/token/introspect" method:"POST" tags:"认证管理" summary:"令牌内省"`
	Token  string `json:"token"`
}
