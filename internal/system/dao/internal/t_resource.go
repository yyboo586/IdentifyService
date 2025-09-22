package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResourceDao is the data access object for table t_resource.
type ResourceDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ResourceColumns // columns contains all the column names of Table for convenient usage.
}

// ResourceColumns defines and stores column names for table t_resource.
type ResourceColumns struct {
	ID        string // 资源ID
	Type      string // 资源类型
	Code      string // 资源代码
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// resourceColumns holds the columns for table t_resource.
var resourceColumns = ResourceColumns{
	ID:        "id",
	Type:      "type",
	Code:      "code",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewResourceDao creates and returns a new DAO object for table data access.
func NewResourceDao() *ResourceDao {
	return &ResourceDao{
		group:   "default",
		table:   "t_resource",
		columns: resourceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ResourceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ResourceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ResourceDao) Columns() ResourceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ResourceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ResourceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *ResourceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
