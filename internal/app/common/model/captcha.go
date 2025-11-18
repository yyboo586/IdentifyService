package model

import (
	"time"

	"IdentifyService/internal/app/common/model/entity"
)

type SMSBusinessType int

const (
	BusinessTypeUnknown          SMSBusinessType = iota
	SMSBusinessTypeAccountLogin                  // 验证码注册/登录
	SMSBusinessTypeBindPhone                     // 绑定/换绑手机号
	SMSBusinessTypeResetPassword                 // 修改密码
	SMSBusinessTypeBindIDCard                    // 绑定身份证号
	SMSBusinessTypeUnRegister                    // 注销账户
)

func GetSMSBusinessType(businessType string) SMSBusinessType {
	switch businessType {
	case "验证码注册/登录":
		return SMSBusinessTypeAccountLogin
	case "绑定/换绑手机号":
		return SMSBusinessTypeBindPhone
	case "修改密码":
		return SMSBusinessTypeResetPassword
	case "绑定身份证号":
		return SMSBusinessTypeBindIDCard
	case "注销账户":
		return SMSBusinessTypeUnRegister
	default:
		return BusinessTypeUnknown
	}
}

type SMSCodeStatus int

const (
	SMSCodeStatusInit SMSCodeStatus = iota // 未使用
	SMSCodeStatusUsed                      // 已使用
)

type SMS struct {
	Id           int64           `json:"id"`
	BusinessType SMSBusinessType `json:"business_type"`
	Phone        string          `json:"phone"`
	Code         string          `json:"code"`
	Status       SMSCodeStatus   `json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
	ExpiredAt    time.Time       `json:"expired_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func ConvertSmsEntity(in *entity.TSmsCode) (out *SMS) {
	out = &SMS{
		Id:           in.Id,
		BusinessType: SMSBusinessType(in.BusinessType),
		Phone:        in.Phone,
		Code:         in.Code,
		Status:       SMSCodeStatus(in.Status),
		CreatedAt:    time.Unix(in.CreatedAt, 0),
		ExpiredAt:    time.Unix(in.ExpiredAt, 0),
		UpdatedAt:    time.Unix(in.UpdatedAt, 0),
	}

	return out
}
