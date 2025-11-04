/*
* @desc:token功能
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/8 15:54
 */

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yyboo586/common/authUtils/tokenUtils"
)

type IGfToken interface {
	Generate(ctx context.Context, data interface{}) (token string, err error)
	Refresh(ctx context.Context, token string) (newToken string, err error)

	GetTokenFromRequest(r *ghttp.Request) (token string)
	Parse(r *ghttp.Request) (*tokenUtils.CustomClaims, error)
	Middleware(group *ghttp.RouterGroup) error
}

var gt IGfToken

func RegisterGToken(gtk IGfToken) {
	gt = gtk
}

func GfToken() IGfToken {
	if gt == nil {
		panic("implement not found for interface IGfToken, forgot register?")
	}
	return gt
}
