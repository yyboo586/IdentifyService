package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// sysNoticeDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysNoticeDao struct {
	*internal.SysNoticeDao
}

var (
	// SysNotice is globally public accessible object for table tools_gen_table operations.
	SysNotice = sysNoticeDao{
		internal.NewSysNoticeDao(),
	}
)

// Fill with you ideas below.
