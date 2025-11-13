package sysUser

import (
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/do"
	"IdentifyService/library/libSecurity"
	"IdentifyService/library/libUtils"
	"context"
	"strings"
	"unicode/utf8"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/google/uuid"
)

func (s *sSysUser) GetUserInfoByPhone(ctx context.Context, phone string) (out *model.User, exists bool, err error) {
	userEntity, exists, err := dao.SysUser.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, false, err
	}
	if !exists {
		return nil, false, nil
	}

	return model.ConvertToUser(userEntity), true, nil
}

func (s *sSysUser) SelfRegister(ctx context.Context, tx gdb.TX, userID string, deptID int64, phone string) (err error) {
	err = s.CheckNameOrPhoneExists(ctx, "", phone)
	if err != nil {
		return err
	}

	salt := grand.S(10)
	password := libUtils.EncryptPassword("123456", salt)
	dataInsert := map[string]interface{}{
		dao.SysUser.Columns().Id:           userID,
		dao.SysUser.Columns().DeptId:       deptID,
		dao.SysUser.Columns().UserName:     "用户" + phone,
		dao.SysUser.Columns().UserNickname: "用户" + phone,
		dao.SysUser.Columns().UserPassword: password,
		dao.SysUser.Columns().UserSalt:     salt,
		dao.SysUser.Columns().UserStatus:   model.UserStatusNormal,
		dao.SysUser.Columns().IsAdmin:      0,

		dao.SysUser.Columns().Mobile:    phone,
		dao.SysUser.Columns().UserEmail: "",
		dao.SysUser.Columns().Sex:       0,
		dao.SysUser.Columns().Birthday:  "",
		dao.SysUser.Columns().City:      "",
		dao.SysUser.Columns().Avatar:    "",
		dao.SysUser.Columns().CardType:  "",
		dao.SysUser.Columns().RealName:  "",
		dao.SysUser.Columns().IDCard:    "",
	}

	err = dao.SysUser.Insert(ctx, tx, dataInsert)
	if err != nil {
		return err
	}

	return nil
}

func (s *sSysUser) UnRegister(ctx context.Context, userID string) (err error) {
	err = dao.SysUser.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysUser) CheckNameOrPhoneExists(ctx context.Context, name, mobile string) error {
	query := dao.SysUser.Ctx(ctx)
	if name != "" {
		query = query.Where(dao.SysUser.Columns().UserName, name)
	}
	if mobile != "" {
		query = query.Where(dao.SysUser.Columns().Mobile, mobile)
	}
	num, err := query.Count()
	if err != nil {
		return err
	}
	if num != 0 {
		return gerror.New("用户名或手机号已存在")
	}

	return nil
}

func (s *sSysUser) GetUserPersonalInfo(ctx context.Context, userID string) (userInfo *model.User, err error) {
	userEntity, err := dao.SysUser.GetUserPersonalInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	userInfo = model.ConvertToUser(userEntity)
	return userInfo, nil
}

func (s *sSysUser) EditUserPersonalInfo(ctx context.Context, userID string, userInfo *model.UserPersonalInfo) (err error) {
	if !model.IsValidUserSex(model.UserSex(userInfo.Sex)) {
		return gerror.New("性别格式不正确,请输入0:保密,1:男,2:女")
	}

	// 校验Birthday是否是一个合法的日期
	if userInfo.Birthday != "" {
		_, err := gtime.StrToTimeFormat(userInfo.Birthday, "Y-m-d")
		if err != nil {
			return gerror.New("生日格式不正确,请输入YYYY-MM-DD格式的日期")
		}
	}

	// 校验City长度是否超过10
	if utf8.RuneCountInString(userInfo.City) > 10 {
		return gerror.New("城市名称长度不能超过10个字符")
	}

	updateData := map[string]interface{}{
		dao.SysUser.Columns().UserNickname: userInfo.Nickname,
		dao.SysUser.Columns().Avatar:       userInfo.Avatar,
		dao.SysUser.Columns().Sex:          userInfo.Sex,
		dao.SysUser.Columns().Birthday:     userInfo.Birthday,
		dao.SysUser.Columns().City:         userInfo.City,
	}

	err = dao.SysUser.Update(ctx, userID, updateData)
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysUser) UpdateUserPhone(ctx context.Context, userID string, phone string) (err error) {
	if phone != "" {
		if !ValidateChinaMobile(phone) {
			return gerror.New("手机号格式不正确, 请输入11位中国大陆手机号")
		}
	}

	updateData := map[string]interface{}{
		dao.SysUser.Columns().Mobile: phone,
	}
	err = dao.SysUser.Update(ctx, userID, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (s *sSysUser) EditUserPassword(ctx context.Context, userID string, phone string, password string) (err error) {
	userInfo, exists, err := s.GetUserInfoByPhone(ctx, phone)
	if err != nil {
		return err
	}
	if !exists {
		return gerror.New("用户不存在")
	}
	if userInfo.Mobile == "" {
		return gerror.New("手机号未绑定, 请先绑定手机号")
	}
	if userInfo.Mobile != phone {
		return gerror.New("手机号不匹配, 请使用绑定的手机号进行修改密码")
	}

	if len(password) < 6 {
		return gerror.New("密码长度不能小于6位")
	}
	if len(password) > 16 {
		return gerror.New("密码长度不能大于16位")
	}

	salt := grand.S(10)
	passwordHash := libUtils.EncryptPassword(password, salt)
	updateData := map[string]interface{}{
		dao.SysUser.Columns().UserPassword: passwordHash,
		dao.SysUser.Columns().UserSalt:     salt,
	}
	err = dao.SysUser.Update(ctx, userID, updateData)
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysUser) EditUserIDCard(ctx context.Context, userID string, phone string, idCard string, cardType string, realName string) (err error) {
	userEntity, err := dao.SysUser.Get(ctx, userID)
	if err != nil {
		return err
	}
	if userEntity.Mobile == "" {
		return gerror.New("手机号未绑定, 请先绑定手机号")
	}
	if userEntity.Mobile != phone {
		return gerror.New("手机号不匹配, 请使用绑定的手机号进行修改身份证号")
	}

	if !model.IsValidCardType(cardType) {
		return gerror.New("证件类型不正确")
	}

	realName = strings.TrimSpace(realName)
	if realName == "" {
		return gerror.New("真实姓名不能为空")
	}

	cipherText, err := libSecurity.EncryptIDCard(idCard)
	if err != nil {
		return err
	}
	updateData := map[string]interface{}{
		dao.SysUser.Columns().IDCard:   cipherText,
		dao.SysUser.Columns().CardType: cardType,
		dao.SysUser.Columns().RealName: realName,
	}
	err = dao.SysUser.Update(ctx, userID, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (s *sSysUser) GetUserIDCard(ctx context.Context, userID string) (info *model.UserIDCardInfo, err error) {
	userEntity, err := dao.SysUser.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	userInfo := model.ConvertToUser(userEntity)
	info = &model.UserIDCardInfo{
		CardType: userEntity.CardType,
		IDCard:   libSecurity.MaskIDCard(userInfo.IDCard),
		RealName: libSecurity.MaskRealName(userInfo.RealName),
	}
	return info, nil
}

// ValidateChinaMobile 校验手机号是否符合中国大陆手机号格式
// 规则：11位数字，以1开头，第二位是3-9
func ValidateChinaMobile(mobile string) bool {
	if mobile == "" {
		return true // 空值不校验，由业务逻辑决定是否必填
	}
	// 正则表达式：^1[3-9]\d{9}$
	return gregex.IsMatchString(`^1[3-9]\d{9}$`, mobile)
}

func checkUserNameExists(ctx context.Context, userName string) (err error) {
	user, err := dao.SysUser.GetUserByUserName(ctx, userName)
	if err != nil {
		return err
	}
	if user != nil {
		return gerror.New("用户名已存在")
	}
	return nil
}

func (s *sSysUser) Add3(ctx context.Context, deptID int64, phone, userNickname string) (userID string, err error) {
	_, exists, err := dao.SysUser.GetUserByPhone(ctx, phone)
	if err != nil {
		return "", err
	}
	if exists {
		return "", gerror.New("手机号已存在")
	}

	userID = uuid.New().String()
	salt := grand.S(10)
	password := libUtils.EncryptPassword("123456", salt)
	_, err = dao.SysUser.Ctx(ctx).Insert(do.SysUser{
		Id:           userID,
		UserName:     "用户" + phone,
		UserNickname: userNickname,
		UserPassword: password,
		UserSalt:     salt,
		UserStatus:   model.UserStatusNormal,
		DeptId:       deptID,
		IsAdmin:      false,
	})
	return
}

func (s *sSysUser) Delete3(ctx context.Context, userID string) (err error) {
	err = dao.SysUser.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
