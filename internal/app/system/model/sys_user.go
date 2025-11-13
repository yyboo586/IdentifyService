package model

import (
	"IdentifyService/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

// LoginUserRes 登录返回
type LoginUserRes struct {
	Id           string `orm:"id,primary"       json:"id"`           //
	UserName     string `orm:"user_name,unique" json:"userName"`     // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	Mobile       string `orm:"mobile" json:"mobile"`                 //手机号
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // 用户状态;0:禁用,1:正常,2:未验证
	Sex          int    `orm:"sex"              json:"sex"`          // 性别;0:保密,1:男,2:女
	Birthday     string `orm:"birthday"         json:"birthday"`     // 生日
	City         string `orm:"city"             json:"city"`         // 城市
	Avatar       string `orm:"avatar"           json:"avatar"`       // 头像

	UserPassword string   `orm:"user_password"    json:"-"`       // 登录密码;cmf_password加密
	UserSalt     string   `orm:"user_salt"        json:"-"`       // 加密盐
	IsAdmin      int      `orm:"is_admin"         json:"isAdmin"` // 是否后台管理员 1 是  0   否
	DeptId       uint64   `orm:"dept_id"       json:"deptId"`     //部门id
	IUQTID       string   `orm:"iuqt_id"       json:"iuqtId"`     // IUQT ID
	UserType     UserType `orm:"user_type"      json:"userType"`  // 用户类型(1:服务提供商,2:展商)
}

// SysUserRoleDeptRes 带有部门、角色、岗位信息的用户数据
type SysUserRoleDeptRes struct {
	*entity.SysUser
	Dept     *entity.SysDept       `json:"dept"`
	RoleInfo []*SysUserRoleInfoRes `json:"roleInfo"`
	Post     []*SysUserPostInfoRes `json:"post"`
}

type SysUserRoleInfoRes struct {
	RoleId uint   `json:"roleId"`
	Name   string `json:"name"`
}

type SysUserPostInfoRes struct {
	PostId   int64  `json:"postId"`
	PostName string `json:"postName"`
}

type SysUserSimpleRes struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`                   //
	Avatar       string `orm:"avatar" json:"avatar"`                 // 头像
	Sex          int    `orm:"sex" json:"sex"`                       // 性别
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
}

type LinkUserRes struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`
	UserNickname string `orm:"user_nickname"    json:"userNickname"`
}
