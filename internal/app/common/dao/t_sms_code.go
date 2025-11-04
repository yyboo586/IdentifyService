package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)

type TSmsCodeDao struct {
	*internal.TSmsCodeDao
}

var (
	TSmsCode = TSmsCodeDao{
		internal.NewTSmsCodeDao(),
	}
)
