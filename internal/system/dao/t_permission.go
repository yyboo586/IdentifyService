package dao

import (
	"IdentifyService/internal/system/dao/internal"
)

// permissionDao is the data access object for table t_permission.
type permissionDao struct {
	*internal.PermissionDao
}

var (
	// Permission is globally public accessible object for table t_permission operations.
	Permission = permissionDao{
		internal.NewPermissionDao(),
	}
)
