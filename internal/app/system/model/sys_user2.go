package model

import (
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/library/libSecurity"

	"github.com/gogf/gf/v2/os/gtime"
)

type UserSex int

const (
	UserSexSecret UserSex = iota
	UserSexMale
	UserSexFemale
)

func IsValidUserSex(sex UserSex) bool {
	return sex == UserSexSecret || sex == UserSexMale || sex == UserSexFemale
}

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

func IsValidCardType(cardType string) bool {
	switch cardType {
	case "居民身份证", "港澳居民往来内地通行证", "台湾居民往来大陆通行证":
		return true
	default:
		return false
	}
}

type User struct {
	ID           string     `json:"id"`
	DeptId       uint64     `json:"dept_id"`
	UserName     string     `json:"user_name"`
	UserNickname string     `json:"user_nickname"`
	UserPassword string     `json:"user_password"`
	UserSalt     string     `json:"user_salt"`
	UserStatus   UserStatus `json:"user_status"`
	IsAdmin      bool       `json:"is_admin"`

	Mobile    string `json:"mobile"`
	UserEmail string `json:"user_email"`
	Sex       int    `json:"sex"`
	Avatar    string `json:"avatar"`
	City      string `json:"city"`
	Birthday  string `json:"birthday"`
	Address   string `json:"address"`
	Describe  string `json:"describe"`
	Remark    string `json:"remark"`

	OpenId   string   `json:"openId"`
	IUQTID   string   `json:"iuqt_id"`
	UserType UserType `json:"user_type"`

	CardType string `json:"card_type"`
	IDCard   string `json:"id_card"`
	RealName string `json:"real_name"`

	LastLoginIp   string      `json:"last_login_ip"`
	LastLoginTime *gtime.Time `json:"last_login_time"`

	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
	DeletedAt *gtime.Time `json:"deleted_at"`
}

func ConvertToUser(in *entity.SysUser) (out *User) {
	idCard, _ := libSecurity.DecryptIDCard(in.IDCard)

	out = &User{
		ID:           in.Id,
		UserName:     in.UserName,
		UserNickname: in.UserNickname,
		UserPassword: in.UserPassword,
		UserSalt:     in.UserSalt,
		UserStatus:   UserStatus(in.UserStatus),
		IsAdmin:      in.IsAdmin == 1,

		Mobile:    in.Mobile,
		UserEmail: in.UserEmail,
		Sex:       in.Sex,
		Avatar:    in.Avatar,
		City:      in.City,
		Birthday:  in.Birthday,
		DeptId:    in.DeptId,
		Address:   in.Address,
		Describe:  in.Describe,
		Remark:    in.Remark,

		CardType: in.CardType,
		IDCard:   idCard,
		RealName: in.RealName,

		LastLoginIp:   in.LastLoginIp,
		LastLoginTime: in.LastLoginTime,

		OpenId:   in.OpenId,
		IUQTID:   in.IUQTID,
		UserType: UserType(in.UserType),

		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		DeletedAt: in.DeletedAt,
	}
	return out
}

type UserPersonalInfo struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Sex      int    `json:"sex"`
	Birthday string `json:"birthday"`
	City     string `json:"city"`
}

type UserIDCardInfo struct {
	CardType string `json:"card_type"`
	IDCard   string `json:"id_card"`
	RealName string `json:"real_name"`
}
