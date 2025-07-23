package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// roleDao is the data access object for table t_role.
// You can define custom methods on it to extend its functionality as you wish.
type roleDao struct {
	*internal.RoleDao
}

var (
	// SysRole is globally public accessible object for table t_role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// Fill with you ideas below.
