package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// sysNoticeReadDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysNoticeReadDao struct {
	*internal.SysNoticeReadDao
}

var (
	// SysNoticeRead is globally public accessible object for table tools_gen_table operations.
	SysNoticeRead = sysNoticeReadDao{
		internal.NewSysNoticeReadDao(),
	}
)

// Fill with you ideas below.
