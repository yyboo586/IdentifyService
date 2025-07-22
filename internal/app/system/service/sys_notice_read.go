package service

import (
	"context"

	"IdentifyService/internal/app/system/model"
)

type ISysNoticeRead interface {
	List(ctx context.Context, req *model.SysNoticeReadSearchReq) (res *model.SysNoticeReadSearchRes, err error)
	ReadNotice(ctx context.Context, req *model.SysNoticeReadNoticeReq) (err error)
	GetById(ctx context.Context, Id int64) (res *model.SysNoticeReadInfoRes, err error)
	Add(ctx context.Context, req *model.SysNoticeReadAddReq) (err error)
	Edit(ctx context.Context, req *model.SysNoticeReadEditReq) (err error)
	Delete(ctx context.Context, Id []int64) (err error)
}

var localSysNoticeRead ISysNoticeRead

func SysNoticeRead() ISysNoticeRead {
	if localSysNoticeRead == nil {
		panic("implement not found for interface ISysNoticeRead, forgot register?")
	}
	return localSysNoticeRead
}

func RegisterSysNoticeRead(i ISysNoticeRead) {
	localSysNoticeRead = i
}
