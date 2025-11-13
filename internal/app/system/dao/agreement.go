package dao

import (
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
	Version      string
	Content      string
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
		Version:      "version",
		Content:      "content",
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

func (dao agreementDao) List(ctx context.Context, name string, page, size int) ([]*entity.Agreement, int, error) {
	model := dao.Ctx(ctx)
	if name != "" {
		model = model.WhereLike(dao.columns.Name, name)
	}
	count, err := model.Clone().Count()
	if err != nil {
		return nil, 0, err
	}

	var list []*entity.Agreement
	err = model.
		OrderDesc(dao.columns.MajorVersion).
		OrderDesc(dao.columns.MinorVersion).
		OrderDesc(dao.columns.PatchVersion).
		Page(page, size).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// GetMaxMajorVersion 获取指定协议名称的最大主版本号
func (dao agreementDao) GetMaxMajorVersion(ctx context.Context, name string) (string, error) {
	var list []*entity.Agreement
	err := dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		OrderDesc(dao.columns.MajorVersion).
		OrderDesc(dao.columns.MinorVersion).
		OrderDesc(dao.columns.PatchVersion).
		Limit(1).
		Scan(&list)
	if err != nil {
		return "0", err
	}
	if len(list) == 0 {
		return "0", nil
	}
	return list[0].MajorVersion, nil
}

// GetMaxMinorVersion 获取指定协议名称和主版本号的最大次版本号
func (dao agreementDao) GetMaxMinorVersion(ctx context.Context, name, majorVersion string) (string, error) {
	var list []*entity.Agreement
	err := dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		Where(dao.columns.MajorVersion, majorVersion).
		OrderDesc(dao.columns.MinorVersion).
		OrderDesc(dao.columns.PatchVersion).
		Limit(1).
		Scan(&list)
	if err != nil {
		return "0", err
	}
	if len(list) == 0 {
		return "0", nil
	}
	return list[0].MinorVersion, nil
}

// GetMaxPatchVersion 获取指定协议名称、主版本号和次版本号的最大补丁版本号
func (dao agreementDao) GetMaxPatchVersion(ctx context.Context, name, majorVersion, minorVersion string) (string, error) {
	var list []*entity.Agreement
	err := dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		Where(dao.columns.MajorVersion, majorVersion).
		Where(dao.columns.MinorVersion, minorVersion).
		OrderDesc(dao.columns.PatchVersion).
		Limit(1).
		Scan(&list)
	if err != nil {
		return "0", err
	}
	if len(list) == 0 {
		return "0", nil
	}
	return list[0].PatchVersion, nil
}

// GetLatestAgreement 获取指定协议名称的最新版本
func (dao agreementDao) GetLatestAgreement(ctx context.Context, name string) (*entity.Agreement, error) {
	// 1. 获取最大主版本号
	maxMajor, err := dao.GetMaxMajorVersion(ctx, name)
	if err != nil {
		return nil, err
	}
	if maxMajor == "0" {
		return nil, gerror.New("协议不存在")
	}

	// 2. 获取该主版本下的最大次版本号
	maxMinor, err := dao.GetMaxMinorVersion(ctx, name, maxMajor)
	if err != nil {
		return nil, err
	}

	// 3. 获取该主版本+次版本下的最大补丁版本号
	maxPatch, err := dao.GetMaxPatchVersion(ctx, name, maxMajor, maxMinor)
	if err != nil {
		return nil, err
	}

	// 4. 查询该版本记录
	var agreement entity.Agreement
	err = dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		Where(dao.columns.MajorVersion, maxMajor).
		Where(dao.columns.MinorVersion, maxMinor).
		Where(dao.columns.PatchVersion, maxPatch).
		Scan(&agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}

// GetByMajorVersion 根据协议名称和主版本号获取协议
func (dao agreementDao) GetByMajorVersion(ctx context.Context, name, majorVersion string) (*entity.Agreement, error) {
	var agreement entity.Agreement
	err := dao.Ctx(ctx).
		Where(dao.columns.Name, name).
		Where(dao.columns.MajorVersion, majorVersion).
		Where(dao.columns.MinorVersion, "0").
		Where(dao.columns.PatchVersion, "0").
		Scan(&agreement)
	if err != nil {
		return nil, err
	}
	return &agreement, nil
}

func gerrorfNotFound(resource string, value interface{}) error {
	return gerror.Newf("%s not found: %v", resource, value)
}
