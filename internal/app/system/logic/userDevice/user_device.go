package userDevice

import (
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sUserDevice struct{}

func init() {
	service.RegisterUserDevice(New())
}

func New() service.IUserDevice {
	return &sUserDevice{}
}

func (s *sUserDevice) RecordDevice(ctx context.Context, input *model.UserDeviceRecordInput) error {
	if input == nil {
		return gerror.New("设备信息不能为空")
	}
	if strings.TrimSpace(input.UserID) == "" || strings.TrimSpace(input.DeviceID) == "" {
		return gerror.New("用户ID或设备ID不能为空")
	}

	now := time.Now().Unix()
	data := g.Map{
		dao.UserDevice.Columns().UserId:     input.UserID,
		dao.UserDevice.Columns().DeviceId:   input.DeviceID,
		dao.UserDevice.Columns().DeviceName: input.DeviceName,
		dao.UserDevice.Columns().DeviceIp:   input.DeviceIP,
		dao.UserDevice.Columns().LoginType:  input.LoginType,
		dao.UserDevice.Columns().CreatedAt:  now,
		dao.UserDevice.Columns().UpdatedAt:  now,
	}

	device, exists, err := dao.UserDevice.GetByUserAndDevice(ctx, input.UserID, input.DeviceID)
	if err != nil {
		return err
	}
	if !exists { // 设备不存在，则插入
		data[dao.UserDevice.Columns().CreatedAt] = now
		return dao.UserDevice.Insert(ctx, data)
	}

	// 设备存在，则更新
	return dao.UserDevice.Update(ctx, device.Id, data)
}

func (s *sUserDevice) GetUserDevice(ctx context.Context, userID, deviceID string) (*model.UserDevice, error) {
	if strings.TrimSpace(userID) == "" || strings.TrimSpace(deviceID) == "" {
		return nil, gerror.New("用户ID或设备ID不能为空")
	}
	device, exists, err := dao.UserDevice.GetByUserAndDevice(ctx, userID, deviceID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, gerror.New("设备不存在")
	}
	return convertUserDeviceEntity(device), nil
}

func (s *sUserDevice) ListUserDevices(ctx context.Context, userID string, pageReq *model.PageReq) ([]*model.UserDevice, *model.PageRes, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, nil, gerror.New("用户ID不能为空")
	}

	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	entities, total, err := dao.UserDevice.ListByUser(ctx, userID, pageReq.Page, pageReq.Size)
	if err != nil {
		return nil, nil, err
	}

	var list []*model.UserDevice
	for _, item := range entities {
		list = append(list, convertUserDeviceEntity(item))
	}
	return list, &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}, nil
}

func (s *sUserDevice) DeleteUserDevice(ctx context.Context, userID, deviceID string) error {
	if strings.TrimSpace(userID) == "" || strings.TrimSpace(deviceID) == "" {
		return gerror.New("用户ID或设备ID不能为空")
	}
	return dao.UserDevice.Delete(ctx, userID, deviceID)
}

func convertUserDeviceEntity(entity *entity.UserDevice) *model.UserDevice {
	if entity == nil {
		return nil
	}
	return &model.UserDevice{
		ID:         entity.Id,
		UserID:     entity.UserId,
		DeviceID:   entity.DeviceId,
		DeviceName: entity.DeviceName,
		DeviceIP:   entity.DeviceIp,
		LoginType:  entity.LoginType,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}
