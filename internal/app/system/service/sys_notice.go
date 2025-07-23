package service

import (
	"IdentifyService/internal/app/system/model"
	"context"
)

type ISysNotice interface {
	List(ctx context.Context, req *model.SysNoticeSearchReq) (res *model.SysNoticeSearchRes, err error)
	ListShow(ctx context.Context, req *model.SysNoticeSearchReq) (res *model.SysNoticeSearchRes, err error)

	GetById(ctx context.Context, ID int64) (res *model.SysNoticeInfoRes, err error)
	Add(ctx context.Context, req *model.SysNoticeAddReq) (err error)
	Edit(ctx context.Context, req *model.SysNoticeEditReq) (err error)
	Delete(ctx context.Context, ID []int64) (err error)
	//IndexData(ctx context.Context) (res *model.SysNoticeIndexRes, err error)
	/*首页相关数据接口*/
	UnReadList(ctx context.Context) (res *model.SysNoticeListRes, err error)
	UnReadCount(ctx context.Context, currentUser string) (sysNoticeUnreadIds *model.SysNoticeUnreadCount, err error)
	ReadAll(ctx context.Context, nType string) (err error)
	GetUserNameList(ctx context.Context, search string) (res []*model.SysNoticeUserNickname, err error)
	//NoticeReadLengthAdd(ctx context.Context, id int64) (err error)
	NoticeReadAddUserId(ctx context.Context, req *model.SysNoticeReadAddUserReq) (err error)
	//获取有我的消息的所有私信
	//CurrentUseWithIds(ctx context.Context, currentUserId uint64) (ids []int64, err error)
}

var localSysNotice ISysNotice

func SysNotice() ISysNotice {
	if localSysNotice == nil {
		panic("implement not found for interface ISysNotice, forgot register?")
	}
	return localSysNotice
}

func RegisterSysNotice(i ISysNotice) {
	localSysNotice = i
}
