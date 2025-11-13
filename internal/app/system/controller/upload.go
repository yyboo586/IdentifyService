package controller

import (
	"context"

	"IdentifyService/api/v1/system"
	commonConsts "IdentifyService/internal/app/common/consts"
	commonService "IdentifyService/internal/app/common/service"
	"IdentifyService/internal/app/system/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yyboo586/common/MiddleWare"
)

var Upload = new(uploadController)

type uploadController struct{}

// 上传单图
func (c *uploadController) SingleImg(ctx context.Context, req *system.UploadSingleImgReq) (res *system.UploadSingleRes, err error) {
	file := req.File
	v, _ := g.Cfg().Get(ctx, "upload.default")
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	response, err := commonService.Upload().UploadFile(ctx, file, commonConsts.CheckFileTypeImg, v.Int(), operator.UserID, consts.UploadAppId)
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
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	res, err = commonService.Upload().UploadFiles(ctx,
		files,
		commonConsts.CheckFileTypeImg,
		g.Cfg().MustGet(ctx, "upload.default").Int(),
		operator.UserID,
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
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	response, err := commonService.Upload().UploadFile(ctx, file, commonConsts.CheckFileTypeFile, v.Int(), operator.UserID, consts.UploadAppId)
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
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	res, err = commonService.Upload().UploadFiles(ctx,
		files,
		commonConsts.CheckFileTypeFile,
		g.Cfg().MustGet(ctx, "upload.default").Int(),
		operator.UserID,
		consts.UploadAppId)
	if err != nil {
		return
	}
	return
}

func (c *uploadController) CheckMultipart(ctx context.Context, req *system.CheckMultipartReq) (res *system.CheckMultipartRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	req.DriverType = g.Cfg().MustGet(ctx, "upload.default").Int()
	req.UploadType = commonConsts.CheckFileTypeFile
	req.UserId = operator.UserID
	req.AppId = consts.UploadAppId
	res = new(system.CheckMultipartRes)
	res.CheckMultipartRes, err = commonService.Upload().CheckMultipart(ctx, req.CheckMultipartReq)
	return
}

func (c *uploadController) UploadPart(ctx context.Context, req *system.UploadPartReq) (res *system.UploadPartRes, err error) {
	operator, err := MiddleWare.GetContextUser(ctx)
	if err != nil {
		return nil, err
	}
	res = new(system.UploadPartRes)
	req.DriverType = g.Cfg().MustGet(ctx, "upload.default").Int()
	req.UploadType = commonConsts.CheckFileTypeFile
	req.UserId = operator.UserID
	req.AppId = consts.UploadAppId
	res.UploadPartRes, err = commonService.Upload().UploadPart(ctx, req.UploadPartReq)
	return
}
