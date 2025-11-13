package agreement

import (
	"IdentifyService/internal/app/system/dao"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"IdentifyService/internal/app/system/service"
	"context"
	"fmt"
	"strconv"
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
	if input.MajorVersion == "" {
		return 0, gerror.New("主版本号不能为空")
	}

	// 检查主版本号不能小于之前最大的主版本号
	maxMajor, err := dao.Agreement.GetMaxMajorVersion(ctx, input.Name)
	if err != nil {
		return 0, gerror.Wrap(err, "获取最大主版本号失败")
	}
	if maxMajor != "0" {
		compare := model.CompareVersion(input.MajorVersion, maxMajor)
		if compare < 0 {
			return 0, gerror.Newf("新主版本号(%s)不能小于当前最大主版本号(%s)", input.MajorVersion, maxMajor)
		}
		if compare == 0 {
			return 0, gerror.Newf("主版本号(%s)已存在", input.MajorVersion)
		}
	}

	now := gtime.Now().Unix()
	data := g.Map{
		dao.Agreement.Columns().Name:         input.Name,
		dao.Agreement.Columns().MajorVersion: input.MajorVersion,
		dao.Agreement.Columns().MinorVersion: "",
		dao.Agreement.Columns().PatchVersion: "",
		dao.Agreement.Columns().Version:      0,
		dao.Agreement.Columns().Content:      input.Content,
		dao.Agreement.Columns().CreatedAt:    now,
		dao.Agreement.Columns().UpdatedAt:    now,
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
	if !model.IsValidAgreementName(input.Name) {
		return 0, gerror.New("协议名称不正确")
	}
	if input.UpdateType != "minor" && input.UpdateType != "patch" {
		return 0, gerror.New("更新类型必须是 minor 或 patch")
	}

	oldAgreement, err := dao.Agreement.GetByID(ctx, input.ID)
	if err != nil {
		return 0, gerror.Wrap(err, "获取协议信息失败")
	}

	var newMajor, newMinor, newPatch string

	if input.UpdateType == "minor" {
		// 修改次版本：找到该主版本下的最大次版本，然后+1，补丁版本设为0
		maxMinor, err := dao.Agreement.GetMaxMinorVersion(ctx, input.Name, oldAgreement.MajorVersion)
		if err != nil {
			return 0, gerror.Wrap(err, "获取最大次版本号失败")
		}
		maxMinorInt, _ := strconv.Atoi(maxMinor)
		newMinorInt := maxMinorInt + 1
		newMajor = oldAgreement.MajorVersion
		newMinor = fmt.Sprintf("%d", newMinorInt)
		newPatch = "0"
	} else if input.UpdateType == "patch" {
		// 修改补丁版本：找到该主版本+次版本下的最大补丁版本，然后+1
		maxPatch, err := dao.Agreement.GetMaxPatchVersion(ctx, input.Name, oldAgreement.MajorVersion, oldAgreement.MinorVersion)
		if err != nil {
			return 0, gerror.Wrap(err, "获取最大补丁版本号失败")
		}
		maxPatchInt, _ := strconv.Atoi(maxPatch)
		newPatchInt := maxPatchInt + 1
		newMajor = oldAgreement.MajorVersion
		newMinor = oldAgreement.MinorVersion
		newPatch = fmt.Sprintf("%d", newPatchInt)
	} else {
		return 0, gerror.New("更新类型必须是 minor 或 patch")
	}

	now := gtime.Now().Unix()
	data := g.Map{
		dao.Agreement.Columns().Name:         input.Name,
		dao.Agreement.Columns().MajorVersion: newMajor,
		dao.Agreement.Columns().MinorVersion: newMinor,
		dao.Agreement.Columns().PatchVersion: newPatch,
		dao.Agreement.Columns().Version:      0,
		dao.Agreement.Columns().Content:      input.Content,
		dao.Agreement.Columns().CreatedAt:    now,
		dao.Agreement.Columns().UpdatedAt:    now,
	}

	id, err := dao.Agreement.Insert(ctx, data)
	if err != nil {
		return 0, gerror.Wrap(err, "新增协议版本失败")
	}
	return id, nil
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

	entities, total, err := dao.Agreement.List(ctx, input.Name, input.Page, input.Size)
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
	data := g.Map{
		dao.UserAgreement.Columns().UserId:        info.UserID,
		dao.UserAgreement.Columns().AgreementID:   info.AgreementID,
		dao.UserAgreement.Columns().AgreementName: info.AgreementName,
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
		Content:      item.Content,
		CreatedAt:    gtime.New(item.CreatedAt),
		UpdatedAt:    gtime.New(item.UpdatedAt),
	}
}
