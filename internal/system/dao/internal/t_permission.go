package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionDao is the data access object for table t_permission.
type PermissionDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PermissionColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionColumns defines and stores column names for table t_permission.
type PermissionColumns struct {
	ID          string // 权限点ID
	Code        string // 权限点代码
	CodeName    string // 权限点名称
	ResourceID  string // 资源ID
	Description string // 权限描述
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// permissionColumns holds the columns for table t_permission.
var permissionColumns = PermissionColumns{
	ID:          "id",
	Code:        "code",
	CodeName:    "code_name",
	ResourceID:  "resource_id",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewPermissionDao creates and returns a new DAO object for table data access.
func NewPermissionDao() *PermissionDao {
	return &PermissionDao{
		group:   "default",
		table:   "t_permission",
		columns: permissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionDao) Columns() PermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *PermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
