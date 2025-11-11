package cache

import (
	"IdentifyService/internal/app/common/consts"
	"IdentifyService/internal/app/common/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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
