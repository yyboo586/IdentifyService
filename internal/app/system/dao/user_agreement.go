package dao

import (
	"IdentifyService/internal/app/system/model/entity"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type userAgreementDao struct {
	table   string
	group   string
	columns userAgreementColumns
}

type userAgreementColumns struct {
	Id            string
	UserId        string
	AgreementID   string
	AgreementName string
	VersionCode   string
	Agreed        string
	CreatedAt     string
}

var UserAgreement = userAgreementDao{
	table: "t_user_agreement",
	group: "default",
	columns: userAgreementColumns{
		Id:            "id",
		UserId:        "user_id",
		AgreementID:   "agreement_id",
		AgreementName: "agreement_name",
		VersionCode:   "version_code",
		Agreed:        "agreed",
		CreatedAt:     "created_at",
	},
}

func (dao userAgreementDao) Columns() userAgreementColumns {
	return dao.columns
}

func (dao userAgreementDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB(dao.group).Model(dao.table).Ctx(ctx)
}

func (dao userAgreementDao) Insert(ctx context.Context, data g.Map) error {
	_, err := dao.Ctx(ctx).Data(data).Insert()
	return err
}

// GetUserAgreements 获取用户同意的协议列表
func (dao userAgreementDao) GetUserAgreements(ctx context.Context, userID string) ([]*entity.UserAgreement, error) {
	var list []*entity.UserAgreement
	err := dao.Ctx(ctx).
		Where(dao.columns.UserId, userID).
		Where(dao.columns.Agreed, 1).
		OrderDesc(dao.columns.CreatedAt).
		Scan(&list)
	return list, err
}
