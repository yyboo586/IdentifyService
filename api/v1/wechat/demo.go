package wechat

import (
	"IdentifyService/api/v1/common"

	"github.com/gogf/gf/v2/frame/g"
)

type DemoReq struct {
	g.Meta `path:"/demo" tags:"微信接口/小程序测试" method:"get" summary:"测试"`
	common.Author
}

type DemoRes struct {
	common.EmptyRes
	Info string `json:"info"`
}
