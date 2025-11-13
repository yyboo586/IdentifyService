package controller

import (
	"context"
	"sync"

	"IdentifyService/api/v1/system"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yyboo586/common/FileModule"
)

var once sync.Once
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
		VisitLink: preUploadRes.FileLink,
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
