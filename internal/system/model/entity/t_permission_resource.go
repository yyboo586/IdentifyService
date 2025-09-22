package entity

import "github.com/gogf/gf/v2/os/gtime"

type PermissionResource struct {
	ID           int64       `orm:"id"`
	PermissionID int64       `orm:"permission_id"`
	ResourceID   int64       `orm:"resource_id"`
	CreatedAt    *gtime.Time `orm:"created_at"`
}
