package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDevice is the golang structure of table t_user_device for DAO operations.
type UserDevice struct {
	g.Meta     `orm:"table:t_user_device, do:true"`
	Id         interface{}
	UserId     interface{}
	DeviceId   interface{}
	DeviceName interface{}
	DeviceIp   interface{}
	LoginType  interface{}
	CreatedAt  *gtime.Time
	UpdatedAt  *gtime.Time
}
