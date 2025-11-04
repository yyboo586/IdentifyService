package sysUser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"

	commonModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
)

func (s *sSysUser) Login2(ctx context.Context, req *system.UserLogin2Req) (out *model.LoginUserRes, err error) {
	var userInfo *model.User
	switch model.LoginType(req.LoginType) {
	case model.LoginTypePhonePassword:
		userInfo, err = s.loginByPhonePassword(ctx, req.Phone, req.Code)
		if err != nil {
			return nil, err
		}
	case model.LoginTypePhoneSMSCode:
		userInfo, err = s.loginByPhoneSMSCode(ctx, req.Phone, req.Code)
		if err != nil {
			return nil, err
		}
	case model.LoginTypeSSO:
		userInfo, err = s.loginBySSO(ctx, req.Code)
		if err != nil {
			return nil, err
		}
	default:
		return nil, gerror.New("位置的登录类型")
	}

	// 账号状态
	if userInfo.UserStatus != model.UserStatusNormal {
		return nil, gerror.Newf("账户状态异常[%s]，请联系管理员", model.GetUserStatus(userInfo.UserStatus))
	}

	out = &model.LoginUserRes{
		Id:           uint64(userInfo.ID),
		UserName:     userInfo.UserName,
		Mobile:       userInfo.Mobile,
		UserStatus:   uint(userInfo.UserStatus),
		UserNickname: userInfo.UserNickname,
		UserPassword: userInfo.UserPassword,
		UserSalt:     userInfo.UserSalt,
		IUQTID:       userInfo.IUQTID,
		UserType:     userInfo.UserType,
	}
	return out, nil
}

// - 如果是未注册用户，提示其只能通过手机验证码登录。
// - 如果是已注册用户，但没设置密码，提示其只能通过手机验证码登录。
// - 如果是已注册用户，验证其手机号、密码是否匹配。
func (s *sSysUser) loginByPhonePassword(ctx context.Context, phone, password string) (userInfo *model.User, err error) {
	userInfo, err = s.GetUserByPhone2(ctx, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, gerror.New("该手机号未注册，请选择其他登录方式")
		}
		return nil, gerror.Wrap(err, "获取用户信息失败")
	}

	if userInfo.UserPassword == "" {
		return nil, gerror.New("该账户未设置密码, 无法通过密码登录")
	}

	// 验证密码
	if libUtils.EncryptPassword(password, userInfo.UserSalt) != userInfo.UserPassword {
		return nil, gerror.New("账户密码错误")
	}

	return userInfo, nil
}

// - 如果是未注册用户，自动注册。
// - 如果是已注册用户，验证手机验证码是否正确。
func (s *sSysUser) loginByPhoneSMSCode(ctx context.Context, phone, code string) (userInfo *model.User, err error) {
	err = commonService.Captcha().ValidateSmsCode(ctx, phone, commonModel.SMSBusinessTypeLogin, code)
	if err != nil {
		return nil, err
	}

	userInfo, err = s.GetUserByPhone2(ctx, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			data := map[string]interface{}{
				dao.SysUser.Columns().UserName:     fmt.Sprintf("user_%s", phone),
				dao.SysUser.Columns().Mobile:       phone,
				dao.SysUser.Columns().UserNickname: fmt.Sprintf("user_%s", phone),
				dao.SysUser.Columns().UserStatus:   model.UserStatusNormal,
				dao.SysUser.Columns().UserEmail:    "",
				dao.SysUser.Columns().Sex:          0,
			}
			userID, err := s.registerUser(ctx, data)
			if err != nil {
				return nil, err
			}

			userInfo = &model.User{
				ID:         userID,
				UserName:   fmt.Sprintf("user_%s", phone),
				UserStatus: model.UserStatusNormal,
			}
			return userInfo, nil
		}
		return nil, gerror.Wrap(err, "获取用户信息失败")
	}

	return
}

func (s *sSysUser) loginBySSO(ctx context.Context, code string) (userInfo *model.User, err error) {
	iuqtUserInfo, err := service.ThirdService().LoginByTicket(ctx, code)
	if err != nil {
		return nil, err
	}

	userInfo, err = s.GetUserByIUQTID(ctx, iuqtUserInfo.IUQTID)
	if err != nil {
		if err == sql.ErrNoRows {
			data := map[string]interface{}{
				dao.SysUser.Columns().UserName:     iuqtUserInfo.UserName,
				dao.SysUser.Columns().UserNickname: iuqtUserInfo.UserName,
				dao.SysUser.Columns().UserStatus:   model.UserStatusNormal,
				dao.SysUser.Columns().UserEmail:    "",
				dao.SysUser.Columns().Sex:          0,
				dao.SysUser.Columns().IuqtID:       iuqtUserInfo.IUQTID,
			}

			userID, err := s.registerUser(ctx, data)
			if err != nil {
				return nil, err
			}
			userInfo = &model.User{
				ID:         userID,
				UserName:   iuqtUserInfo.UserName,
				UserStatus: model.UserStatusNormal,
			}
			return userInfo, nil
		}
		return nil, gerror.Wrap(err, "获取用户信息失败")
	}

	return userInfo, nil
}

// TODO:
// 组织信息、部门信息
// 角色信息
// IsAdmin
func (s *sSysUser) registerUser(ctx context.Context, userInfo map[string]interface{}) (userID int64, err error) {
	salt := grand.S(10)
	password := libUtils.EncryptPassword("123456", salt)
	userInfo[dao.SysUser.Columns().UserPassword] = password
	userInfo[dao.SysUser.Columns().UserSalt] = salt

	result, err := dao.SysUser.Ctx(ctx).Data(userInfo).Insert()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = gerror.New("该手机号已注册,请更换手机号登录")
			return
		}
		return
	}

	userID, err = result.LastInsertId()
	if err != nil {
		return
	}

	return userID, nil
}

func (s *sSysUser) UpdateUserType(ctx context.Context, userID int64, userType string) (err error) {
	userTypeEnum := model.GetUserType(userType)
	if userTypeEnum == model.UserTypeUnknown {
		return gerror.New("无效的用户类型")
	}

	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, userID).Update(map[string]interface{}{
		dao.SysUser.Columns().UserType: userTypeEnum,
	})
	return err
}
