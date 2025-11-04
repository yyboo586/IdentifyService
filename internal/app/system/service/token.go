package service

import (
	"github.com/yyboo586/common/authUtils/tokenUtils"
)

var t tokenUtils.IToken

func RegisterToken(in tokenUtils.IToken) {
	t = in
}

func Token() tokenUtils.IToken {
	if t == nil {
		panic("implement not found for interface tokenUtils.IToken, forgot register?")
	}
	return t
}
