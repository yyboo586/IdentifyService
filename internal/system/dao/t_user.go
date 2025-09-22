package dao

import (
	"IdentifyService/internal/system/dao/internal"
)

type userDao struct {
	*internal.UserDao
}

var (
	User = userDao{
		internal.NewUserDao(),
	}
)
