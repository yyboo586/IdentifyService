/*
* @desc:登录
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/27 21:52
 */

package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
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
		token       string
		permissions []string
		menuList    []*model.UserMenus
	)

	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.SysUser().GetAdminUserByUsernamePassword(ctx, req)
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
	token, err = service.Token().Generate(ctx, user)
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
		Token:       token,
		MenuList:    menuList,
		Permissions: permissions,
	}
	//用户在线状态保存
	service.SysUserOnline().Invoke(gctx.New(), &model.SysUserOnlineParams{
		UserAgent: userAgent,
		Uuid:      gmd5.MustEncrypt(token),
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
		token      string
		settleInfo *model.SettleInfo
	)
	userInfo, err = service.SysUser().Login2(ctx, req)
	if err != nil {
		return
	}

	settleInfo, err = service.ThirdService().GetSettleInfo(ctx, int64(userInfo.Id), userInfo.UserType)
	if err != nil {
		return
	}

	// ip := libUtils.GetClientIp(ctx)
	// userAgent := libUtils.GetUserAgent(ctx)
	token, err = service.Token().Generate(ctx, userInfo)
	if err != nil {
		return nil, gerror.Wrap(err, "登录失败，后端服务出现错误")
	}

	res = &system.UserLogin2Res{
		UserInfo: &system.UserInfo2{
			UserID:   fmt.Sprintf("%d", userInfo.Id),
			IUQTID:   userInfo.IUQTID,
			UserName: userInfo.UserName,
			Mobile:   userInfo.Mobile,
			UserType: model.GetUserTypeText(userInfo.UserType),
		},
		SettleInfo: settleInfo,
		Token:      token,
	}

	return res, nil
}

func (c *loginController) TokenInspect(ctx context.Context, req *system.TokenIntrospectReq) (res *system.TokenIntrospectRes, err error) {
	// 初始化登录用户信息
	data, err := service.Token().Parse(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return nil, gerror.Wrap(err, "解析令牌失败")
	}

	var userInfo model.LoginUserRes
	err = gconv.Struct(data.Data, &userInfo)
	if err != nil {
		return nil, gerror.Wrap(err, "解析令牌失败")
	}

	res = &system.TokenIntrospectRes{}
	res.UserInfo2 = system.UserInfo2{
		UserID:   fmt.Sprintf("%d", userInfo.Id),
		IUQTID:   userInfo.IUQTID,
		UserName: userInfo.UserName,
		Mobile:   userInfo.Mobile,
		UserType: model.GetUserTypeText(userInfo.UserType),
	}
	return res, nil
}

func (c *loginController) TokenRefresh(ctx context.Context, req *system.TokenRefreshReq) (res *system.TokenRefreshRes, err error) {
	oldToken := service.Token().GetTokenFromRequest(g.RequestFromCtx(ctx))
	if oldToken == "" {
		return nil, gerror.New("令牌不存在")
	}

	token, err := service.Token().Refresh(ctx, oldToken)
	if err != nil {
		return nil, gerror.Wrap(err, "刷新令牌失败")
	}

	res = &system.TokenRefreshRes{
		Token: token,
	}
	return res, nil
}
