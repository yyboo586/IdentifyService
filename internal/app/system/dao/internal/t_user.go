package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table t_user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table t_user.
type UserColumns struct {
	ID        string //
	Name      string // 用户名
	Nickname  string // 用户昵称
	Password  string // 登录密码;cmf_password加密
	Salt      string // 加密盐
	Status    string // 用户状态;0:禁用,1:正常,2:未验证
	OrgID     string // 组织id
	Sex       string // 性别;0:保密,1:男,2:女
	Email     string // 用户登录邮箱
	Avatar    string // 用户头像
	Mobile    string // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Address   string // 联系地址
	Describe  string // 描述信息
	IsAdmin   string // 是否后台管理员
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// userColumns holds the columns for table t_user.
var userColumns = UserColumns{
	ID:        "id",
	Name:      "name",
	Nickname:  "nickname",
	Password:  "password",
	Salt:      "salt",
	Status:    "status",
	OrgID:     "org_id",
	Sex:       "sex",
	Email:     "email",
	Avatar:    "avatar",
	Mobile:    "mobile",
	Address:   "address",
	Describe:  "describe",
	IsAdmin:   "is_admin",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "t_user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
