package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

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
		userInfo *model.User
		token    string
		menuList []*model.AuthRuleNode
		roleList []map[int64]string = make([]map[int64]string, 0)
	)

	defer func(ctx context.Context) {
		loginLog := &model.LoginLog{
			OrgID:     "",
			LoginName: req.Username,
			IP:        libUtils.GetClientIp(ctx),
			Browser:   libUtils.GetUserAgent(ctx),
			Success:   true,
			Message:   "登录成功",
			LoginTime: gtime.Now(),
		}
		if err != nil {
			if !strings.Contains(err.Error(), "用户不存在") {
				loginLog.OrgID = userInfo.OrgID
			}
			loginLog.Success = false
			loginLog.Message = "登录失败:" + err.Error()
			g.Log().Error(ctx, err)
		} else {
			loginLog.OrgID = userInfo.OrgID
		}
		service.Log().InvokeLoginLog(loginLog)
	}(ctx)

	userInfo, err = service.User().GetByUsername(ctx, req.Username)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if userInfo.Status != model.UserStatusEnabled {
		err = fmt.Errorf("账号已被冻结")
		return
	}

	orgInfo, err := service.Org().Get(ctx, userInfo.OrgID)
	if err != nil {
		return
	}
	if !orgInfo.Enabled {
		err = fmt.Errorf("组织已禁用")
		return
	}

	err = service.User().ValidateUsernameAndPassword(ctx, userInfo.Password, userInfo.Salt, req.Password)
	if err != nil {
		return
	}

	roles, err := service.Role().GetRolesByUserID(ctx, userInfo.ID)
	if err != nil {
		return
	}

	for _, v := range roles {
		item := map[int64]string{v.ID: v.Name}
		roleList = append(roleList, item)
	}
	userInfo.Roles = roleList
	g.Log().Info(ctx, userInfo.Roles)

	key := gconv.String(userInfo.ID) + "-" + gmd5.MustEncryptString(userInfo.Name) + gmd5.MustEncryptString(userInfo.Password)
	token, err = service.TokenService().Generate(ctx, key, userInfo)
	if err != nil {
		return
	}

	menuList, err = service.AuthRule().GetMenuTreeByUserID(ctx, userInfo.ID)
	if err != nil {
		return
	}

	buttons, err := service.AuthRule().GetButtonListByUserID(ctx, userInfo.ID)
	if err != nil {
		return
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
		ButtonList: buttons,
	}
	return
}

func (c *authController) LoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {

	return
}

func (c *authController) Introspect(ctx context.Context, req *system.IntrospectReq) (res *model.IntrospectRes, err error) {
	res, err = service.TokenService().Introspect(ctx, req.Authorization)
	if err != nil {
		if errors.Is(err, service.ErrTokenInvalid) || errors.Is(err, service.ErrTokenExpired) {
			g.RequestFromCtx(ctx).Response.Status = http.StatusOK
			g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		return
	}
	return
}
