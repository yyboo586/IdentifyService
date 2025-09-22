package system

import (
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserRegistrationReq struct {
	g.Meta   `path:"/auth/registration" tags:"认证管理" method:"post" summary:"自注册账户"`
	Username string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type UserRegistrationRes struct {
	g.Meta `mime:"application/json"`
}

type UserLoginReq struct {
	g.Meta   `path:"/auth/login" tags:"认证管理" method:"post" summary:"账户登入"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	g.Meta               `mime:"application/json"`
	UserInfo             *model.UserLoginRes         `json:"user_info"`
	Token                string                      `json:"token"`
	PermissionPointCodes []model.PermissionPointCode `json:"permission_point_codes"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/auth/logout" tags:"认证管理" method:"post" summary:"账户登出"`
	model.Author
}

type UserLoginOutRes struct {
	g.Meta `mime:"application/json"`
}

type IntrospectReq struct {
	g.Meta `path:"/auth/token/introspect" method:"POST" tags:"认证管理" summary:"令牌内省"`
	model.Author
}
