package controller

import (
	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/service"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yyboo586/common/MiddleWare"
)

var UserDevice = userDeviceController{}

type userDeviceController struct {
	BaseController
}

func (c *userDeviceController) ListUserDevices(ctx context.Context, req *system.ListUserDeviceReq) (res *system.ListUserDeviceRes, err error) {
	devices, pageRes, err := service.UserDevice().ListUserDevices(ctx, req.UserID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	items := make([]*system.UserDeviceItem, 0, len(devices))
	for _, device := range devices {
		createTime := ""
		if device.CreatedAt != nil {
			createTime = device.CreatedAt.Format("Y-m-d H:i:s")
		}
		items = append(items, &system.UserDeviceItem{
			DeviceID:   device.DeviceID,
			DeviceName: device.DeviceName,
			DeviceIP:   device.DeviceIP,
			CreateTime: createTime,
			LoginType:  device.LoginType,
		})
	}

	return &system.ListUserDeviceRes{
		List:    items,
		PageRes: pageRes,
	}, nil
}

func (c *userDeviceController) DeleteUserDevice(ctx context.Context, req *system.DeleteUserDeviceReq) (res *system.DeleteUserDeviceRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	g.Log().Info(ctx, operator.UserID, operator.UserName)

	device, err := service.UserDevice().GetUserDevice(ctx, req.UserID, req.DeviceID)
	if err != nil {
		return nil, err
	}

	// TODO: 事务
	err = service.Token().RevokeDeviceToken(ctx, req.DeviceID)
	if err != nil {
		return nil, err
	}

	err = service.UserDevice().DeleteUserDevice(ctx, req.UserID, req.DeviceID)
	if err != nil {
		return nil, err
	}

	service.MQ().Publish(ctx, "core.users.notify", []string{req.UserID}, map[string]interface{}{
		"title":       "强制设备下线",
		"user_id":     req.UserID,
		"device_id":   req.DeviceID,
		"device_name": device.DeviceName,
		"force_time":  time.Now().Format("2006-01-02 15:04:05"),
	})

	return &system.DeleteUserDeviceRes{}, nil
}
