/*
* @desc:token功能
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/27 17:01
 */

package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	commonModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
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
		),
	}
}

func init() {
	service.RegisterToken(New())
}
