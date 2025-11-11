package service

import (
	"context"

	"IdentifyService/internal/app/common/model"
)

type ISysAttachment interface {
	List(ctx context.Context, req *model.SysAttachmentSearchReq) (res *model.SysAttachmentSearchRes, err error)
	GetById(ctx context.Context, Id int64) (res *model.SysAttachmentInfoRes, err error)
	GetByMd5(ctx context.Context, md5 string) (res *model.SysAttachmentInfoRes, err error)
	AddUpload(ctx context.Context, req *model.UploadResponse, attr *model.SysAttachmentAddAttribute) (err error)
	Add(ctx context.Context, req *model.SysAttachmentAddReq) (err error)
	Edit(ctx context.Context, req *model.SysAttachmentEditReq) (err error)
	Delete(ctx context.Context, Id []int64) (err error)
	// 附件管理状态修改（状态）
	ChangeStatus(ctx context.Context, id int64, status bool) (err error)
}

var localSysAttachment ISysAttachment

func SysAttachment() ISysAttachment {
	if localSysAttachment == nil {
		panic("implement not found for interface ISysAttachment, forgot register?")
	}
	return localSysAttachment
}

func RegisterSysAttachment(i ISysAttachment) {
	localSysAttachment = i
}
