package system

import (
	commonApi "IdentifyService/api/v1/common"
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"系统后台/认证模块" method:"post" summary:"用户登录"`
	UserName string `p:"user_name" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	MenuList    []*model.UserMenus  `json:"menuList"`
	Permissions []string            `json:"permissions"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"系统后台/认证模块" method:"get" summary:"退出登录"`
	commonApi.Author
}

type UserLoginOutRes struct {
}

type UserLogin2Req struct {
	g.Meta    `path:"/login2" tags:"系统后台/认证模块" method:"post" summary:"展会用户登录"`
	Phone     string `p:"phone" dc:"手机号"`
	Code      string `p:"code" dc:"验证码"`
	LoginType int    `p:"loginType" v:"required#登录类型不能为空" dc:"1:手机号+密码 2:手机号+验证码 3:IUQT SSO"`
}

type UserLogin2Res struct {
	g.Meta     `mime:"application/json"`
	Token      string            `json:"token"`
	UserInfo   *UserInfo2        `json:"user_info"`
	SettleInfo *model.SettleInfo `json:"settle_status" dc:"入驻状态(未入驻、审核中、已入驻)"`
}

type UserInfo2 struct {
	UserID   string `json:"user_id"`
	IUQTID   string `json:"iuqt_id"`
	UserName string `json:"user_name"`
	Mobile   string `json:"mobile"`
	UserType string `json:"user_type"`
}

type TokenIntrospectReq struct {
	g.Meta `path:"/token/introspect" tags:"系统后台/认证模块" method:"post" summary:"令牌内省"`
	commonApi.Author
}

type TokenIntrospectRes struct {
	g.Meta `mime:"application/json"`
	UserInfo2
}

type TokenRefreshReq struct {
	g.Meta `path:"/token/refresh" tags:"系统后台/认证模块" method:"post" summary:"刷新令牌"`
	commonApi.Author
}

type TokenRefreshRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token" dc:"新令牌"`
}
