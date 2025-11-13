package personal

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/do"
	"IdentifyService/internal/app/system/model/entity"
	service "IdentifyService/internal/app/system/service"
	"IdentifyService/library/libUtils"
	"IdentifyService/library/liberr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/yyboo586/common/MiddleWare"
)

func init() {
	service.RegisterPersonal(New())
}

type sPersonal struct {
}

func New() service.IPersonal {
	return &sPersonal{}
}

func (s *sPersonal) GetPersonalInfo(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error) {
	res = new(system.PersonalInfoRes)
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	userId := operator.UserID
	res.User, err = service.SysUser().GetUserInfoById(ctx, userId)
	var dept *entity.SysDept
	dept, err = service.SysDept().GetByDeptId(ctx, res.User.DeptId)
	res.DeptName = dept.DeptName
	allRoles, err := service.SysRole().GetRoleList(ctx)
	roles, err := service.SysUser().GetAdminRole(ctx, userId, allRoles)
	name := make([]string, len(roles))
	roleIds := make([]uint, len(roles))
	for k, v := range roles {
		name[k] = v.Name
		roleIds[k] = v.Id
	}
	res.Roles = name
	if err != nil {
		return
	}
	return
}

func (s *sPersonal) EditPersonal(ctx context.Context, req *system.PersonalEditReq) (user *model.LoginUserRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	userId := operator.UserID
	err = service.SysUser().CheckUserNameOrMobileExists(ctx, "", req.Mobile, userId)
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).WherePri(userId).Update(do.SysUser{
				Mobile:       req.Mobile,
				UserNickname: req.Nickname,
				Remark:       req.Remark,
				Sex:          req.Sex,
				UserEmail:    req.UserEmail,
				Describe:     req.Describe,
				Avatar:       req.Avatar,
			})
			liberr.ErrIsNil(ctx, err, "修改用户信息失败")
			user, err = service.SysUser().GetUserById(ctx, userId)
			liberr.ErrIsNil(ctx, err)
		})
		return err
	})
	return
}

func (s *sPersonal) ResetPwdPersonal(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	userId := operator.UserID
	salt := grand.S(10)
	password := libUtils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Update(g.Map{
			dao.SysUser.Columns().UserSalt:     salt,
			dao.SysUser.Columns().UserPassword: password,
		})
		liberr.ErrIsNil(ctx, err, "重置用户密码失败")
	})
	return
}
