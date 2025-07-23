package dao

import (
	"IdentifyService/internal/app/system/dao/internal"
)

// authRuleDao is the data access object for table t_auth_rule.
// You can define custom methods on it to extend its functionality as you wish.
type authRuleDao struct {
	*internal.AuthRuleDao
}

var (
	// AuthRule is globally public accessible object for table t_auth_rule operations.
	AuthRule = authRuleDao{
		internal.NewAuthRuleDao(),
	}
)

// Fill with you ideas below.
