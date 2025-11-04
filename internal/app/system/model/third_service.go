package model

type IUQTUserInfo struct {
	IUQTID   string `json:"iuqt_id"`
	UserName string `json:"user_name"`
}

type SettleInfo struct {
	UserID int64  `json:"user_id"`
	Status string `json:"status"` // 入驻状态(未入驻、审核中、已入驻)
	ID     string `json:"id"`     // 服务提供商ID/展商ID
}

type SPSettleCheckRes struct {
	HttpResponse
	Data SPSettleCheckResData `json:"data"`
}

type SPSettleCheckResData struct {
	ApplyStatus       bool   `json:"apply_status" dc:"是否已申请" v:"required#是否已申请不能为空"`
	ServiceProviderID string `json:"service_provider_id" dc:"服务提供商ID"`
	ReviewStatus      string `json:"review_status" dc:"审核状态(待审核、已审核、已驳回、已禁用、已注销)" v:"required#审核状态不能为空"`
}

type MerchantSettleCheckRes struct {
	HttpResponse
	Data MerchantSettleCheckResData `json:"data"`
}

type MerchantSettleCheckResData struct {
	ApplyStatus  bool   `json:"apply_status" dc:"是否已申请" v:"required#是否已申请不能为空"`
	MerchantID   string `json:"merchant_id" dc:"展商ID"`
	ReviewStatus string `json:"review_status" dc:"审核状态(待审核、已审核、已驳回、已禁用、已注销)" v:"required#审核状态不能为空"`
}

type HttpResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
