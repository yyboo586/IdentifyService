package system

import (
	v1 "IdentifyService/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

type CacheRemoveReq struct {
	g.Meta `path:"/cache/remove" tags:"系统后台/缓存管理" method:"delete" summary:"清除缓存"`
	v1.Author
}

type CacheRemoveRes struct {
	v1.EmptyRes
}
