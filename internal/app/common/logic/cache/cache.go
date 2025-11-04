/*
* @desc:缓存处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/27 16:33
 */

package cache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast/v3/internal/app/common/consts"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/yyboo586/common/cacheUtils"
)

func init() {
	service.RegisterCache(New())
}

type Cache struct {
	cacheUtils.ICache
	prefix string
}

func New() cacheUtils.ICache {
	var (
		ctx            = gctx.New()
		cacheContainer cacheUtils.ICache
	)
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	model := g.Cfg().MustGet(ctx, "system.cache.model").String()

	switch model {
	case consts.CacheModelRedis:
		cacheContainer = cacheUtils.NewRedis(prefix)
	case consts.CacheModelMem:
		cacheContainer = cacheUtils.NewMemory(prefix)
	default:
		panic("invalid cache model, only support redis and memory")
	}
	return &Cache{
		ICache: cacheContainer,
		prefix: prefix,
	}
}
