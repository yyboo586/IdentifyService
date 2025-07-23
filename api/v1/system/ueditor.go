package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UEditorConfigReq struct {
	g.Meta `path:"/uEditor/action" tags:"系统后台/UEditor" method:"get" summary:"获取UEditor配置"`
	UEditorReq
}

type UEditorUpFileReq struct {
	g.Meta `path:"/uEditor/action" tags:"系统后台/UEditor" method:"post" summary:"UEditor上传"`
	UEditorReq
}

type UEditorReq struct {
	Action   string            `p:"action"`
	Callback string            `p:"callback"`
	File     *ghttp.UploadFile `p:"upfile" type:"file"`
	Start    int               `p:"start"`
	Size     int               `p:"size"`
}

type UEditorRes struct {
	g.Meta `mime:"application/json"`
	g.Map
}
