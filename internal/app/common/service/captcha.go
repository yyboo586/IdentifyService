package service

import (
	"context"

	"IdentifyService/internal/app/common/model"

	"github.com/wenlng/go-captcha/captcha"
)

type ICaptcha interface {
	GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
	VerifyString(id, answer string) bool
	GetCaptchaV2(ctx context.Context) (dots map[int]captcha.CharDot, img, thumb, key string, err error)
	CheckCaptchaV2(ctx context.Context, key string, dots string, removeKey ...bool) (err error)

	SendSmsCode(ctx context.Context, phone string, bussinessType model.SMSBusinessType) (code string, err error)
	ValidateSmsCode(ctx context.Context, phone string, bussinessType model.SMSBusinessType, code string) (err error)
}

var localCaptcha ICaptcha

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
