package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionResourceDao is the data access object for table t_permission_resource.
type PermissionResourceDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns PermissionResourceColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionResourceColumns defines and stores column names for table t_permission_resource.
type PermissionResourceColumns struct {
	ID           string // 权限点资源ID
	PermissionID string // 权限点ID
	ResourceID   string // 资源ID
	CreatedAt    string // 创建时间
}

// permissionResourceColumns holds the columns for table t_permission_resource.
var permissionResourceColumns = PermissionResourceColumns{
	ID:           "id",
	PermissionID: "permission_id",
	ResourceID:   "resource_id",
	CreatedAt:    "created_at",
}

// NewPermissionResourceDao creates and returns a new DAO object for table data access.
func NewPermissionResourceDao() *PermissionResourceDao {
	return &PermissionResourceDao{
		group:   "default",
		table:   "t_permission_resource",
		columns: permissionResourceColumns,
	}
}

func (dao *PermissionResourceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *PermissionResourceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *PermissionResourceDao) Table() string {
	return dao.table
}

func (dao *PermissionResourceDao) Columns() PermissionResourceColumns {
	return dao.columns
}

func (dao *PermissionResourceDao) Group() string {
	return dao.group
}

func (dao *PermissionResourceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
