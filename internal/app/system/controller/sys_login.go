package controller

import (
	"context"
	"strings"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"

	commonModel "IdentifyService/internal/app/common/model"
	commonService "IdentifyService/internal/app/common/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
	"github.com/yyboo586/common/LogModule"
	"github.com/yyboo586/common/authUtils/tokenUtils"
)

var (
	Login = loginController{}
)

type loginController struct {
	BaseController
}

func (c *loginController) Login(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		user        *model.LoginUserRes
		tokenPair   *tokenUtils.TokenPair
		permissions []string
		menuList    []*model.UserMenus
	)

	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.SysUser().ValidateByUserNameAndPassword(ctx, req)
	if err != nil {
		// 保存登录失败的日志信息
		service.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
			Status:    0,
			Username:  req.UserName,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "系统后台",
		})
		return
	}
	err = service.SysUser().UpdateLoginInfo(ctx, user.Id, ip)
	if err != nil {
		return
	}
	// 报存登录成功的日志信息
	service.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
		Status:    1,
		Username:  req.UserName,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "系统后台",
	})
	tokenData := map[string]interface{}{
		"user_id":       user.Id,
		"iuqt_id":       user.IUQTID,
		"user_name":     user.UserName,
		"user_nickname": user.UserNickname,
		"mobile":        user.Mobile,
		"device_id":     "",
		"user_type":     user.UserType,
	}
	tokenPair, err = service.Token().Generate(ctx, tokenData)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	//获取用户菜单数据
	menuList, permissions, err = service.SysUser().GetAdminRules(ctx, user.Id)
	if err != nil {
		return
	}
	res = &system.UserLoginRes{
		UserInfo:    user,
		Token:       tokenPair.AccessToken,
		MenuList:    menuList,
		Permissions: permissions,
	}
	//用户在线状态保存
	service.SysUserOnline().Invoke(gctx.New(), &model.SysUserOnlineParams{
		UserAgent: userAgent,
		Uuid:      gmd5.MustEncrypt(tokenPair.AccessToken),
		Token:     "",
		Username:  user.UserName,
		Ip:        ip,
	})
	return
}

// LoginOut 退出登录
func (c *loginController) LoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {
	//_ = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}

func (c *loginController) Login2(ctx context.Context, req *system.UserLogin2Req) (res *system.UserLogin2Res, err error) {
	var (
		userInfo   *model.LoginUserRes
		tokenPair  *tokenUtils.TokenPair
		settleInfo *model.SettleInfo
	)
	userInfo, err = service.SysUser().Login2(ctx, req)
	if err != nil {
		return
	}

	settleInfo, err = service.ThirdService().GetSettleInfo(ctx, userInfo.Id, userInfo.UserType)
	if err != nil {
		return
	}

	// ip := libUtils.GetClientIp(ctx)
	// userAgent := libUtils.GetUserAgent(ctx)
	tokenData := map[string]interface{}{
		"user_id":       userInfo.Id,
		"iuqt_id":       userInfo.IUQTID,
		"user_name":     userInfo.UserName,
		"user_nickname": userInfo.UserNickname,
		"mobile":        userInfo.Mobile,
		"device_id":     "",
		"user_type":     userInfo.UserType,
	}
	tokenPair, err = service.Token().Generate(ctx, tokenData)
	if err != nil {
		return nil, gerror.Wrap(err, "登录失败，后端服务出现错误")
	}

	res = &system.UserLogin2Res{
		UserInfo: &system.UserInfo2{
			UserID:   userInfo.Id,
			IUQTID:   userInfo.IUQTID,
			UserName: userInfo.UserName,
			Mobile:   userInfo.Mobile,
			UserType: model.GetUserTypeText(userInfo.UserType),
		},
		SettleInfo: settleInfo,
		Token:      tokenPair.AccessToken,
	}

	return res, nil
}

func (c *loginController) TokenInspect(ctx context.Context, req *system.TokenIntrospectReq) (res *system.TokenIntrospectRes, err error) {
	// 初始化登录用户信息
	customClaims, IsActive, err := service.Token().Parse(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return nil, gerror.Wrap(err, "解析令牌失败")
	}
	if !IsActive {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code": 401,
			"msg":  "令牌已失效",
		})
		return nil, nil
	}

	g.Log().Info(ctx, customClaims)
	userID, _ := customClaims.Data.(map[string]interface{})["user_id"].(string)
	iuqtID, _ := customClaims.Data.(map[string]interface{})["iuqt_id"].(string)
	userName, _ := customClaims.Data.(map[string]interface{})["user_name"].(string)
	mobile, _ := customClaims.Data.(map[string]interface{})["mobile"].(string)
	deptID, _ := customClaims.Data.(map[string]interface{})["dept_id"].(float64)
	userType, _ := customClaims.Data.(map[string]interface{})["user_type"].(float64)

	res = &system.TokenIntrospectRes{}
	res.UserInfo2 = system.UserInfo2{
		UserID:   userID,
		IUQTID:   iuqtID,
		UserName: userName,
		Mobile:   mobile,
		DeptID:   uint64(deptID),
		UserType: model.GetUserTypeText(model.UserType(int(userType))),
	}
	return res, nil
}

func (c *loginController) RevokeToken(ctx context.Context, req *system.RevokeTokenReq) (res *system.RevokeTokenRes, err error) {
	err = service.Token().RevokeToken(ctx, service.Token().GetTokenFromRequest(ghttp.RequestFromCtx(ctx)))
	if err != nil {
		return nil, err
	}
	return &system.RevokeTokenRes{}, nil
}

func (c *loginController) LoginOrRegister(ctx context.Context, req *system.UserLogin3Req) (res *system.UserLogin3Res, err error) {
	var (
		exists    bool
		userInfo  *model.User
		tokenPair *tokenUtils.TokenPair
	)

	if req.Aggrement == nil {
		return nil, gerror.New("用户协议不能为空")
	}
	if !req.Aggrement.Accepted {
		return nil, gerror.New("用户协议未接受, 请先接受用户协议")
	}
	if strings.TrimSpace(req.Aggrement.Version) == "" {
		return nil, gerror.New("协议版本不能为空")
	}

	switch req.BussinessType {
	case "验证码登录":
		err = commonService.Captcha().ValidateSmsCode(ctx, req.Phone, commonModel.SMSBusinessTypeAccountLogin, req.Code)
		if err != nil {
			return nil, err
		}

		userInfo, exists, err = service.SysUser().GetUserInfoByPhone(ctx, req.Phone)
		if err != nil {
			return nil, err
		}
		if !exists {
			var userID string = uuid.New().String()
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				deptID, err := service.SysDept().CreateDept(ctx, tx, "部门"+req.Phone, userID)
				if err != nil {
					return err
				}
				err = service.SysUser().SelfRegister(ctx, tx, userID, deptID, req.Phone)
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return nil, err
			}

			userInfo = &model.User{
				ID:           userID,
				UserName:     "用户" + req.Phone,
				UserNickname: "用户" + req.Phone,
				Mobile:       req.Phone,
			}
		}
	case "密码登录":
		userInfo, exists, err = service.SysUser().GetUserInfoByPhone(ctx, req.Phone)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, gerror.New("用户不存在")
		}
		if userInfo.UserStatus != model.UserStatusNormal {
			return nil, gerror.New("账户已被冻结")
		}
		err = service.SysUser().ValidatePassword(ctx, userInfo, req.Code)
		if err != nil {
			return nil, err
		}
	default:
		return nil, gerror.New("不支持的业务类型")
	}

	userAgreement := &model.UserAgreement{
		UserID:        userInfo.ID,
		AgreementID:   req.Aggrement.ID,
		AgreementName: req.Aggrement.Name,
		Agreed:        req.Aggrement.Accepted,
	}
	if err := service.Agreement().RecordUserAgreements(ctx, userAgreement); err != nil {
		return nil, err
	}

	tokenData := map[string]interface{}{
		"user_id":       userInfo.ID,
		"iuqt_id":       userInfo.IUQTID,
		"user_name":     userInfo.UserName,
		"user_nickname": userInfo.UserNickname,
		"mobile":        userInfo.Mobile,
		"device_id":     req.DeviceInfo.DeviceID,
		"user_type":     userInfo.UserType,
		"dept_id":       userInfo.DeptId,
	}
	tokenPair, err = service.Token().Generate(ctx, tokenData)
	if err != nil {
		return nil, gerror.Wrap(err, "登录失败，后端服务出现错误")
	}
	g.Log().Info(ctx, "tokenPair.AccessToken", tokenPair.AccessToken)
	g.Log().Info(ctx, "tokenPair.RefreshToken", tokenPair.RefreshToken)

	if err := c.recordUserDevice(ctx, req, userInfo); err != nil {
		return nil, err
	}

	commonService.Log().WriteLog(ctx, []*LogModule.LogItem{
		{
			Module:     model.LogModuleUser,
			Action:     model.LogActionUserLogin,
			OperatorID: userInfo.ID,
			Message:    "用户登录",
			Detail:     nil,
			IP:         libUtils.GetClientIp(ctx),
		},
	})

	return &system.UserLogin3Res{
		Token:        tokenPair.AccessToken,
		UserID:       userInfo.ID,
		UserNickname: userInfo.UserNickname,
		Avatar:       userInfo.Avatar,
		Phone:        userInfo.Mobile,
		Sex:          userInfo.Sex,
		Birthday:     userInfo.Birthday,
		City:         userInfo.City,
	}, nil
}

func (c *loginController) recordUserDevice(ctx context.Context, req *system.UserLogin3Req, userInfo *model.User) error {
	if req.DeviceInfo == nil {
		return nil
	}

	deviceType := strings.ToLower(req.DeviceInfo.DeviceType)
	if deviceType != "android" && deviceType != "ios" {
		return nil
	}

	loginType := c.getLoginType(req.BussinessType)
	record := &model.UserDeviceRecordInput{
		UserID:     userInfo.ID,
		DeviceID:   req.DeviceInfo.DeviceID,
		DeviceName: req.DeviceInfo.DeviceName,
		DeviceIP:   libUtils.GetClientIp(ctx),
		LoginType:  loginType,
	}

	return service.UserDevice().RecordDevice(ctx, record)
}

func (c *loginController) getLoginType(businessType string) string {
	switch businessType {
	case "验证码登录":
		return "手机验证码"
	case "密码登录":
		return "手机密码"
	default:
		return businessType
	}
}
