package controller

import (
	"context"
	"errors"
	"io"
	"net/url"
	"sync"

	"IdentifyService/api/v1/system"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yyboo586/common/FileModule"
	"github.com/yyboo586/common/httpUtils"
)

var once sync.Once
var client httpUtils.HTTPClient
var fileManager FileModule.IFileManager

var File = new(fileController)

type fileController struct {
}

func InitFileManager() {
	once.Do(func() {
		var err error
		fileManager, err = FileModule.NewFileManager(&FileModule.Config{
			DSN:            g.Cfg().MustGet(context.Background(), "database.default.link").String(),
			Database:       "game_admin",
			Group:          "default",
			FileEngineAddr: g.Cfg().MustGet(context.Background(), "server.thirdService.fileEngine.addr").String(),
			EnableDebug:    true,
		})
		client = httpUtils.NewHTTPClient()
		if err != nil {
			panic(err)
		}
	})
}

func (c *fileController) PreUploadFile(ctx context.Context, req *system.PreUploadFileReq) (res *system.PreUploadFileRes, err error) {
	if fileManager == nil {
		InitFileManager()
	}

	preUploadRes, err := fileManager.PreUpload(ctx, &FileModule.PreUploadReq{
		FileName:    req.FileName,
		ContentType: req.ContentType,
		Size:        req.FileSize,
		BucketID:    "public-bucket",
	})
	if err != nil {
		return nil, err
	}

	res = &system.PreUploadFileRes{
		FileID:    preUploadRes.FileID,
		UploadUrl: preUploadRes.UploadURL,
		VisitLink: buildPreviewLink(preUploadRes.FileLink),
	}
	return
}

func (c *fileController) ReportUploadResult(ctx context.Context, req *system.ReportUploadResultReq) (res *system.ReportUploadResultRes, err error) {
	res = &system.ReportUploadResultRes{}
	if fileManager == nil {
		InitFileManager()
	}

	fileStatus := FileModule.FileStatusUploadSuccess
	if !req.Success {
		fileStatus = FileModule.FileStatusUploadFailed
	}
	err = fileManager.UpdateStatus(ctx, req.FileID, fileStatus)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *fileController) UploadFile(ctx context.Context, req *system.UploadFileReq) (res *system.UploadFileRes, err error) {
	if fileManager == nil {
		InitFileManager()
	}

	preUploadRes, err := fileManager.PreUpload(ctx, &FileModule.PreUploadReq{
		FileName:    req.File.Filename,
		ContentType: req.File.Header.Get("Content-Type"),
		Size:        req.File.Size,
		BucketID:    "public-bucket",
	})
	if err != nil {
		return nil, err
	}

	file, err := req.File.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	contentType := req.File.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	status, _, err := client.PUT(ctx, preUploadRes.UploadURL, map[string]interface{}{
		"Content-Type": contentType,
	}, fileBytes)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, errors.New("上传文件失败")
	}

	res = &system.UploadFileRes{
		FileID:   preUploadRes.FileID,
		FileLink: buildPreviewLink(preUploadRes.FileLink),
	}
	return res, nil
}

func buildPreviewLink(link string) string {
	if link == "" {
		return link
	}
	u, err := url.Parse(link)
	if err != nil {
		return link
	}
	q := u.Query()
	q.Set("response-content-disposition", "inline")
	u.RawQuery = q.Encode()
	return u.String()
}
