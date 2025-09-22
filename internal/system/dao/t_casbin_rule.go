package dao

import (
	"IdentifyService/internal/system/dao/internal"
)

// casbinRuleDao is the data access object for table t_casbin_rule.
// You can define custom methods on it to extend its functionality as you wish.
type casbinRuleDao struct {
	*internal.CasbinRuleDao
}

var (
	// CasbinRule is globally public accessible object for table t_casbin_rule operations.
	CasbinRule = casbinRuleDao{
		internal.NewCasbinRuleDao(),
	}
)

// Fill with you ideas below.
