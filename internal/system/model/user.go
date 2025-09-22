package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

const (
	DefaultSuperAdminID = "11111111-0000-0000-0000-000000000000"
)

type UserStatus int

const (
	_                 UserStatus = iota
	UserStatusEnabled            // 正常
	UserStatusDisable            // 禁用
)

func IsValidUserStatus(status UserStatus) bool {
	return status == UserStatusDisable || status == UserStatusEnabled
}

type UserSex int

const (
	UserSexUnknown UserSex = iota // 保密
	UserSexMale                   // 男
	UserSexFemale                 // 女
)

// UserLoginRes 登录返回
type UserLoginRes struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Nickname string  `json:"nick_name"`
	IsAdmin  bool    `json:"is_admin"`
	Avatar   string  `json:"avatar"`
	OrgID    string  `json:"org_id"`
	RoleIDs  []int64 `json:"role_ids"`
}

type LinkUserRes struct {
	gmeta.Meta `orm:"table:t_user"`
	ID         uint64 `orm:"id"       json:"id"`
	Nickname   string `orm:"nickname"    json:"nick_name"`
}

type User struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Nickname  string             `json:"nick_name"`
	Password  string             `json:"password"`
	Salt      string             `json:"salt"`
	Status    UserStatus         `json:"status"`
	OrgID     string             `json:"org_id"`
	Roles     []map[int64]string `json:"roles"`
	Sex       UserSex            `json:"sex"`
	Email     string             `json:"email"`
	Avatar    string             `json:"avatar"`
	Mobile    string             `json:"mobile"`
	Address   string             `json:"address"`
	Describe  string             `json:"describe"`
	IsAdmin   bool               `json:"is_admin"`
	CreatedAt *gtime.Time        `json:"created_at"`
	UpdatedAt *gtime.Time        `json:"updated_at"`
}

// 前端资源（用户管理）
const (
	DirUserManager  = "dir:user:manage"
	MenuUserManager = "menu:user:list"

	ButtonUserCreate = "button:user:create"
	ButtonUserDelete = "button:user:delete"
	ButtonUserEdit   = "button:user:edit"
	ButtonUserView   = "button:user:view"
)

// API 资源（用户管理）
const (
	APIUserCreate = "api:user:create"
	APIUserDelete = "api:user:delete"
	APIUserEdit   = "api:user:edit"
	APIUserView   = "api:user:view"
	APIUserList   = "api:user:list"
)

var UserResources []*Resource = []*Resource{
	{Type: ResourceTypeDir, Code: DirUserManager},
	{Type: ResourceTypeMenu, Code: MenuUserManager},
	{Type: ResourceTypeButton, Code: ButtonUserCreate},
	{Type: ResourceTypeButton, Code: ButtonUserDelete},
	{Type: ResourceTypeButton, Code: ButtonUserEdit},
	{Type: ResourceTypeButton, Code: ButtonUserView},
	{Type: ResourceTypeAPI, Code: APIUserCreate},
	{Type: ResourceTypeAPI, Code: APIUserDelete},
	{Type: ResourceTypeAPI, Code: APIUserEdit},
	{Type: ResourceTypeAPI, Code: APIUserView},
	{Type: ResourceTypeAPI, Code: APIUserList},
}

// 用户管理权限点常量
const (
	UserCreate PermissionPointCode = "user:create"
	UserDelete PermissionPointCode = "user:delete"
	UserEdit   PermissionPointCode = "user:edit"
	UserView   PermissionPointCode = "user:view"
	UserList   PermissionPointCode = "user:list"
)

var UserPermissions []*PermissionPoint = []*PermissionPoint{
	{
		Code:     UserCreate,
		CodeName: "创建用户",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirUserManager},
			{Type: ResourceTypeMenu, Code: MenuUserManager},
			{Type: ResourceTypeAPI, Code: APIUserList},
			{Type: ResourceTypeButton, Code: ButtonUserCreate},
			{Type: ResourceTypeAPI, Code: APIUserCreate},
		},
	},
	{
		Code:     UserDelete,
		CodeName: "删除用户",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirUserManager},
			{Type: ResourceTypeMenu, Code: MenuUserManager},
			{Type: ResourceTypeAPI, Code: APIUserList},
			{Type: ResourceTypeButton, Code: ButtonUserDelete},
			{Type: ResourceTypeAPI, Code: APIUserDelete},
		},
	},
	{
		Code:     UserEdit,
		CodeName: "编辑用户",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirUserManager},
			{Type: ResourceTypeMenu, Code: MenuUserManager},
			{Type: ResourceTypeAPI, Code: APIUserList},
			{Type: ResourceTypeButton, Code: ButtonUserEdit},
			{Type: ResourceTypeAPI, Code: APIUserEdit},
		},
	},
	{
		Code:     UserView,
		CodeName: "用户详情",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirUserManager},
			{Type: ResourceTypeMenu, Code: MenuUserManager},
			{Type: ResourceTypeAPI, Code: APIUserList},
			{Type: ResourceTypeButton, Code: ButtonUserView},
			{Type: ResourceTypeAPI, Code: APIUserView},
		},
	},
	{
		Code:     UserList,
		CodeName: "用户列表",
		Resources: []*Resource{
			{Type: ResourceTypeDir, Code: DirUserManager},
			{Type: ResourceTypeMenu, Code: MenuUserManager},
			{Type: ResourceTypeAPI, Code: APIUserList},
		},
	},
}
