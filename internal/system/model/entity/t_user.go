package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table t_user.
type User struct {
	ID        string      `orm:"id"`
	Name      string      `orm:"name"`
	Nickname  string      `orm:"nickname"`
	Password  string      `orm:"password"`
	Salt      string      `orm:"salt"`
	Status    int         `orm:"status"`
	OrgID     string      `orm:"org_id"`
	Sex       int         `orm:"sex"`
	Email     string      `orm:"email"`
	Avatar    string      `orm:"avatar"`
	Mobile    string      `orm:"mobile"`
	Address   string      `orm:"address"`
	Describe  string      `orm:"describe"`
	IsAdmin   int         `orm:"is_admin"`
	CreatedAt *gtime.Time `orm:"created_at"`
	UpdatedAt *gtime.Time `orm:"updated_at"`
}
