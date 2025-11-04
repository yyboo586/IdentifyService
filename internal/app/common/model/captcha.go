package model

import (
	"time"

	"github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)

type SMSBusinessType int

const (
	BusinessTypeUnknown  SMSBusinessType = iota
	SMSBusinessTypeLogin                 // 验证码登录
)

func GetSMSBusinessType(businessType string) SMSBusinessType {
	switch businessType {
	case "验证码登录":
		return SMSBusinessTypeLogin
	default:
		return BusinessTypeUnknown
	}
}

type SMSCodeStatus int

const (
	SMSCodeStatusUnknown SMSCodeStatus = iota // 未知状态
	SMSCodeStatusInit                         // 未使用
	SMSCodeStatusUsed                         // 已使用
)

func GetTSmsCodeStatus(status string) SMSCodeStatus {
	switch status {
	case "未使用":
		return SMSCodeStatusInit
	case "已使用":
		return SMSCodeStatusUsed
	default:
		return SMSCodeStatusUnknown
	}
}

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
