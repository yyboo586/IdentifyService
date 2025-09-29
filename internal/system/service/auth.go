package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"IdentifyService/api/v1/system"

	commonModel "IdentifyService/internal/common/model"
	commonService "IdentifyService/internal/common/service"
	"IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type authService struct {
	roleDomain       interfaces.IRole
	userDomain       interfaces.IUser
	orgDomain        interfaces.IOrg
	permissionDomain interfaces.IPermission
}

func NewAuthService(roleDomain interfaces.IRole, userDomain interfaces.IUser, orgDomain interfaces.IOrg, permissionDomain interfaces.IPermission) *authService {
	return &authService{
		roleDomain:       roleDomain,
		userDomain:       userDomain,
		orgDomain:        orgDomain,
		permissionDomain: permissionDomain,
	}
}

// TODO: CASBIN的操作，如何和事务关联起来？
func (c *authService) Registration(ctx context.Context, req *system.UserRegistrationReq) (res *system.UserRegistrationRes, err error) {
	res = &system.UserRegistrationRes{}

	userID := uuid.New().String()
	orgID := uuid.New().String()
	userInfo := &model.User{
		ID:       userID,
		Name:     req.Username,
		Nickname: req.Username,
		Password: req.Password,
		OrgID:    orgID,
	}
	orgInfo := &model.Org{
		ID:          orgID,
		PID:         "",
		Name:        fmt.Sprintf("Org-%v", userID),
		ManagerID:   userID,
		ManagerName: req.Username,
	}

	defer func(ctx context.Context) {
		contextUser := &commonModel.ContextUser{
			UserID:       "",
			UserName:     req.Username,
			UserNickname: req.Username,
			OrgID:        "",
			RoleIDs:      []int64{},
		}
		if err == nil {
			contextUser.UserID = userInfo.ID
			contextUser.OrgID = userInfo.OrgID
		}
		commonService.ContextService().SetUserInfo(g.RequestFromCtx(ctx), contextUser)
	}(ctx)

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if err = c.userDomain.Create(ctx, tx, userInfo); err != nil {
			return
		}

		if err = c.orgDomain.Create(ctx, tx, orgInfo); err != nil {
			return
		}

		defaultOrgAdminRole, err := c.roleDomain.GetByName(ctx, model.DefaultRoleOrgAdmin)
		if err != nil {
			return
		}

		if err = c.permissionDomain.AssignUserRoles(ctx, userID, []int64{defaultOrgAdminRole.ID}); err != nil {
			return
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return
}

func (c *authService) Login(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		userInfo             *model.User
		token                string
		permissionPointCodes []model.PermissionPointCode
	)

	// 补充上下文信息，方便记录日志
	defer func(ctx context.Context) {
		contextUser := &commonModel.ContextUser{
			UserID:       "",
			UserName:     req.Username,
			UserNickname: req.Username,
			OrgID:        "",
			RoleIDs:      []int64{},
		}
		if userInfo != nil {
			contextUser.UserID = userInfo.ID
			contextUser.OrgID = userInfo.OrgID
		}
		commonService.ContextService().SetUserInfo(g.RequestFromCtx(ctx), contextUser)
	}(ctx)

	// 获取用户信息
	userInfo, err = c.userDomain.GetByUsername(ctx, req.Username)
	if err != nil {
		return
	}

	// 判断账户所属组织是否被禁用
	_, err = c.orgDomain.IsEnabled(ctx, userInfo.OrgID)
	if err != nil {
		return
	}

	// 判断账户密码是否正确
	err = c.userDomain.ValidateUsernameAndPassword(ctx, userInfo.Password, userInfo.Salt, req.Password)
	if err != nil {
		return
	}

	// 根据用户ID获取角色ID列表
	roleIDs, err := c.permissionDomain.GetRoleIDsByUserID(ctx, userInfo.ID)
	if err != nil {
		return
	}

	data := map[string]interface{}{
		"user_id":       userInfo.ID,
		"user_name":     userInfo.Name,
		"user_nickname": userInfo.Nickname,
		"org_id":        userInfo.OrgID,
		"role_ids":      roleIDs,
	}
	token, err = commonService.TokenService().Generate(ctx, data)
	if err != nil {
		return
	}

	// 根据用户ID获取权限点代码列表
	permissionPointCodes, err = c.permissionDomain.GetPermissionPointCodesByUserID(ctx, userInfo.ID)
	if err != nil {
		return
	}

	res = &system.UserLoginRes{
		UserInfo: &model.UserLoginRes{
			ID:       userInfo.ID,
			Name:     userInfo.Name,
			Nickname: userInfo.Nickname,
			IsAdmin:  userInfo.IsAdmin,
			Avatar:   userInfo.Avatar,
			OrgID:    userInfo.OrgID,
			RoleIDs:  roleIDs,
		},
		Token:                token,
		PermissionPointCodes: permissionPointCodes,
	}
	return
}

func (c *authService) LoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {
	token := strings.TrimPrefix(req.Authorization, "Bearer ")

	customClaims, err := commonService.TokenService().Parse(ctx, token)
	if err != nil {
		return
	}

	err = commonService.TokenService().Revoke(ctx, token, customClaims.CustomData["user_id"].(string))
	if err != nil {
		return
	}

	varRoleIDs, ok := customClaims.CustomData["role_ids"].([]interface{})
	if !ok {
		return
	}
	roleIDs := make([]int64, 0, len(varRoleIDs))
	for _, roleID := range varRoleIDs {
		v, ok := roleID.(float64)
		if !ok {
			return
		}
		roleIDs = append(roleIDs, int64(v))
	}
	commonService.ContextService().Init(g.RequestFromCtx(ctx), &commonModel.ContextUser{
		UserID:       customClaims.CustomData["user_id"].(string),
		UserName:     customClaims.CustomData["user_name"].(string),
		UserNickname: customClaims.CustomData["user_nickname"].(string),
		OrgID:        customClaims.CustomData["org_id"].(string),
		RoleIDs:      roleIDs,
	})
	return
}

func (c *authService) Introspect(ctx context.Context, req *system.IntrospectReq) (res *model.IntrospectRes, err error) {
	token := strings.TrimPrefix(req.Authorization, "Bearer ")
	res, err = commonService.TokenService().Introspect(ctx, token)
	if err != nil {
		if errors.Is(err, commonService.ErrTokenExpired) {
			g.RequestFromCtx(ctx).Response.Status = http.StatusOK
			g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "token expired",
			})
			return
		} else if errors.Is(err, commonService.ErrTokenInvalid) {
			g.RequestFromCtx(ctx).Response.Status = http.StatusOK
			g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "token invalid",
			})
			return
		}
		return
	}
	return
}
