package agreement

import (
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sAgreement struct{}

func init() {
	service.RegisterAgreement(New())
}

func New() service.IAgreement {
	return &sAgreement{}
}

func (s *sAgreement) CreateAgreement(ctx context.Context, input *model.AgreementCreateInput) (int64, error) {
	if input == nil {
		return 0, gerror.New("协议信息不能为空")
	}
	if !model.IsValidAgreementName(input.Name) {
		return 0, gerror.New("协议名称不正确")
	}
	if input.Major < 0 || input.Minor < 0 || input.Patch < 0 {
		return 0, gerror.New("版本号不能为负数")
	}

	exists, err := dao.Agreement.ExistsVersion(ctx, input.Name, input.Major, input.Minor, input.Patch)
	if err != nil {
		return 0, gerror.Wrap(err, "校验协议版本失败")
	}
	if exists {
		return 0, gerror.New("协议版本已存在")
	}

	now := gtime.Now().Unix()
	data := g.Map{
		dao.Agreement.Columns().Name:         input.Name,
		dao.Agreement.Columns().MajorVersion: input.Major,
		dao.Agreement.Columns().MinorVersion: input.Minor,
		dao.Agreement.Columns().PatchVersion: input.Patch,
		dao.Agreement.Columns().VersionCode:  model.BuildVersionCode(input.Major, input.Minor, input.Patch),
		dao.Agreement.Columns().Status:       model.AgreementStatusDraft,
		dao.Agreement.Columns().Content:      input.Content,
		dao.Agreement.Columns().CreatedAt:    now,
		dao.Agreement.Columns().UpdatedAt:    now,
	}
	if input.PublishNow {
		data[dao.Agreement.Columns().Status] = model.AgreementStatusPublished
		data[dao.Agreement.Columns().PublishedAt] = now
	}

	id, err := dao.Agreement.Insert(ctx, data)
	if err != nil {
		return 0, gerror.Wrap(err, "新增协议失败")
	}
	return id, nil
}

func (s *sAgreement) UpdateAgreement(ctx context.Context, input *model.AgreementUpdateInput) (int64, error) {
	if input == nil {
		return 0, gerror.New("协议信息不能为空")
	}
	if input.ID == 0 {
		return 0, gerror.New("协议ID不能为空")
	}

	oldAgreement, err := dao.Agreement.GetByID(ctx, input.ID)
	if err != nil {
		return 0, gerror.Wrap(err, "获取协议信息失败")
	}

	now := gtime.Now().Unix()
	data := g.Map{
		dao.Agreement.Columns().Content:   input.Content,
		dao.Agreement.Columns().UpdatedAt: now,
	}

	if input.PublishNow && oldAgreement.Status != model.AgreementStatusPublished {
		data[dao.Agreement.Columns().Status] = model.AgreementStatusPublished
		data[dao.Agreement.Columns().PublishedAt] = now
	}
	if input.Status != nil {
		data[dao.Agreement.Columns().Status] = *input.Status
		if *input.Status == model.AgreementStatusPublished {
			data[dao.Agreement.Columns().PublishedAt] = now
		} else {
			data[dao.Agreement.Columns().PublishedAt] = 0
		}
	}

	if err := dao.Agreement.Update(ctx, input.ID, data); err != nil {
		return 0, gerror.Wrap(err, "更新协议失败")
	}
	return input.ID, nil
}

func (s *sAgreement) DeleteAgreement(ctx context.Context, id int64) error {
	if id == 0 {
		return gerror.New("协议ID不能为空")
	}
	if err := dao.Agreement.Delete(ctx, id); err != nil {
		return gerror.Wrap(err, "删除协议失败")
	}
	return nil
}

func (s *sAgreement) GetAgreement(ctx context.Context, id int64) (*model.Agreement, error) {
	if id == 0 {
		return nil, gerror.New("协议ID不能为空")
	}
	item, err := dao.Agreement.GetByID(ctx, id)
	if err != nil {
		return nil, gerror.Wrap(err, "获取协议失败")
	}
	return convertAgreementEntity(item), nil
}

func (s *sAgreement) ListAgreements(ctx context.Context, input *model.AgreementListInput) (out []*model.Agreement, pageRes *model.PageRes, err error) {
	if input == nil {
		input = &model.AgreementListInput{}
	}

	if input.Page <= 0 {
		input.Page = 1
	}
	if input.Size <= 0 {
		input.Size = 10
	}

	entities, total, err := dao.Agreement.List(ctx, input.Name, input.Status, input.Page, input.Size)
	if err != nil {
		return nil, nil, gerror.Wrap(err, "获取协议列表失败")
	}

	for _, entity := range entities {
		out = append(out, convertAgreementEntity(entity))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: input.Page,
	}
	return out, pageRes, nil
}

func (s *sAgreement) GetLatestAgreement(ctx context.Context, name string) (*model.Agreement, error) {
	if !model.IsValidAgreementName(name) {
		return nil, gerror.New("协议名称不正确")
	}

	entity, err := dao.Agreement.GetLatestAgreement(ctx, name)
	if err != nil {
		return nil, gerror.Wrap(err, "获取最新协议失败")
	}

	return convertAgreementEntity(entity), nil
}

func (s *sAgreement) RecordUserAgreements(ctx context.Context, info *model.UserAgreement) error {
	if info == nil {
		return gerror.New("用户协议信息不能为空")
	}
	agreement, err := dao.Agreement.GetByID(ctx, info.AgreementID)
	if err != nil {
		return gerror.Wrap(err, "获取协议版本失败")
	}
	data := g.Map{
		dao.UserAgreement.Columns().UserId:        info.UserID,
		dao.UserAgreement.Columns().AgreementID:   info.AgreementID,
		dao.UserAgreement.Columns().AgreementName: info.AgreementName,
		dao.UserAgreement.Columns().VersionCode:   agreement.VersionCode,
		dao.UserAgreement.Columns().CreatedAt:     time.Now().Unix(),
	}
	if info.Agreed {
		data[dao.UserAgreement.Columns().Agreed] = 1
	} else {
		data[dao.UserAgreement.Columns().Agreed] = 0
	}
	if err := dao.UserAgreement.Insert(ctx, data); err != nil {
		return gerror.Wrap(err, "记录用户协议失败")
	}
	return nil
}

func (s *sAgreement) GetUserAgreements(ctx context.Context, userID string) ([]*model.UserAgreement, error) {
	if userID == "" {
		return nil, gerror.New("用户ID不能为空")
	}

	entities, err := dao.UserAgreement.GetUserAgreements(ctx, userID)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户协议列表失败")
	}

	var result []*model.UserAgreement
	for _, entity := range entities {
		result = append(result, &model.UserAgreement{
			ID:            entity.Id,
			UserID:        entity.UserID,
			AgreementID:   entity.AgreementID,
			AgreementName: entity.AgreementName,
			VersionCode:   entity.VersionCode,
			Agreed:        entity.Agreed == 1,
			CreatedAt:     gtime.New(time.Unix(entity.CreatedAt, 0)),
		})
	}
	return result, nil
}

func convertAgreementEntity(item *entity.Agreement) *model.Agreement {
	if item == nil {
		return nil
	}
	return &model.Agreement{
		ID:           item.Id,
		Name:         item.Name,
		MajorVersion: item.MajorVersion,
		MinorVersion: item.MinorVersion,
		PatchVersion: item.PatchVersion,
		VersionCode:  item.VersionCode,
		Status:       item.Status,
		Content:      item.Content,
		PublishedAt:  gtime.New(item.PublishedAt),
		CreatedAt:    gtime.New(item.CreatedAt),
		UpdatedAt:    gtime.New(item.UpdatedAt),
	}
}
