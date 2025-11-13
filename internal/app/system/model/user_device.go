package model

import "github.com/gogf/gf/v2/os/gtime"

type UserDevice struct {
	ID         int64       `json:"id"`
	UserID     string      `json:"user_id"`
	DeviceID   string      `json:"device_id"`
	DeviceName string      `json:"device_name"`
	DeviceIP   string      `json:"device_ip"`
	LoginType  string      `json:"login_type"`
	CreatedAt  *gtime.Time `json:"created_at"`
	UpdatedAt  *gtime.Time `json:"updated_at"`
}

type UserDeviceRecordInput struct {
	UserID     string
	DeviceID   string
	DeviceName string
	DeviceIP   string
	LoginType  string
}
