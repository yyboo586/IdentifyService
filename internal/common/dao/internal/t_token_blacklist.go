package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TokenBlacklistDao is the data access object for table t_token_blacklist.
type TokenBlacklistDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TokenBlacklistColumns // columns contains all the column names of Table for convenient usage.
}

// TokenBlacklistColumns defines and stores column names for table t_token_blacklist.
type TokenBlacklistColumns struct {
	ID         string // 令牌id
	OperatorID string // 操作者ID
	CreatedAt  string // 创建时间
}

// tokenBlacklistColumns holds the columns for table t_token_blacklist.
var tokenBlacklistColumns = TokenBlacklistColumns{
	ID:         "id",
	OperatorID: "operator_id",
	CreatedAt:  "created_at",
}

// NewTokenBlacklistDao creates and returns a new DAO object for table data access.
func NewTokenBlacklistDao() *TokenBlacklistDao {
	return &TokenBlacklistDao{
		group:   "default",
		table:   "t_token_blacklist",
		columns: tokenBlacklistColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TokenBlacklistDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TokenBlacklistDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TokenBlacklistDao) Columns() TokenBlacklistColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TokenBlacklistDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TokenBlacklistDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TokenBlacklistDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
