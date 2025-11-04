/*
* @desc:验证码获取
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/2 17:45
 */

package controller

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
)

var Captcha = captchaController{}

type captchaController struct {
}

// Get 获取验证码
func (c *captchaController) Get(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	idKeyC, base64stringC, err := service.Captcha().GetVerifyImgString(ctx)
	res = &common.CaptchaRes{
		Key:          idKeyC,
		Img:          base64stringC,
		VerifyStatus: g.Cfg().MustGet(ctx, "system.verifyStatus").Int(),
	}
	return
}

// V2 验证码
func (c *captchaController) V2(ctx context.Context, req *common.CaptchaV2Req) (res *common.CaptchaV2Res, err error) {
	dots, img, thumb, key, err := service.Captcha().GetCaptchaV2(ctx)
	// 写入缓存
	service.Cache().Set(ctx, "captchaV2_"+key, dots, 10*60*time.Second)
	res = &common.CaptchaV2Res{
		Key:   key,
		Img:   img,
		Thumb: thumb,
	}
	return
}

func (c *captchaController) V2Check(ctx context.Context, req *common.CheckCaptchaV2Req) (res *common.CheckCaptchaV2Res, err error) {
	if req.Key == "" || req.Dots == "" {
		return nil, errors.New("验证码无效")
	}
	err = service.Captcha().CheckCaptchaV2(ctx, req.Key, req.Dots)
	return
}

func (c *captchaController) SendSmsCode(ctx context.Context, req *common.SendSmsCodeReq) (res *common.SendSmsCodeRes, err error) {
	// 发送短信
	err = service.Captcha().SendSmsCode(ctx, req.Phone, model.GetSMSBusinessType(req.BusinessType))

	return
}

func (c *captchaController) ValidateSMSCode(ctx context.Context, req *common.ValidateSMSCodeReq) (res *common.ValidateSMSCodeRes, err error) {
	err = service.Captcha().ValidateSmsCode(ctx, req.Phone, model.GetSMSBusinessType(req.BusinessType), req.Code)
	return
}
