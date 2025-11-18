package controller

import (
	"context"

	"IdentifyService/api/v1/system"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
)

var Agreement = new(agreementController)

type agreementController struct{}

func (c *agreementController) Create(ctx context.Context, req *system.AgreementCreateReq) (res *system.AgreementCreateRes, err error) {
	major, minor, patch, err := model.ParseVersion(req.Version)
	if err != nil {
		return nil, err
	}
	input := &model.AgreementCreateInput{
		Name:       req.Name,
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		Content:    req.Content,
		PublishNow: req.PublishNow,
	}
	id, err := service.Agreement().CreateAgreement(ctx, input)
	if err != nil {
		return nil, err
	}
	return &system.AgreementCreateRes{ID: id}, nil
}

func (c *agreementController) Update(ctx context.Context, req *system.AgreementUpdateReq) (res *system.AgreementUpdateRes, err error) {
	input := &model.AgreementUpdateInput{
		ID:         req.ID,
		Content:    req.Content,
		PublishNow: req.PublishNow,
		Status:     req.Status,
	}
	id, err := service.Agreement().UpdateAgreement(ctx, input)
	if err != nil {
		return nil, err
	}
	return &system.AgreementUpdateRes{ID: id}, nil
}

func (c *agreementController) Delete(ctx context.Context, req *system.AgreementDeleteReq) (res *system.AgreementDeleteRes, err error) {
	if err := service.Agreement().DeleteAgreement(ctx, req.ID); err != nil {
		return nil, err
	}
	return &system.AgreementDeleteRes{}, nil
}

func (c *agreementController) Get(ctx context.Context, req *system.AgreementGetReq) (res *system.AgreementGetRes, err error) {
	item, err := service.Agreement().GetAgreement(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &system.AgreementGetRes{
		AgreementItem: convertAgreementToAPI(item),
	}, nil
}

func (c *agreementController) List(ctx context.Context, req *system.AgreementListReq) (res *system.AgreementListRes, err error) {
	input := &model.AgreementListInput{
		Name:   req.Name,
		Status: req.Status,
		PageReq: model.PageReq{
			Page: req.Page,
			Size: req.Size,
		},
	}
	result, pageRes, err := service.Agreement().ListAgreements(ctx, input)
	if err != nil {
		return nil, err
	}

	list := make([]*system.AgreementItem, 0, len(result))
	for _, item := range result {
		list = append(list, convertAgreementToAPI(item))
	}

	return &system.AgreementListRes{
		List:    list,
		PageRes: pageRes,
	}, nil
}

func (c *agreementController) GetLatest(ctx context.Context, req *system.AgreementGetLatestReq) (res *system.AgreementGetLatestRes, err error) {
	item, err := service.Agreement().GetLatestAgreement(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &system.AgreementGetLatestRes{
		AgreementItem: convertAgreementToAPI(item),
	}, nil
}

func (c *agreementController) GetUserAgreements(ctx context.Context, req *system.UserAgreementListReq) (res *system.UserAgreementListRes, err error) {
	list, err := service.Agreement().GetUserAgreements(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	items := make([]*system.UserAgreementItem, 0, len(list))
	for _, item := range list {
		items = append(items, &system.UserAgreementItem{
			ID:            item.ID,
			UserID:        item.UserID,
			AgreementID:   item.AgreementID,
			AgreementName: item.AgreementName,
			Agreed:        item.Agreed,
			CreatedAt:     item.CreatedAt.Unix(),
		})
	}

	return &system.UserAgreementListRes{
		List: items,
	}, nil
}

func convertAgreementToAPI(item *model.Agreement) *system.AgreementItem {
	if item == nil {
		return nil
	}
	var createdAt, updatedAt, publishedAt int64
	if item.CreatedAt != nil {
		createdAt = item.CreatedAt.Unix()
	}
	if item.UpdatedAt != nil {
		updatedAt = item.UpdatedAt.Unix()
	}
	if item.PublishedAt != nil {
		publishedAt = item.PublishedAt.Unix()
	}
	return &system.AgreementItem{
		ID:           item.ID,
		Name:         item.Name,
		MajorVersion: item.MajorVersion,
		MinorVersion: item.MinorVersion,
		PatchVersion: item.PatchVersion,
		VersionCode:  item.VersionCode,
		Status:       item.Status,
		Content:      item.Content,
		Version:      item.VersionString(),
		PublishedAt:  publishedAt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
