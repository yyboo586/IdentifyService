package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

type userDao struct {
	*internal.UserDao
}

var (
	User = userDao{
		internal.NewUserDao(),
	}
)
