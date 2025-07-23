package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// internalSysOperLogDao is internal type for wrapping internal DAO implements.
type internalOperLogDao = *internal.OperLogDao

// operLogDao is the data access object for table sys_oper_log.
// You can define custom methods on it to extend its functionality as you wish.
type operLogDao struct {
	internalOperLogDao
}

var (
	// OperLog is globally public accessible object for table sys_oper_log operations.
	OperLog = operLogDao{
		internal.NewOperLogDao(),
	}
)

// Fill with you ideas below.
