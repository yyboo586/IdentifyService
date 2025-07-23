package service

import (
	"context"

	"IdentifyService/api/v1/system"
)

type ISysDictData interface {
	GetDictWithDataByType(ctx context.Context, dictType, defaultValue string) (dict *system.GetDictRes, err error)
	List(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error)
	Add(ctx context.Context, req *system.DictDataAddReq, userId uint64) (err error)
	Get(ctx context.Context, dictCode uint) (res *system.DictDataGetRes, err error)
	Edit(ctx context.Context, req *system.DictDataEditReq, userId uint64) (err error)
	Delete(ctx context.Context, ids []int) (err error)
}

var localSysDictData ISysDictData

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}
