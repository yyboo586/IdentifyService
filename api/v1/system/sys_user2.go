package system

import "github.com/gogf/gf/v2/frame/g"

type UpdateUserTypeReq struct {
	g.Meta   `path:"/user/updateUserType" tags:"系统后台/用户管理" method:"patch" summary:"更新用户类型"`
	UserID   int64  `p:"userID" v:"required#用户ID不能为空"`
	UserType string `p:"userType" v:"required#用户类型不能为空"`
}

type UpdateUserTypeRes struct {
	g.Meta `mime:"application/json"`
}
