package domain

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"

	"IdentifyService/api/v1/system"

	"IdentifyService/internal/system/dao"
	service "IdentifyService/internal/system/interfaces"
	"IdentifyService/internal/system/model"
	"IdentifyService/internal/system/model/entity"
	"IdentifyService/library/libUtils"

	commonService "IdentifyService/internal/common/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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

func NewUser(enforcer *casbin.SyncedEnforcer) service.IUser {
	userOnce.Do(func() {
		userInstance = &user{
			enforcer:         enforcer,
			casBinUserPrefix: "u_",
			defaultPassword:  "123456",
		}
	})
	return userInstance
}

// checked
func (u *user) Create(ctx context.Context, tx gdb.TX, in *model.User) (err error) {
	salt := grand.S(10)
	isAdmin := 0
	if in.IsAdmin {
		isAdmin = 1
	}
	userInsertData := map[string]interface{}{
		dao.User.Columns().ID:       in.ID,
		dao.User.Columns().Name:     in.Name,
		dao.User.Columns().Nickname: in.Nickname,
		dao.User.Columns().Password: libUtils.EncryptPassword(in.Password, salt),
		dao.User.Columns().Salt:     salt,
		dao.User.Columns().Status:   model.UserStatusEnabled,
		dao.User.Columns().OrgID:    in.OrgID,
		dao.User.Columns().Sex:      in.Sex,
		dao.User.Columns().Email:    in.Email,
		dao.User.Columns().Avatar:   in.Avatar,
		dao.User.Columns().Mobile:   in.Mobile,
		dao.User.Columns().Address:  in.Address,
		dao.User.Columns().Describe: in.Describe,
		dao.User.Columns().IsAdmin:  isAdmin,
	}

	_, err = dao.User.Ctx(ctx).TX(tx).Insert(userInsertData)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fmt.Errorf("%w, %v", model.ErrRecordAlreadyExists, err)
		}
		return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	return nil
}

func (u *user) Delete(ctx context.Context, tx gdb.TX, ids []string) (err error) {
	if len(ids) == 0 {
		return nil
	}

	_, err = dao.User.Ctx(ctx).TX(tx).Where(dao.User.Columns().ID+" in(?)", ids).Delete()
	if err != nil {
		return fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	return nil
}

func (u *user) Edit(ctx context.Context, req *system.UserEditReq) (err error) {
	var data map[string]interface{} = make(map[string]interface{})

	if req.NickName != "" {
		data[dao.User.Columns().Nickname] = req.NickName
	}
	if req.Email != "" {
		data[dao.User.Columns().Email] = req.Email
	}
	if req.Avatar != "" {
		data[dao.User.Columns().Avatar] = req.Avatar
	}
	if req.Mobile != "" {
		data[dao.User.Columns().Mobile] = req.Mobile
	}
	if req.Address != "" {
		data[dao.User.Columns().Address] = req.Address
	}
	if req.Describe != "" {
		data[dao.User.Columns().Describe] = req.Describe
	}
	if model.IsValidUserStatus(model.UserStatus(req.Status)) {
		data[dao.User.Columns().Status] = req.Status
	}

	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().ID, req.UserID).Update(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

func (u *user) ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error) {
	salt := grand.S(10)
	password := libUtils.EncryptPassword(u.defaultPassword, salt)

	_, err = dao.User.Ctx(ctx).WherePri(req.UserID).Update(g.Map{
		dao.User.Columns().Salt:     salt,
		dao.User.Columns().Password: password,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}

// checked
func (u *user) ValidateUsernameAndPassword(ctx context.Context, hashPassword, salt, password string) (err error) {
	if libUtils.EncryptPassword(password, salt) != hashPassword {
		return fmt.Errorf("%w, %v", model.ErrBadRequest, fmt.Errorf("账户/密码错误"))
	}
	return nil
}

func (u *user) IsSuperAdmin(ctx context.Context, userID string) bool {
	return userID == model.DefaultSuperAdminID
}

func (u *user) IsEnabled(ctx context.Context, in *model.User) (err error) {
	if in.Status != model.UserStatusEnabled {
		err = fmt.Errorf("账户已被禁用")
	}
	return
}
func (u *user) GetByID(ctx context.Context, id string) (out *model.User, err error) {
	var userEntity entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().ID, id).Scan(&userEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("用户不存在")
		}
		g.Log().Error(ctx, err)
		return
	}
	out = u.convertEntityToModel(&userEntity)
	return
}

// checked
func (u *user) GetByUsername(ctx context.Context, username string) (out *model.User, err error) {
	var userEntity entity.User

	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Name, username).Scan(&userEntity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%w, %v", model.ErrRecordNotFound, fmt.Errorf("user_name: %s", username))
		}
		return nil, fmt.Errorf("%w, %v", model.ErrServerInternal, err)
	}

	return u.convertEntityToModel(&userEntity), nil
}

// List 根据组织ID获取用户列表
// 1、如果当前用户是超级管理员，则可以获取所有用户列表
// 2、如果当前用户不是超级管理员，则只能获取当前组织下的用户列表
func (u *user) Search(ctx context.Context, req *system.UserSearchReq) (total interface{}, out []*model.User, err error) {
	currentUserId, err := commonService.ContextService().GetUserID(ctx)
	if err != nil {
		return nil, nil, err
	}
	orgID, err := commonService.ContextService().GetOrgID(ctx)
	if err != nil {
		return nil, nil, err
	}
	if !u.IsSuperAdmin(ctx, currentUserId) && req.OrgID != orgID {
		err = fmt.Errorf("无数据权限访问")
		return
	}

	if req.PageSize == 0 {
		req.PageSize = model.DefaultPageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}

	m := dao.User.Ctx(ctx)
	if !u.IsSuperAdmin(ctx, currentUserId) {
		m = m.Where(dao.User.Columns().OrgID, req.OrgID)
	}

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

func (u *user) IsOrgHasUsers(ctx context.Context, orgID string) (hasUsers bool, err error) {
	hasUsers, err = dao.User.Ctx(ctx).Where(dao.User.Columns().OrgID, orgID).Exist()
	if err != nil {
		return false, err
	}
	return hasUsers, nil
}

// 私有方法

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
