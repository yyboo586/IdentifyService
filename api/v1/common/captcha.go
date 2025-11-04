/*
* @desc:验证码参数
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/2 17:47
 */

package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `path:"/get" tags:"通用接口/验证码" method:"get" summary:"获取验证码"`
}

type CaptchaV2Req struct {
	g.Meta `path:"/v2" tags:"通用接口/验证码" method:"get" summary:"获取v2验证码"`
}

type CheckCaptchaV2Req struct {
	g.Meta `path:"/v2Check" tags:"通用接口/验证码" method:"post" summary:"检查v2验证码"`
	Key    string `json:"key"`
	Dots   string `json:"dots"`
}

type CaptchaRes struct {
	g.Meta       `mime:"application/json"`
	Key          string `json:"key"`
	Img          string `json:"img"`
	VerifyStatus int    `json:"verifyStatus"`
}

type CaptchaV2Res struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
	Thumb  string `json:"thumb"`
}

type CheckCaptchaV2Res struct {
	g.Meta `mime:"application/json"`
}

type SendSmsCodeReq struct {
	g.Meta       `path:"/send_sms_code" tags:"通用接口/验证码" method:"post" summary:"发送短信验证码"`
	Phone        string `json:"phone" v:"required#手机号不能为空"`
	BusinessType string `json:"business_type" v:"required#业务类型不能为空" dc:"业务类型(验证码登录)"`
}

type SendSmsCodeRes struct {
	g.Meta `mime:"application/json"`
}

type ValidateSMSCodeReq struct {
	g.Meta       `path:"/validate_sms_code" tags:"通用接口/验证码" method:"post" summary:"验证短信验证码"`
	Phone        string `json:"phone" v:"required#手机号不能为空" dc:"手机号"`
	BusinessType string `json:"business_type" v:"required#业务类型不能为空" dc:"业务类型(验证码登录)"`
	Code         string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
}

type ValidateSMSCodeRes struct {
	g.Meta `mime:"application/json"`
}
