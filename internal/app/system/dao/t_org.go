package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// internalOrgDao is internal type for wrapping internal DAO implements.
type internalOrgDao = *internal.OrgDao

// orgDao is the data access object for table t_org.
// You can define custom methods on it to extend its functionality as you wish.
type orgDao struct {
	internalOrgDao
}

var (
	// Org is globally public accessible object for table t_org operations.
	Org = orgDao{
		internal.NewOrgDao(),
	}
)

// Fill with you ideas below.
