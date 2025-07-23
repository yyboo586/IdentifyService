package controller

import (
	"context"

	"IdentifyService/api/v1/system"

	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	AuthController = authController{}
)

type authController struct {
}

// 1、验证账户密码是否正确
// 2、生成访问令牌
// 3、根据用户ID获取拥有的角色ID列表
// 4、根据角色ID列表获取菜单树形结构
// 5、返回登录结果
func (c *authController) Login(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		userInfo   *model.User
		token      string
		buttonList []*system.Button
		menuList   []*model.AuthRuleNode
	)

	defer func(ctx context.Context) {
		loginLog := &model.LoginLog{
			LoginName: req.Username,
			IP:        libUtils.GetClientIp(ctx),
			Browser:   libUtils.GetUserAgent(ctx),
			Success:   true,
			Message:   "登录成功",
			LoginTime: gtime.Now(),
		}
		if err != nil {
			loginLog.Success = false
			loginLog.Message = "登录失败:" + err.Error()
			g.Log().Error(ctx, err)
		}
		service.Log().InvokeLoginLog(loginLog)
	}(ctx)

	userInfo, err = service.User().ValidateUsernameAndPassword(ctx, req.Username, req.Password)
	if err != nil {
		return
	}

	key := gconv.String(userInfo.ID) + "-" + gmd5.MustEncryptString(userInfo.Name) + gmd5.MustEncryptString(userInfo.Password)
	token, err = service.TokenService().GenerateToken(ctx, key, userInfo)
	if err != nil {
		return
	}

	menuList, err = service.AuthRule().GetMenuTreesByUserID(ctx, userInfo.ID, false)
	if err != nil {
		return
	}

	buttons, err := service.AuthRule().GetButtonListByUserID(ctx, userInfo.ID)
	if err != nil {
		return
	}
	for _, v := range buttons {
		buttonList = append(buttonList, &system.Button{
			ID:   v.ID,
			Pid:  v.Pid,
			Name: v.Name,
			Type: int64(v.Type),
		})
	}

	res = &system.UserLoginRes{
		UserInfo: &model.UserLoginRes{
			ID:       userInfo.ID,
			Name:     userInfo.Name,
			Nickname: userInfo.Nickname,
			Mobile:   userInfo.Mobile,
			Status:   int64(userInfo.Status),
			IsAdmin:  userInfo.IsAdmin,
			Avatar:   userInfo.Avatar,
			OrgID:    userInfo.OrgID,
		},
		Token:      token,
		MenuList:   menuList,
		ButtonList: buttonList,
	}
	return
}

func (c *authController) LoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {

	return
}

func (c *authController) Introspect(ctx context.Context, req *system.IntrospectReq) (res *model.IntrospectRes, err error) {
	res, err = service.TokenService().Introspect(ctx, req.Token)
	if err != nil {
		return
	}
	return
}
