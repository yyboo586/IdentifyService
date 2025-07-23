package logics

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"

	"IdentifyService/api/v1/system"

	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"IdentifyService/library/liberr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

var (
	userOnce     sync.Once
	userInstance *user
)

type user struct {
	enforcer         *casbin.SyncedEnforcer
	casBinUserPrefix string // CasBin 用户id前缀
	defaultPassword  string // 默认密码
}

func NewUser() service.IUser {
	userOnce.Do(func() {
		enforcer, err := service.CasbinEnforcer(context.Background())
		if err != nil {
			panic(err)
		}
		userInstance = &user{
			enforcer:         enforcer,
			casBinUserPrefix: "u_",
			defaultPassword:  "123456",
		}
	})
	return userInstance
}

/*
1、创建角色
2、创建用户，并绑定到角色
3、修改角色的权限
4、用户登录，似乎没有新增的权限。
*/
func (u *user) Create(ctx context.Context, req *system.UserCreateReq) (userID string, err error) {
	operatorInfo := service.ContextService().Get(ctx)
	if req.OrgID == "" {
		req.OrgID = operatorInfo.User.OrgID // 默认使用当前登录用户的组织ID
	}
	userID = uuid.New().String()
	salt := grand.S(10)
	userInsertData := map[string]interface{}{
		dao.User.Columns().ID:       userID,
		dao.User.Columns().Name:     req.UserName,
		dao.User.Columns().Nickname: req.NickName,
		dao.User.Columns().Password: libUtils.EncryptPassword(req.Password, salt),
		dao.User.Columns().Salt:     salt,
		dao.User.Columns().Status:   model.UserStatusEnabled,
		dao.User.Columns().OrgID:    req.OrgID,
		dao.User.Columns().Sex:      req.Sex,
		dao.User.Columns().Email:    req.Email,
		dao.User.Columns().Avatar:   req.Avatar,
		dao.User.Columns().Mobile:   req.Mobile,
		dao.User.Columns().Address:  req.Address,
		dao.User.Columns().Describe: req.Describe,
		dao.User.Columns().IsAdmin:  req.IsAdmin,
	}
	if len(req.RoleIDs) > 0 {
		req.RoleIDs, err = roleInstance.FilterRoleIDs(ctx, req.RoleIDs, operatorInfo.User.ID, false)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
	g.Log().Info(ctx, "[DEBUG] roleIDs: ", "roleIDs", req.RoleIDs)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.User.Ctx(ctx).TX(tx).Insert(userInsertData)
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = fmt.Errorf("添加用户失败，用户名已存在")
				return
			}
			return
		}
		if len(req.RoleIDs) > 0 {
			err = u.assignUserRoles(ctx, req.RoleIDs, userID)
			if err != nil {
				err = fmt.Errorf("设置用户权限失败: %w", err)
				return
			}
		}
		return
	})
	return
}

// 前台用户自注册
// 1、应该创建一个组织
// 2、应该为当前用户绑定 前台组织管理员 的角色
func (u *user) Register(ctx context.Context, req *system.UserRegisterReq) (userID string, err error) {
	defer func() {
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}()

	salt := grand.S(10)
	userID = uuid.New().String()
	req.Password = libUtils.EncryptPassword(req.Password, salt)
	userInsertData := map[string]interface{}{
		dao.User.Columns().ID:       userID,
		dao.User.Columns().Name:     req.UserName,
		dao.User.Columns().Nickname: req.UserName,
		dao.User.Columns().Password: req.Password,
		dao.User.Columns().Salt:     salt,
		dao.User.Columns().Status:   model.UserStatusEnabled,
	}
	orgInsetData := &model.Org{
		PID:         consts.DefaultFrontOrgID,
		Name:        fmt.Sprintf("Org-%v", userID),
		ManagerID:   userID,
		ManagerName: req.UserName,
	}
	roleIDs := []int64{2}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var orgID string
		orgID, err = orgInstance.Add(ctx, orgInsetData)
		if err != nil {
			err = fmt.Errorf("添加组织失败: %w", err)
			return
		}

		userInsertData[dao.User.Columns().OrgID] = orgID
		_, err = dao.User.Ctx(ctx).TX(tx).Insert(userInsertData)
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = fmt.Errorf("添加用户失败，用户名已存在")
				return err
			}
			return err
		}

		err = u.assignUserRoles(ctx, roleIDs, userID)
		if err != nil {
			err = fmt.Errorf("设置用户权限失败: %w", err)
			return
		}
		return
	})

	return
}

func (u *user) Delete(ctx context.Context, ids []string) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.User.Ctx(ctx).TX(tx).Where(dao.User.Columns().ID+" in(?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除用户失败")
			// 删除对应权限
			for _, v := range ids {
				_, err = u.enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%s", u.casBinUserPrefix, v))
				liberr.ErrIsNil(ctx, err)
			}
		})
		return err
	})
	return
}

func (u *user) EditPersonalInfo(ctx context.Context, req *system.EditUserPersonalInfoReq) (err error) {
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().ID, req.ID).Update(map[string]interface{}{
		dao.User.Columns().Nickname: req.NickName,
		dao.User.Columns().Email:    req.Email,
		dao.User.Columns().Avatar:   req.Avatar,
		dao.User.Columns().Mobile:   req.Mobile,
		dao.User.Columns().Address:  req.Address,
		dao.User.Columns().Describe: req.Describe,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (u *user) EditUserPermission(ctx context.Context, req *system.EditUserPermissionReq) (err error) {
	operatorInfo := service.ContextService().Get(ctx)
	dataUpdate := map[string]interface{}{
		dao.User.Columns().Status:  model.UserStatusEnabled,
		dao.User.Columns().IsAdmin: req.IsAdmin,
	}
	if !req.Enabled {
		dataUpdate[dao.User.Columns().Status] = model.UserStatusDisable
	}

	if len(req.RoleIDs) > 0 {
		req.RoleIDs, err = roleInstance.FilterRoleIDs(ctx, req.RoleIDs, operatorInfo.User.ID, false)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.User.Ctx(ctx).TX(tx).Where(dao.User.Columns().ID, req.ID).Update(dataUpdate)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}

		if len(req.RoleIDs) > 0 {
			err = u.removeUserRoles(ctx, req.ID)
			if err != nil {
				err = fmt.Errorf("删除用户旧角色信息失败: %w", err)
				return
			}

			err = u.assignUserRoles(ctx, req.RoleIDs, req.ID)
			if err != nil {
				err = fmt.Errorf("设置用户权限失败: %w", err)
				return
			}
		}

		return
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (u *user) ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error) {
	salt := grand.S(10)
	password := libUtils.EncryptPassword(u.defaultPassword, salt)

	_, err = dao.User.Ctx(ctx).WherePri(req.ID).Update(g.Map{
		dao.User.Columns().Salt:     salt,
		dao.User.Columns().Password: password,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (u *user) IsSuperAdmin(ctx context.Context, userID string) bool {
	return userID == consts.DefaultSuperAdminID
}

func (u *user) ValidateUsernameAndPassword(ctx context.Context, userName, password string) (out *model.User, err error) {
	out, err = u.getUserInfoByUsername(ctx, userName)
	if err != nil {
		err = fmt.Errorf("账号/密码错误")
		return
	}
	if libUtils.EncryptPassword(password, out.Salt) != out.Password {
		err = fmt.Errorf("账号/密码错误")
		g.Log().Error(ctx, err)
		return
	}
	if out.Status != model.UserStatusEnabled {
		err = fmt.Errorf("账号已被冻结")
		return
	}

	return
}

func (u *user) GetUserInfoByID(ctx context.Context, id string) (out *model.User, err error) {
	var userEntity entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().ID, id).Scan(&userEntity)
	if err != nil {
		return
	}
	out = u.convertEntityToModel(&userEntity)
	return
}

func (u *user) List(ctx context.Context, req *system.UserListReq) (total interface{}, out []*model.User, err error) {
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}

	m := dao.User.Ctx(ctx).Where(dao.User.Columns().OrgID, req.OrgID)
	if req.Name != "" {
		m = m.Where(dao.User.Columns().Name, req.Name)
	}

	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	var userEntityList []*entity.User
	err = m.Page(req.PageNum, req.PageSize).Scan(&userEntityList)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range userEntityList {
		out = append(out, u.convertEntityToModel(v))
	}
	return
}

// 私有方法
// removeUserRoles 删除用户角色
func (u *user) removeUserRoles(ctx context.Context, userID string) (err error) {
	_, err = u.enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, fmt.Sprintf("%s%s", u.casBinUserPrefix, userID))
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

// assignUserRoles 将角色分配给用户
func (u *user) assignUserRoles(ctx context.Context, roleIDs []int64, userID string) (err error) {
	for _, v := range roleIDs {
		_, err = u.enforcer.AddNamedGroupingPolicy("g", fmt.Sprintf("%s%s", u.casBinUserPrefix, userID), gconv.String(v))
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}

	return
}

// getUserInfoByUsername 通过用户名获取用户信息
func (u *user) getUserInfoByUsername(ctx context.Context, userName string) (user *model.User, err error) {
	var userEntity entity.User
	err = dao.User.Ctx(ctx).Fields(userEntity).Where(dao.User.Columns().Name, userName).Scan(&userEntity)
	if err != nil {
		return
	}

	user = u.convertEntityToModel(&userEntity)
	return
}

func (u *user) convertEntityToModel(in *entity.User) (out *model.User) {
	out = &model.User{
		ID:        in.ID,
		Name:      in.Name,
		Nickname:  in.Nickname,
		Password:  in.Password,
		Salt:      in.Salt,
		Status:    model.UserStatus(in.Status),
		OrgID:     in.OrgID,
		Sex:       model.UserSex(in.Sex),
		Email:     in.Email,
		Avatar:    in.Avatar,
		Mobile:    in.Mobile,
		Address:   in.Address,
		Describe:  in.Describe,
		IsAdmin:   in.IsAdmin == 1,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return
}
