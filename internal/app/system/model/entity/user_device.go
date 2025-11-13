package entity

import "github.com/gogf/gf/v2/os/gtime"

// UserDevice 对应数据表 t_user_device
type UserDevice struct {
	Id         int64       `orm:"id"`
	UserId     string      `orm:"user_id"`
	DeviceId   string      `orm:"device_id"`
	DeviceName string      `orm:"device_name"`
	DeviceIp   string      `orm:"device_ip"`
	LoginType  string      `orm:"login_type"`
	CreatedAt  *gtime.Time `orm:"created_at"`
	UpdatedAt  *gtime.Time `orm:"updated_at"`
}
