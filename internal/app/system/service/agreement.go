package service

import (
	"context"

	"IdentifyService/internal/app/system/model"
)

type (
	IAgreement interface {
		CreateAgreement(ctx context.Context, input *model.AgreementCreateInput) (int64, error)
		UpdateAgreement(ctx context.Context, input *model.AgreementUpdateInput) (int64, error)
		DeleteAgreement(ctx context.Context, id int64) error
		GetAgreement(ctx context.Context, id int64) (*model.Agreement, error)
		ListAgreements(ctx context.Context, input *model.AgreementListInput) ([]*model.Agreement, *model.PageRes, error)
		GetLatestAgreement(ctx context.Context, name string) (*model.Agreement, error)

		RecordUserAgreements(ctx context.Context, info *model.UserAgreement) error
		GetUserAgreements(ctx context.Context, userID string) ([]*model.UserAgreement, error)
	}
)

var (
	localAgreement IAgreement
)

func Agreement() IAgreement {
	if localAgreement == nil {
		panic("implement not found for interface IAgreement, forgot register?")
	}
	return localAgreement
}

func RegisterAgreement(srv IAgreement) {
	localAgreement = srv
}
