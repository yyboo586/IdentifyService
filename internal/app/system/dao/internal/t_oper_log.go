package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OperLogDao is the data access object for table sys_oper_log.
type OperLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns OperLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysOperLogColumns defines and stores column names for table sys_oper_log.
type OperLogColumns struct {
	ID         string // 日志主键
	OrgID      string // 组织ID
	OperName   string // 操作人员
	OperUrl    string // 请求URL
	OperMethod string // 操作方法
	OperIP     string // 主机地址
	OperTime   string // 操作时间
	CreatedAt  string // 创建时间
}

// operLogColumns holds the columns for table sys_oper_log.
var operLogColumns = OperLogColumns{
	ID:         "id",
	OrgID:      "org_id",
	OperName:   "oper_name",
	OperUrl:    "oper_url",
	OperMethod: "oper_method",
	OperIP:     "oper_ip",
	OperTime:   "oper_time",
	CreatedAt:  "created_at",
}

// NewOperLogDao creates and returns a new DAO object for table data access.
func NewOperLogDao() *OperLogDao {
	return &OperLogDao{
		group:   "default",
		table:   "t_oper_log",
		columns: operLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OperLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OperLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OperLogDao) Columns() OperLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OperLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OperLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OperLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
