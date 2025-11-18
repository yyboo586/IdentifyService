package service

import (
	"IdentifyService/internal/app/system/model"
	"context"
)

type (
	IUserDevice interface {
		RecordDevice(ctx context.Context, input *model.UserDeviceRecordInput) error
		GetUserDevice(ctx context.Context, userID, deviceID string) (*model.UserDevice, error)
		ListUserDevices(ctx context.Context, userID string, pageReq *model.PageReq) ([]*model.UserDevice, *model.PageRes, error)
		DeleteUserDevice(ctx context.Context, userID, deviceID string) error
	}
)

var (
	localUserDevice IUserDevice
)

func UserDevice() IUserDevice {
	if localUserDevice == nil {
		panic("implement not found for interface IUserDevice, forgot register?")
	}
	return localUserDevice
}

func RegisterUserDevice(srv IUserDevice) {
	localUserDevice = srv
}
