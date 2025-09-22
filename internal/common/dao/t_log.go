package dao

import (
	"IdentifyService/internal/common/dao/internal"
)

// internalLogDao is internal type for wrapping internal DAO implements.
type internalLogDao = *internal.LogDao

// logDao is the data access object for table sys_oper_log.
// You can define custom methods on it to extend its functionality as you wish.
type logDao struct {
	internalLogDao
}

var (
	// Log is globally public accessible object for table sys_oper_log operations.
	Log = logDao{
		internal.NewLogDao(),
	}
)

// Fill with you ideas below.
