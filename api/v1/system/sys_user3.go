package system

import (
	commonApi "IdentifyService/api/v1/common"
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UnRegisterReq struct {
	g.Meta `path:"/users/{user_id}/unregister" tags:"个人中心" method:"post" summary:"注销账户"`
	commonApi.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
}

type UnRegisterRes struct {
	g.Meta `mime:"application/json"`
	commonApi.EmptyRes
}

type GetUserPersonalInfoReq struct {
	g.Meta `path:"/users/{user_id}/personal-info" tags:"个人中心" method:"get" summary:"获取用户个人信息"`
	model.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
}

type GetUserPersonalInfoRes struct {
	g.Meta `mime:"application/json"`
	*UserPersonalInfo
}

type EditUserPersonalInfoReq struct {
	g.Meta `path:"/users/{user_id}/personal-info" tags:"个人中心" method:"put" summary:"编辑用户个人信息(全量更新)"`
	model.Author
	UserID   string `p:"user_id" v:"required#用户ID不能为空"`
	Nickname string `json:"nickname" dc:"用户昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Sex      int    `json:"sex" dc:"性别(0:保密,1:男,2:女)"`
	Birthday string `json:"birthday" dc:"生日(yyyy-mm-dd)"`
	City     string `json:"city" dc:"城市"`
}

type EditUserPersonalInfoRes struct {
	model.EmptyRes
}

type UserPersonalInfo struct {
	Nickname          string `json:"nickname" dc:"用户昵称"`
	Avatar            string `json:"avatar" dc:"头像"`
	Mobile            string `json:"mobile" dc:"手机号"`
	Sex               int    `json:"sex" dc:"性别(0:保密,1:男,2:女)"`
	Birthday          string `json:"birthday" dc:"生日(yyyy-mm-dd)"`
	City              string `json:"city" dc:"城市"`
	CreateTime        string `json:"create_time" dc:"创建时间"`
	IsAlreadyRealname bool   `json:"is_already_realname" dc:"是否已经实名认证"`
	UserRealName      string `json:"user_real_name" dc:"用户真实姓名"`
}

type BindPhoneReq struct {
	g.Meta `path:"/users/{user_id}/bind-phone" tags:"个人中心" method:"put" summary:"绑定手机号"`
	commonApi.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
	Phone  string `p:"phone" v:"required#手机号不能为空"`
	Code   string `p:"code" v:"required#验证码不能为空"`
}

type BindPhoneRes struct {
	model.EmptyRes
}

type EditUserPasswordReq struct {
	g.Meta `path:"/users/{user_id}/password" tags:"个人中心" method:"put" summary:"修改密码"`
	commonApi.Author
	UserID      string `p:"user_id" v:"required#用户ID不能为空"`
	Phone       string `p:"phone" v:"required#手机号不能为空"`
	NewPassword string `p:"new_password" v:"required#新密码不能为空"`
}

type EditUserPasswordRes struct {
	model.EmptyRes
}

type EditUserIDCardReq struct {
	g.Meta `path:"/users/{user_id}/id-card" tags:"个人中心" method:"put" summary:"绑定身份证号"`
	commonApi.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
	// 证件类型(1:居民身份证,2:港澳居民来往内地通行证,3:台湾居民来往大陆通行证)
	CardType string `p:"card_type" v:"required|in:居民身份证,港澳居民往来内地通行证,台湾居民往来大陆通行证#证件类型不正确"`
	IDCard   string `p:"id_card" v:"required#证件号不能为空"`
	RealName string `p:"real_name" v:"required#真实姓名不能为空"`
}

type EditUserIDCardRes struct {
	model.EmptyRes
}

type GetUserIDCardReq struct {
	g.Meta `path:"/users/{user_id}/id-card" tags:"个人中心" method:"get" summary:"获取身份证号"`
	commonApi.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
}

type GetUserIDCardRes struct {
	g.Meta   `mime:"application/json"`
	CardType string `json:"card_type"`
	IDCard   string `json:"id_card"`
	RealName string `json:"real_name"`
}

type Add3Req struct {
	g.Meta `path:"/users" tags:"账户管理" method:"post" summary:"添加用户"`
	model.Author
	Phone        string `p:"phone" v:"required#手机号不能为空"`
	UserNickname string `p:"user_nickname" v:"required#用户昵称不能为空"`
}

type Add3Res struct {
	g.Meta `mime:"application/json"`
	UserID string `json:"user_id"`
}

type Delete3Req struct {
	g.Meta `path:"/users/{user_id}" tags:"账户管理" method:"delete" summary:"删除用户"`
	model.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
}

type Delete3Res struct {
	g.Meta `mime:"application/json"`
	model.EmptyRes
}

type ListUserDeviceReq struct {
	g.Meta `path:"/users/{user_id}/device" tags:"个人中心/设备管理" method:"get" summary:"获取用户设备列表"`
	model.Author
	UserID string `p:"user_id" v:"required#用户ID不能为空"`
	model.PageReq
}

type ListUserDeviceRes struct {
	g.Meta `mime:"application/json"`
	List   []*UserDeviceItem `json:"list"`
	*model.PageRes
}

type DeleteUserDeviceReq struct {
	g.Meta `path:"/users/{user_id}/device/{device_id}" tags:"个人中心/设备管理" method:"delete" summary:"删除用户设备"`
	model.Author
	UserID   string `p:"user_id" v:"required#用户ID不能为空"`
	DeviceID string `p:"device_id" v:"required#设备ID不能为空"`
}

type DeleteUserDeviceRes struct {
	g.Meta `mime:"application/json"`
	model.EmptyRes
}

type UserDeviceItem struct {
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	DeviceIP   string `json:"device_ip"`
	LoginType  string `json:"login_type"`
	CreateTime string `json:"create_time"`
}
