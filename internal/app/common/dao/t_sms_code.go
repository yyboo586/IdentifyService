package dao

import (
	"IdentifyService/internal/app/common/dao/internal"
)

type TSmsCodeDao struct {
	*internal.TSmsCodeDao
}

var (
	TSmsCode = TSmsCodeDao{
		internal.NewTSmsCodeDao(),
	}
)
