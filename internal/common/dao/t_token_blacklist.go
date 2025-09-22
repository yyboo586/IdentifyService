package dao

import (
	"IdentifyService/internal/common/dao/internal"
)

// tokenBlacklistDao is the data access object for table t_token_blacklist.
// You can define custom methods on it to extend its functionality as you wish.
type tokenBlacklistDao struct {
	*internal.TokenBlacklistDao
}

var (
	// TokenBlacklist is globally public accessible object for table t_token_blacklist operations.
	TokenBlacklist = tokenBlacklistDao{
		internal.NewTokenBlacklistDao(),
	}
)
