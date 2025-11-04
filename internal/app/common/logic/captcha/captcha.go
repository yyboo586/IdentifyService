/*
* @desc:验证码处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/28 9:01
 */

package captcha

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mojocn/base64Captcha"
	"github.com/tiger1103/gfast/v3/internal/app/common/dao"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
	captchaV2 "github.com/wenlng/go-captcha/captcha"
)

func init() {
	service.RegisterCaptcha(New())
}

func New() service.ICaptcha {
	return &sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
}

type sCaptcha struct {
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
}

// GetCaptchaV2 创建点击验证码数据
func (s *sCaptcha) GetCaptchaV2(ctx context.Context) (dots map[int]captchaV2.CharDot, img, thumb, key string, err error) {
	capt := captchaV2.GetCaptcha()
	dots, img, thumb, key, err = capt.Generate()
	return
}

// CheckCaptchaV2 验证captchaV2数据
func (s *sCaptcha) CheckCaptchaV2(ctx context.Context, key string, dots string, removeKey ...bool) (err error) {
	dotsStr, err := gbase64.DecodeToString(dots)
	if err != nil {
		return err
	}
	// 进行Url转码，防止特殊字符问题
	dotsStr, err = url.QueryUnescape(dotsStr)
	dotsMap := gconv.Maps(dotsStr)
	if dotsMap == nil {
		return errors.New("提交的数据无效")
	}

	cacheDots := service.Cache().Get(ctx, "captchaV2_"+key)
	if cacheDots == nil {
		return errors.New("未找到验证数据")
	}
	var dotsMap2 map[int]captchaV2.CharDot
	err = cacheDots.Scan(&dotsMap2)
	if len(dotsMap) != len(dotsMap2) {
		return errors.New("人机验证失败")
	}
	g.Log().Info(ctx, dotsMap, dotsMap2)
	for i, dot := range dotsMap {
		checkStatus := captchaV2.CheckPointDistWithPadding(gconv.Int64(dot["x"]), gconv.Int64(dot["y"]), int64(dotsMap2[i].Dx), int64(dotsMap2[i].Dy), int64(dotsMap2[i].Width), int64(dotsMap2[i].Height), 10)
		if checkStatus == false {
			return errors.New("人机验证失败")
		}
	}
	if len(removeKey) > 0 && removeKey[0] {
		service.Cache().Remove(ctx, "captchaV2_"+key)
	}
	return
}

var (
	captcha = sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
)

// GetVerifyImgString 获取字母数字混合验证码
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	driver := s.driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.store)
	idKeyC, base64stringC, _, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *sCaptcha) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(s.driver, s.store)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}

// 1、60秒内只能发送一次。如果60秒内发送过，则返回错误（"验证码已发送，请60秒后重试"）
// 2、生成验证码，并保存到数据库，过期时间5分钟
// 3、发送短信
func (s *sCaptcha) SendSmsCode(ctx context.Context, phone string, bussinessType model.SMSBusinessType) (err error) {
	// 检查60秒内是否发送过
	cacheKey := fmt.Sprintf("sms_code_%s_%d", phone, bussinessType)
	lastSendTime := service.Cache().Get(ctx, cacheKey)
	if lastSendTime != nil && !lastSendTime.IsEmpty() {
		return errors.New("验证码已发送, 请60秒后重试")
	}

	// 生成6位数字验证码
	code := grand.Digits(6)

	// 当前时间戳（秒）
	now := gtime.Now().Unix()

	// 设置过期时间为5分钟后
	expiredAt := now + 5*60

	// 保存到数据库
	smsCode := &entity.TSmsCode{
		BusinessType: int(bussinessType),
		Phone:        phone,
		Code:         code,
		Status:       int(model.SMSCodeStatusInit),
		CreatedAt:    now,
		ExpiredAt:    expiredAt,
		UpdatedAt:    now,
	}

	_, err = dao.TSmsCode.Ctx(ctx).Data(smsCode).Insert()
	if err != nil {
		return gerror.Newf(err.Error(), "保存验证码失败")
	}

	// 缓存发送时间，60秒过期
	service.Cache().Set(ctx, cacheKey, now, 60*time.Second)

	// TODO: 发送短信（这里可以调用实际的短信服务）
	g.Log().Infof(ctx, fmt.Sprintf("发送短信验证码: phone=%s, code=%s", phone, code))

	return nil
}

func (s *sCaptcha) ValidateSmsCode(ctx context.Context, phone string, bussinessType model.SMSBusinessType, code string) (err error) {
	// 从数据库获取
	var smsEntity entity.TSmsCode
	err = dao.TSmsCode.Ctx(ctx).
		Where(dao.TSmsCode.Columns().Phone, phone).
		Where(dao.TSmsCode.Columns().BusinessType, bussinessType).
		Order(dao.TSmsCode.Columns().CreatedAt, "DESC").
		Limit(1).
		Scan(&smsEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			return gerror.New("验证码错误")
		}
		return gerror.Newf(err.Error(), "获取验证码失败")
	}

	smsInfo := model.ConvertSmsEntity(&smsEntity)
	if smsInfo.Status != model.SMSCodeStatusInit {
		return gerror.New("验证码已被使用，请重新获取")
	}
	if smsInfo.ExpiredAt.Unix() < time.Now().Unix() {
		return gerror.New("验证码已过期，请重新获取")
	}
	if smsInfo.Code != code {
		return gerror.New("验证码错误")
	}

	dataUpdate := map[string]interface{}{
		dao.TSmsCode.Columns().Status:    model.SMSCodeStatusUsed,
		dao.TSmsCode.Columns().UpdatedAt: gtime.Now().Unix(),
	}
	_, err = dao.TSmsCode.Ctx(ctx).Where(dao.TSmsCode.Columns().Id, smsEntity.Id).Data(dataUpdate).Update()
	if err != nil {
		return gerror.Newf(err.Error(), "更新验证码状态失败")
	}

	return nil
}
