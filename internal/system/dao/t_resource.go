package dao

import "IdentifyService/internal/system/dao/internal"

type resourceDao struct {
	*internal.ResourceDao
}

var (
	Resource = resourceDao{
		internal.NewResourceDao(),
	}
)
