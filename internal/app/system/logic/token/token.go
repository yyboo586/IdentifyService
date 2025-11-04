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
	"github.com/tiger1103/gfast/v3/library/liberr"
	"github.com/yyboo586/common/authUtils/tokenUtils"
)

type sToken struct {
	*tokenUtils.Token
}

func New() service.IGfToken {
	var (
		ctx = gctx.New()
		opt *commonModel.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
	)
	liberr.ErrIsNil(ctx, err)

	return &sToken{
		Token: tokenUtils.NewToken(
			tokenUtils.WithExcludePaths(opt.ExcludePaths),
		),
	}
}

func init() {
	service.RegisterGToken(New())
}
