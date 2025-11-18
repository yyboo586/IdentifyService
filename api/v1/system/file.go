package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type PreUploadFileReq struct {
	g.Meta `path:"/file/pre-upload" tags:"文件管理" method:"post" summary:"获取文件上传地址"`
	model.Author
	FileName    string `json:"file_name" v:"required#文件名不能为空" dc:"文件名"`
	ContentType string `json:"content_type" v:"required#文件类型不能为空" dc:"文件类型"`
	FileSize    int64  `json:"file_size" v:"required#文件大小不能为空" dc:"文件大小"`
}

type PreUploadFileRes struct {
	g.Meta    `mime:"application/json"`
	FileID    string `json:"file_id" dc:"文件ID"`
	UploadUrl string `json:"upload_url" dc:"上传URL"`
	VisitLink string `json:"visit_link" dc:"访问链接"`
}

type ReportUploadResultReq struct {
	g.Meta `path:"/file/:file_id/status" tags:"文件管理" method:"patch" summary:"上报文件上传结果"`
	model.Author
	FileID  string `p:"file_id" v:"required#文件ID不能为空" dc:"文件ID"`
	Success bool   `json:"success" v:"required#上传结果不能为空" dc:"上传结果"`
}

type ReportUploadResultRes struct {
	g.Meta `mime:"application/json"`
}

type UploadFileReq struct {
	g.Meta `path:"/file/upload" mine:"multipart/form-data" tags:"文件管理" method:"post" summary:"上传文件"`
	model.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件" v:"required#上传文件必须"`
}

type UploadFileRes struct {
	g.Meta   `mime:"application/json"`
	FileID   string `json:"file_id" dc:"文件ID"`
	FileLink string `json:"file_link" dc:"文件链接"`
}
