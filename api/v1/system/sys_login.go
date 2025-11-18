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
	g.Meta    `path:"/login2" tags:"认证模块" method:"post" summary:"用户登录(用户不存在时自动注册)"`
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

type UserLogin3Req struct {
	g.Meta        `path:"/login3" tags:"认证模块" method:"post" summary:"用户登录(用户不存在时自动注册)"`
	Phone         string      `json:"phone" dc:"手机号" v:"required#手机号不能为空"`
	Code          string      `json:"code" dc:"验证码/密码" v:"required#code不能为空"`
	BussinessType string      `json:"bussiness_type" v:"required#业务类型不能为空" dc:"业务类型(验证码登录、密码登录)"`
	Aggrement     *Aggrement  `json:"aggrement" dc:"用户协议" v:"required#用户协议不能为空"`
	DeviceInfo    *DeviceInfo `json:"device_info" dc:"设备信息" v:"required#设备信息不能为空"`
}

type DeviceInfo struct {
	DeviceType string `json:"device_type" v:"required|in:android,ios,web#设备类型不能为空" dc:"设备类型(android, ios, web)"`
	DeviceID   string `json:"device_id" v:"required#设备ID不能为空" dc:"设备ID"`
	DeviceName string `json:"device_name" v:"required#设备名称不能为空" dc:"设备名称"`
}

type Aggrement struct {
	ID        int64  `json:"id" dc:"协议ID"`
	Name      string `json:"name" v:"required#协议名称不能为空" dc:"协议名称"`
	Accepted  bool   `json:"accepted" v:"required#是否接受不能为空" dc:"是否接受"`
	Version   string `json:"version" v:"required#协议版本不能为空" dc:"协议版本标识符"`
	Timestamp int64  `json:"timestamp" v:"required#用户操作时间戳不能为空" dc:"用户操作时间戳(单位:秒)"`
}

type UserLogin3Res struct {
	g.Meta       `mime:"application/json"`
	Token        string `json:"token"`
	UserID       string `json:"user_id"`
	UserNickname string `json:"user_nickname"`
	Avatar       string `json:"avatar"`
	Phone        string `json:"phone"`
	Sex          int    `json:"sex"`
	Birthday     string `json:"birthday"`
	City         string `json:"city"`
}

type UserInfo3 struct {
	ID          string `json:"id" dc:"用户ID"`
	Avatar      string `json:"avatar" dc:"头像"`
	UserNickame string `json:"user_nickname" dc:"用户昵称"`
	Phone       string `json:"phone" dc:"手机号"`
	Sex         int    `json:"sex" dc:"性别(0:保密,1:男,2:女)"`
	Birthday    string `json:"birthday" dc:"生日"`
	City        string `json:"city" dc:"城市"`
}

type UserInfo2 struct {
	UserID   string `json:"user_id"`
	IUQTID   string `json:"iuqt_id"`
	UserName string `json:"user_name"`
	Mobile   string `json:"mobile"`
	UserType string `json:"user_type"`
	DeptID   uint64 `json:"dept_id"`
}

type TokenIntrospectReq struct {
	g.Meta `path:"/token/introspect" tags:"系统后台/认证模块" method:"post" summary:"令牌内省"`
	model.Author
}

type TokenIntrospectRes struct {
	g.Meta `mime:"application/json"`
	UserInfo2
}

type RevokeTokenReq struct {
	g.Meta `path:"/token/revoke" tags:"系统后台/认证模块" method:"post" summary:"撤销令牌"`
	model.Author
}

type RevokeTokenRes struct {
	g.Meta `mime:"application/json"`
	model.EmptyRes
}
