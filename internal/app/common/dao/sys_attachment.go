package dao

import (
	"IdentifyService/internal/app/common/dao/internal"
)

// sysAttachmentDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysAttachmentDao struct {
	*internal.SysAttachmentDao
}

var (
	// SysAttachment is globally public accessible object for table tools_gen_table operations.
	SysAttachment = sysAttachmentDao{
		internal.NewSysAttachmentDao(),
	}
)

// Fill with you ideas below.
