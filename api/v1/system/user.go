package system

import (
	"IdentifyService/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserCreateReq struct {
	g.Meta `path:"/user" tags:"账户管理" method:"post" summary:"账户创建"`
	model.Author
	UserName string  `json:"user_ame" v:"required#用户名不能为空" dc:"用户名"`
	NickName string  `json:"nick_name" v:"required#用户昵称不能为空" dc:"用户昵称"`
	Password string  `json:"password" v:"required#用户密码不能为空" dc:"用户密码(必须包含大小写字母和数字,长度在6~18之间)"`
	OrgID    string  `json:"org_id" dc:"用户所属组织ID(如不指定,则使用创建者的组织ID)"`
	Sex      int     `json:"sex" dc:"用户性别(0:保密,1:男,2:女)"`
	Email    string  `json:"email" dc:"用户邮箱"`
	Avatar   string  `json:"avatar" dc:"用户头像"`
	Mobile   string  `json:"mobile" dc:"用户手机号"`
	Address  string  `json:"address" dc:"用户地址"`
	Describe string  `json:"describe" dc:"用户描述"`
	IsAdmin  int     `json:"is_admin" dc:"是否是后台管理员(0:否,1:是)"`
	RoleIDs  []int64 `json:"role_ids" dc:"用户关联角色ID列表"`
}

type UserCreateRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id" dc:"用户ID"`
}

type UserRegisterReq struct {
	g.Meta   `path:"/user/register" tags:"账户管理" method:"post" summary:"账户注册" dc:"适用于前台用户注册"`
	UserName string `json:"user_name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#用户密码不能为空" dc:"用户密码(必须包含大小写字母和数字,长度在6~18之间)"`
}

type UserRegisterRes struct {
	g.Meta `mime:"application/json"`
	ID     string `json:"id" dc:"用户ID"`
}

type UserDeleteReq struct {
	g.Meta `path:"/user" tags:"账户管理" method:"delete" summary:"删除用户"`
	model.Author
	IDs []string `json:"ids"  v:"required#ids不能为空" dc:"用户ID列表"`
}

type UserDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type EditUserPersonalInfoReq struct {
	g.Meta `path:"/user/{id}/personalInfo" tags:"账户管理" method:"put" summary:"修改个人资料(全量更新)"`
	model.Author
	ID       string `p:"id" v:"required#用户ID不能为空" dc:"用户ID"`
	NickName string `json:"nick_name" v:"required#用户昵称不能为空" dc:"用户昵称"`
	Email    string `json:"email" dc:"用户邮箱"`
	Avatar   string `json:"avatar" dc:"用户头像"`
	Mobile   string `json:"mobile" dc:"用户手机号"`
	Address  string `json:"address" dc:"用户地址"`
	Describe string `json:"describe" dc:"用户描述"`
}

type EditUserPersonalInfoRes struct {
	g.Meta `mime:"application/json"`
}

type EditUserRolesReq struct {
	g.Meta `path:"/user/{id}/roles" tags:"账户管理" method:"put" summary:"修改用户角色(全量更新)"`
	model.Author
	ID      string  `p:"id" v:"required#用户ID不能为空" dc:"用户ID"`
	RoleIDs []int64 `json:"role_ids" dc:"用户关联角色ID列表"`
}

type EditUserRolesRes struct {
	g.Meta `mime:"application/json"`
}

type EditUserStatusReq struct {
	g.Meta `path:"/user/{id}/status" tags:"账户管理" method:"put" summary:"修改用户状态"`
	model.Author
	ID      string `p:"id" v:"required#用户ID不能为空" dc:"用户ID"`
	Enabled bool   `json:"enabled" dc:"用户是否启用(true:启用,false:禁用)"`
}

type EditUserStatusRes struct {
	g.Meta `mime:"application/json"`
}

type UserResetPwdReq struct {
	g.Meta `path:"/user/reset-credentials" tags:"账户管理" method:"put" summary:"重置用户密码"`
	model.Author
	ID string `p:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}

type UserResetPwdRes struct {
	g.Meta `mime:"application/json"`
}

type GetUserInfoReq struct {
	g.Meta `path:"/user/{id}" tags:"账户管理" method:"get" summary:"获取用户信息"`
	model.Author
	ID string `p:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}

type GetUserInfoRes struct {
	g.Meta `mime:"application/json"`
	*User
}

type UserListReq struct {
	g.Meta `path:"/user" tags:"账户管理" method:"get" summary:"用户列表"`
	model.Author
	OrgID string `p:"org_id" v:"required#组织ID不能为空" dc:"组织ID"`
	Name  string `p:"name" dc:"用户名称"`
	model.PageReq
}

type UserListRes struct {
	g.Meta `mime:"application/json"`
	List   []*User `json:"list" dc:"用户列表"`
	model.PageRes
}

type User struct {
	ID        string      `json:"id" dc:"用户ID"`
	Name      string      `json:"name" dc:"用户名称"`
	Nickname  string      `json:"nick_name" dc:"用户昵称"`
	Mobile    string      `json:"mobile" dc:"用户手机号"`
	Email     string      `json:"email" dc:"用户邮箱"`
	Enabled   bool        `json:"enabled" dc:"用户状态(true:启用,false:禁用)"`
	RoleIDs   []int64     `json:"role_ids" dc:"用户关联角色ID列表"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}
