package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// loginLogDao is the data access object for table t_login_log.
// You can define custom methods on it to extend its functionality as you wish.
type loginLogDao struct {
	*internal.LoginLogDao
}

var (
	// LoginLog is globally public accessible object for table t_login_log operations.
	LoginLog = loginLogDao{
		internal.NewLoginLogDao(),
	}
)

// Fill with you ideas below.
