package dao

import (
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type agreementDao struct {
	table   string
	group   string
	columns agreementColumns
}

type agreementColumns struct {
	Id           string
	Name         string
	MajorVersion string
	MinorVersion string
	PatchVersion string
	VersionCode  string
	Status       string
	Content      string
	PublishedAt  string
	CreatedAt    string
	UpdatedAt    string
}

var Agreement = agreementDao{
	table: "t_agreement",
	group: "default",
	columns: agreementColumns{
		Id:           "id",
		Name:         "name",
		MajorVersion: "major_version",
		MinorVersion: "minor_version",
		PatchVersion: "patch_version",
		VersionCode:  "version_code",
		Status:       "status",
		Content:      "content",
		PublishedAt:  "published_at",
		CreatedAt:    "created_at",
		UpdatedAt:    "updated_at",
	},
}

func (dao agreementDao) Columns() agreementColumns {
	return dao.columns
}

func (dao agreementDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB(dao.group).Model(dao.table).Ctx(ctx)
}

func (dao agreementDao) Insert(ctx context.Context, data g.Map) (int64, error) {
	result, err := dao.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dao agreementDao) Update(ctx context.Context, id int64, data g.Map) error {
	_, err := dao.Ctx(ctx).Where(dao.columns.Id, id).Data(data).Update()
	return err
}

func (dao agreementDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Ctx(ctx).Where(dao.columns.Id, id).Delete()
	return err
}

func (dao agreementDao) GetByID(ctx context.Context, id int64) (*entity.Agreement, error) {
	var agreement entity.Agreement
	err := dao.Ctx(ctx).Where(dao.columns.Id, id).Scan(&agreement)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.New("协议不存在")
		}
		return nil, err
	}
	return &agreement, nil
}

func (dao agreementDao) List(ctx context.Context, name string, status *int, page, size int) ([]*entity.Agreement, int, error) {
	model := dao.Ctx(ctx)
	if name != "" {
		model = model.WhereLike(dao.columns.Name, name)
	}
	if status != nil {
		model = model.Where(dao.columns.Status, *status)
	}
	count, err := model.Clone().Count()
	if err != nil {
		return nil, 0, err
	}

	var list []*entity.Agreement
	err = model.
		OrderDesc(dao.columns.VersionCode).
		Page(page, size).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// GetLatestAgreement 获取指定协议名称的最新版本
func (dao agreementDao) GetLatestAgreement(ctx context.Context, name string) (*entity.Agreement, error) {
	var agreement entity.Agreement
	err := dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		Where(dao.columns.Status, model.AgreementStatusPublished).
		OrderDesc(dao.columns.VersionCode).
		Limit(1).
		Scan(&agreement)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.New("协议不存在")
		}
		return nil, err
	}
	if agreement.Id == 0 {
		return nil, gerror.New("协议不存在")
	}

	return &agreement, nil
}

func (dao agreementDao) ExistsVersion(ctx context.Context, name string, major, minor, patch int) (bool, error) {
	count, err := dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		Where(dao.columns.MajorVersion, major).
		Where(dao.columns.MinorVersion, minor).
		Where(dao.columns.PatchVersion, patch).
		Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
