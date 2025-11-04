package service

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/yyboo586/common/httpUtils"
)

type IThirdService interface {
	// IUQT
	LoginByTicket(ctx context.Context, ticket string) (iuqtUserInfo *model.IUQTUserInfo, err error)

	// ExhibitionService
	GetSettleInfo(ctx context.Context, userID int64, userType model.UserType) (settleInfo *model.SettleInfo, err error)
}

var localThirdService IThirdService

func ThirdService() IThirdService {
	if localThirdService == nil {
		panic("implement not found for interface IThirdService, forgot register?")
	}
	return localThirdService
}

func RegisterThirdService() {
	localThirdService = NewThirdService()
}

type thirdService struct {
	IUQTAddr     string // 服务地址
	ExEngineAddr string // 引擎服务地址

	client httpUtils.HTTPClient
}

func NewThirdService() IThirdService {
	return &thirdService{
		IUQTAddr:     "",
		ExEngineAddr: g.Cfg().MustGet(context.Background(), "server.service.exEngine.addr").String(),

		client: httpUtils.NewHTTPClientWithDebug(true),
	}
}

// TODO: 对接第三方登录接口
func (i *thirdService) LoginByTicket(ctx context.Context, ticket string) (iuqtUserInfo *model.IUQTUserInfo, err error) {
	return &model.IUQTUserInfo{
		IUQTID:   ticket,
		UserName: fmt.Sprintf("iuqt_%s", ticket),
	}, nil
}

func (i *thirdService) GetSettleInfo(ctx context.Context, userID int64, userType model.UserType) (settleInfo *model.SettleInfo, err error) {
	switch userType {
	case model.UserTypeServiceProvider:
		return i.getSPSettleInfo(ctx, userID)
	case model.UserTypeExhibitor:
		return i.getMerchantSettleInfo(ctx, userID)
	default:
		settleInfo = &model.SettleInfo{
			UserID: userID,
			Status: "未入驻",
		}
		return settleInfo, nil
	}
}

func (i *thirdService) getSPSettleInfo(ctx context.Context, userID int64) (settleInfo *model.SettleInfo, err error) {
	url := fmt.Sprintf("%s/api/v1/exhibition-service/service-provider/check", i.ExEngineAddr)
	body := map[string]interface{}{
		"user_id": userID,
	}
	status, resBody, err := i.client.POST(ctx, url, nil, body)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, gerror.New(string(resBody))
	}

	var settleCheckRes model.SPSettleCheckRes
	err = gconv.Struct(resBody, &settleCheckRes)
	if err != nil {
		return nil, err
	}

	settleInfo = &model.SettleInfo{
		UserID: userID,
	}
	if !settleCheckRes.Data.ApplyStatus {
		settleInfo.Status = "未入驻"
		return
	}
	settleInfo.ID = settleCheckRes.Data.ServiceProviderID
	settleInfo.Status = settleCheckRes.Data.ReviewStatus
	return
}

func (i *thirdService) getMerchantSettleInfo(ctx context.Context, userID int64) (settleInfo *model.SettleInfo, err error) {
	url := fmt.Sprintf("%s/api/v1/exhibition-service/merchant/check", i.ExEngineAddr)
	body := map[string]interface{}{
		"user_id": userID,
	}
	status, resBody, err := i.client.POST(ctx, url, nil, body)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, gerror.New(string(resBody))
	}

	var settleCheckRes model.MerchantSettleCheckRes
	err = gconv.Struct(resBody, &settleCheckRes)
	if err != nil {
		return nil, err
	}

	settleInfo = &model.SettleInfo{
		UserID: userID,
	}
	if !settleCheckRes.Data.ApplyStatus {
		settleInfo.Status = "未入驻"
		return
	}
	settleInfo.ID = settleCheckRes.Data.MerchantID
	settleInfo.Status = settleCheckRes.Data.ReviewStatus
	return
}
