package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginLogDao is the data access object for table t_login_log.
type LoginLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysLoginLogColumns defines and stores column names for table t_login_log.
type LoginLogColumns struct {
	ID        string
	OrgID     string // 组织ID
	LoginName string // 登录账号
	IP        string // 登录IP地址
	Browser   string // 浏览器类型
	Status    string // 登录状态（0成功 1失败）
	Message   string // 提示消息
	LoginTime string // 登录时间
	CreatedAt string // 创建时间
}

// sysLoginLogColumns holds the columns for table t_login_log.
var sysLoginLogColumns = LoginLogColumns{
	ID:        "id",
	OrgID:     "org_id",
	LoginName: "login_name",
	IP:        "ip",
	Browser:   "browser",
	Status:    "status",
	Message:   "message",
	LoginTime: "login_time",
	CreatedAt: "created_at",
}

// NewSysLoginLogDao creates and returns a new DAO object for table data access.
func NewLoginLogDao() *LoginLogDao {
	return &LoginLogDao{
		group:   "default",
		table:   "t_login_log",
		columns: sysLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LoginLogDao) Columns() LoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
