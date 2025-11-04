package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type TSmsCodeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TSmsCodeColumns // columns contains all the column names of Table for convenient usage.
}

type TSmsCodeColumns struct {
	Id           string // ID
	BusinessType string // Business Type
	Phone        string // Phone
	Code         string // Code
	Status       string // Status
	CreatedAt    string // Created At
	ExpiredAt    string // Expired At
	UpdatedAt    string // Updated At
}

var tSmsCodeColumns = TSmsCodeColumns{
	Id:           "id",
	BusinessType: "business_type",
	Phone:        "phone",
	Code:         "code",
	Status:       "status",
	CreatedAt:    "created_at",
	ExpiredAt:    "expired_at",
	UpdatedAt:    "updated_at",
}

func NewTSmsCodeDao() *TSmsCodeDao {
	return &TSmsCodeDao{
		group:   "default",
		table:   "t_sms_code",
		columns: tSmsCodeColumns,
	}
}

func (dao *TSmsCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *TSmsCodeDao) Table() string {
	return dao.table
}

func (dao *TSmsCodeDao) Columns() TSmsCodeColumns {
	return dao.columns
}

func (dao *TSmsCodeDao) Group() string {
	return dao.group
}

func (dao *TSmsCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *TSmsCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
