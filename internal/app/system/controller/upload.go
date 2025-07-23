package controller

import (
	"IdentifyService/api/v1/system"
	commonConsts "IdentifyService/internal/app/common/consts"
	commonService "IdentifyService/internal/app/common/service"
	"IdentifyService/internal/app/system/consts"
	"IdentifyService/internal/app/system/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

var Upload = new(uploadController)

type uploadController struct{}

// 上传单图
func (c *uploadController) SingleImg(ctx context.Context, req *system.UploadSingleImgReq) (res *system.UploadSingleRes, err error) {
	file := req.File
	v, _ := g.Cfg().Get(ctx, "upload.default")
	response, err := commonService.Upload().UploadFile(ctx, file, commonConsts.CheckFileTypeImg, v.Int(), service.Context().Get(ctx).User.Id, consts.UploadAppId)
	if err != nil {
		return
	}
	res = &system.UploadSingleRes{
		UploadResponse: response,
	}
	return
}

// 上传多图
func (c *uploadController) MultipleImg(ctx context.Context, req *system.UploadMultipleImgReq) (res system.UploadMultipleRes, err error) {
	files := req.File
	res, err = commonService.Upload().UploadFiles(ctx,
		files,
		commonConsts.CheckFileTypeImg,
		g.Cfg().MustGet(ctx, "upload.default").Int(),
		service.Context().Get(ctx).User.Id,
		consts.UploadAppId)
	if err != nil {
		return
	}
	return
}

// 上传单文件
func (c *uploadController) SingleFile(ctx context.Context, req *system.UploadSingleFileReq) (res *system.UploadSingleRes, err error) {
	file := req.File
	v, _ := g.Cfg().Get(ctx, "upload.default")
	response, err := commonService.Upload().UploadFile(ctx, file, commonConsts.CheckFileTypeFile, v.Int(), service.Context().Get(ctx).User.Id, consts.UploadAppId)
	if err != nil {
		return
	}
	res = &system.UploadSingleRes{
		UploadResponse: response,
	}
	return
}

// 上传多文件
func (c *uploadController) MultipleFile(ctx context.Context, req *system.UploadMultipleFileReq) (res system.UploadMultipleRes, err error) {
	files := req.File
	res, err = commonService.Upload().UploadFiles(ctx,
		files,
		commonConsts.CheckFileTypeFile,
		g.Cfg().MustGet(ctx, "upload.default").Int(),
		service.Context().Get(ctx).User.Id,
		consts.UploadAppId)
	if err != nil {
		return
	}
	return
}

func (c *uploadController) CheckMultipart(ctx context.Context, req *system.CheckMultipartReq) (res *system.CheckMultipartRes, err error) {
	req.DriverType = g.Cfg().MustGet(ctx, "upload.default").Int()
	req.UploadType = commonConsts.CheckFileTypeFile
	req.UserId = service.Context().Get(ctx).User.Id
	req.AppId = consts.UploadAppId
	res = new(system.CheckMultipartRes)
	res.CheckMultipartRes, err = commonService.Upload().CheckMultipart(ctx, req.CheckMultipartReq)
	return
}

func (c *uploadController) UploadPart(ctx context.Context, req *system.UploadPartReq) (res *system.UploadPartRes, err error) {
	res = new(system.UploadPartRes)
	req.DriverType = g.Cfg().MustGet(ctx, "upload.default").Int()
	req.UploadType = commonConsts.CheckFileTypeFile
	req.UserId = service.Context().Get(ctx).User.Id
	req.AppId = consts.UploadAppId
	res.UploadPartRes, err = commonService.Upload().UploadPart(ctx, req.UploadPartReq)
	return
}
