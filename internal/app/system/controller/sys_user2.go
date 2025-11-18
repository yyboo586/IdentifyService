package controller

import (
	"context"
	"time"

	"IdentifyService/api/v1/system"
	commonModel "IdentifyService/internal/app/common/model"
	commonService "IdentifyService/internal/app/common/service"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/library/libSecurity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yyboo586/common/MiddleWare"
)

func (c *userController) UpdateUserType(ctx context.Context, req *system.UpdateUserTypeReq) (res *system.UpdateUserTypeRes, err error) {
	err = service.SysUser().UpdateUserType(ctx, req.UserID, req.UserType)
	return
}

func (c *userController) GetUserPersonalInfo(ctx context.Context, req *system.GetUserPersonalInfoReq) (res *system.GetUserPersonalInfoRes, err error) {
	userInfo, err := service.SysUser().GetUserPersonalInfo(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	res = &system.GetUserPersonalInfoRes{
		UserPersonalInfo: &system.UserPersonalInfo{
			Nickname:          userInfo.UserNickname,
			Avatar:            userInfo.Avatar,
			Mobile:            userInfo.Mobile,
			Sex:               userInfo.Sex,
			Birthday:          userInfo.Birthday,
			City:              userInfo.City,
			CreateTime:        userInfo.CreatedAt.Time.Format(time.DateTime),
			IsAlreadyRealname: userInfo.IDCard != "",
			UserRealName:      libSecurity.MaskRealName(userInfo.RealName),
		},
	}
	return res, nil
}

func (c *userController) EditUserPersonalInfo(ctx context.Context, req *system.EditUserPersonalInfoReq) (res *system.EditUserPersonalInfoRes, err error) {
	in := &model.UserPersonalInfo{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
		Birthday: req.Birthday,
		City:     req.City,
	}
	err = service.SysUser().EditUserPersonalInfo(ctx, req.UserID, in)
	if err != nil {
		return nil, err
	}

	return &system.EditUserPersonalInfoRes{}, nil
}

func (c *userController) BindPhone(ctx context.Context, req *system.BindPhoneReq) (res *system.BindPhoneRes, err error) {
	err = commonService.Captcha().ValidateSmsCode(ctx, req.Phone, commonModel.SMSBusinessTypeBindPhone, req.Code)
	if err != nil {
		return nil, err
	}

	err = service.SysUser().UpdateUserPhone(ctx, req.UserID, req.Phone)
	if err != nil {
		return nil, err
	}

	return &system.BindPhoneRes{}, nil
}

func (c *userController) EditUserPassword(ctx context.Context, req *system.EditUserPasswordReq) (res *system.EditUserPasswordRes, err error) {
	err = service.SysUser().EditUserPassword(ctx, req.UserID, req.Phone, req.NewPassword)
	if err != nil {
		return nil, err
	}

	return &system.EditUserPasswordRes{}, nil
}

func (c *userController) EditUserIDCard(ctx context.Context, req *system.EditUserIDCardReq) (res *system.EditUserIDCardRes, err error) {
	err = service.SysUser().EditUserIDCard(ctx, req.UserID, req.IDCard, req.CardType, req.RealName)
	if err != nil {
		return nil, err
	}

	return &system.EditUserIDCardRes{}, nil
}

func (c *userController) GetUserIDCard(ctx context.Context, req *system.GetUserIDCardReq) (res *system.GetUserIDCardRes, err error) {
	info, err := service.SysUser().GetUserIDCard(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &system.GetUserIDCardRes{
		CardType: info.CardType,
		IDCard:   info.IDCard,
		RealName: info.RealName,
	}, nil
}

func (c *userController) Add3(ctx context.Context, req *system.Add3Req) (res *system.Add3Res, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	g.Log().Info(ctx, operator.DeptID)
	userID, err := service.SysUser().Add3(ctx, int64(operator.DeptID), req.Phone, req.UserNickname)
	if err != nil {
		return nil, err
	}

	return &system.Add3Res{
		UserID: userID,
	}, nil
}

func (c *userController) Delete3(ctx context.Context, req *system.Delete3Req) (res *system.Delete3Res, err error) {
	err = service.SysUser().Delete3(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &system.Delete3Res{}, nil
}
