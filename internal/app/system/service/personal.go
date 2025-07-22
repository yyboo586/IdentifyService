package service

import (
	"IdentifyService/internal/app/system/model"
	"context"

	"IdentifyService/api/v1/system"
)

type (
	IPersonal interface {
		GetPersonalInfo(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error)
		EditPersonal(ctx context.Context, req *system.PersonalEditReq) (user *model.LoginUserRes, err error)
		ResetPwdPersonal(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error)
	}
)

var (
	localPersonal IPersonal
)

func Personal() IPersonal {
	if localPersonal == nil {
		panic("implement not found for interface IPersonal, forgot register?")
	}
	return localPersonal
}

func RegisterPersonal(i IPersonal) {
	localPersonal = i
}
