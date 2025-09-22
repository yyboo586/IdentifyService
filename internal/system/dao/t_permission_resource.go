package dao

import "IdentifyService/internal/system/dao/internal"

var (
	PermissionResource = permissionResourceDao{
		internal.NewPermissionResourceDao(),
	}
)

type permissionResourceDao struct {
	*internal.PermissionResourceDao
}
