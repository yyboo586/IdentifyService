package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type UserStatus int

const (
	_                 UserStatus = iota
	UserStatusEnabled            // 正常
	UserStatusDisable            // 禁用
)

type UserSex int

const (
	UserSexUnknown UserSex = iota // 保密
	UserSexMale                   // 男
	UserSexFemale                 // 女
)

// UserLoginRes 登录返回
type UserLoginRes struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nick_name"`
	Mobile   string `json:"mobile"`
	Status   int64  `json:"status"`
	IsAdmin  bool   `json:"is_admin"`
	Avatar   string `json:"avatar"`
	OrgID    string `json:"org_id"`
}

type LinkUserRes struct {
	gmeta.Meta `orm:"table:t_user"`
	ID         uint64 `orm:"id"       json:"id"`
	Nickname   string `orm:"nickname"    json:"nick_name"`
}

type User struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Nickname  string      `json:"nick_name"`
	Password  string      `json:"password"`
	Salt      string      `json:"salt"`
	Status    UserStatus  `json:"status"`
	OrgID     string      `json:"org_id"`
	Sex       UserSex     `json:"sex"`
	Email     string      `json:"email"`
	Avatar    string      `json:"avatar"`
	Mobile    string      `json:"mobile"`
	Address   string      `json:"address"`
	Describe  string      `json:"describe"`
	IsAdmin   bool        `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

func ConvertUserToLoginRes(in *User) *UserLoginRes {
	return &UserLoginRes{
		ID:       in.ID,
		Name:     in.Name,
		Nickname: in.Nickname,
		Mobile:   in.Mobile,
		Avatar:   in.Avatar,
		IsAdmin:  in.IsAdmin,
		OrgID:    in.OrgID,
		Status:   int64(in.Status),
	}
}
