package service

import (
	"IdentifyService/internal/app/common/model"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IUpload interface {
	UploadFiles(ctx context.Context, files []*ghttp.UploadFile, checkFileType string, source int, userId uint64, appId string) (result []*model.UploadResponse, err error)
	UploadFile(ctx context.Context, file *ghttp.UploadFile, checkFileType string, source int, userId uint64, appId string) (result *model.UploadResponse, err error)
	CheckSize(ctx context.Context, checkFileType string, fileSize int64) (err error)
	CheckType(ctx context.Context, checkFileType string, fileName string) (err error)
	CheckMultipart(ctx context.Context, req *model.CheckMultipartReq) (res *model.CheckMultipartRes, err error)
	UploadPart(ctx context.Context, req *model.UploadPartReq) (res *model.UploadPartRes, err error)
}

var localUpload IUpload

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}
