package token

import (
	commonModel "IdentifyService/internal/app/common/model"
	"IdentifyService/internal/app/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yyboo586/common/authUtils/tokenUtils"
)

type token struct {
	*tokenUtils.Token
}

func New() tokenUtils.IToken {
	var (
		ctx = gctx.New()
		opt *commonModel.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
	)

	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}

	return &token{
		Token: tokenUtils.NewToken(
			tokenUtils.WithExcludePaths(opt.ExcludePaths),
			tokenUtils.WithTokenStoreConfig(&tokenUtils.TokenStoreConfig{
				DSN:         g.Cfg().MustGet(ctx, "database.default.link").String(),
				Group:       "default",
				TableName:   "t_token",
				EnableDebug: true,
			}),
		),
	}
}

func init() {
	service.RegisterToken(New())
}
