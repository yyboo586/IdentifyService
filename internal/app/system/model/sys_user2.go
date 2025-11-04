package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type LoginType int

const (
	LoginTypeUnknown       LoginType = iota
	LoginTypePhonePassword           // 手机号+密码
	LoginTypePhoneSMSCode            // 手机号+验证码
	LoginTypeSSO                     // IUQT SSO
)

type UserStatus int

const (
	UserStatusDisabled   UserStatus = iota // 禁用
	UserStatusNormal                       // 正常
	UserStatusUnverified                   // 未验证
)

func GetUserStatus(status UserStatus) string {
	switch status {
	case UserStatusDisabled:
		return "禁用"
	case UserStatusNormal:
		return "正常"
	case UserStatusUnverified:
		return "未验证"
	default:
		return "未知"
	}
}

type UserType int

const (
	UserTypeUnknown         UserType = iota // 未知
	UserTypeServiceProvider                 // 服务提供商
	UserTypeExhibitor                       // 展商
)

func GetUserType(userType string) UserType {
	switch userType {
	case "服务提供商":
		return UserTypeServiceProvider
	case "展商":
		return UserTypeExhibitor
	default:
		return UserTypeUnknown
	}
}

func GetUserTypeText(userType UserType) string {
	switch userType {
	case UserTypeServiceProvider:
		return "服务提供商"
	case UserTypeExhibitor:
		return "展商"
	default:
		return "未知"
	}
}

type User struct {
	ID            int64       `json:"id"`
	UserName      string      `json:"user_name"`
	Mobile        string      `json:"mobile"`
	UserNickname  string      `json:"user_nickname"`
	Birthday      int         `json:"birthday"`
	UserPassword  string      `json:"user_password"`
	UserSalt      string      `json:"user_salt"`
	UserStatus    UserStatus  `json:"user_status"`
	UserEmail     string      `json:"user_email"`
	Sex           int         `json:"sex"`
	Avatar        string      `json:"avatar"`
	DeptId        uint64      `json:"dept_id"`
	Remark        string      `json:"remark"`
	IsAdmin       bool        `json:"is_admin"`
	Address       string      `json:"address"`
	Describe      string      `json:"describe"`
	LastLoginIp   string      `json:"last_login_ip"`
	LastLoginTime *gtime.Time `json:"last_login_time"`
	CreatedAt     *gtime.Time `json:"created_at"`
	UpdatedAt     *gtime.Time `json:"updated_at"`
	DeletedAt     *gtime.Time `json:"deleted_at"`
	OpenId        string      `json:"openId"`
	IUQTID        string      `json:"iuqt_id"`
	UserType      UserType    `json:"user_type"`
}

func ConvertToUser(user *entity.SysUser) *User {
	return &User{
		ID:            int64(user.Id),
		IUQTID:        user.IUQTID,
		UserName:      user.UserName,
		Mobile:        user.Mobile,
		UserNickname:  user.UserNickname,
		Birthday:      user.Birthday,
		UserPassword:  user.UserPassword,
		UserSalt:      user.UserSalt,
		UserStatus:    UserStatus(user.UserStatus),
		UserEmail:     user.UserEmail,
		Sex:           user.Sex,
		Avatar:        user.Avatar,
		DeptId:        user.DeptId,
		Remark:        user.Remark,
		IsAdmin:       user.IsAdmin == 1,
		Address:       user.Address,
		Describe:      user.Describe,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		DeletedAt:     user.DeletedAt,
		OpenId:        user.OpenId,
		UserType:      UserType(user.UserType),
	}
}
