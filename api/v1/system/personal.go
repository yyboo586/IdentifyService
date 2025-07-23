package system

import (
	v1 "IdentifyService/api/v1"
	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PersonalInfoReq struct {
	g.Meta `path:"/personal/getPersonalInfo" tags:"系统后台/用户管理" method:"get" summary:"登录用户信息"`
	v1.Author
}

type PersonalInfoRes struct {
	g.Meta   `mime:"application/json"`
	User     *entity.SysUser `json:"user"`
	Roles    []string        `json:"roles"`
	DeptName string          `json:"deptName"`
}

// SetPersonalReq 添加修改用户公用请求字段
type SetPersonalReq struct {
	Nickname  string `p:"nickname" v:"required#用户昵称不能为空"`
	Mobile    string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Remark    string `p:"remark"`
	Sex       int    `p:"sex"`
	UserEmail string `p:"userEmail" v:"required|email#邮箱不能为空|邮箱格式错误"`
	Describe  string `p:"describe"` //签名
	Avatar    string `p:"avatar"`   //签名

}

// PersonalEditReq 修改个人
type PersonalEditReq struct {
	g.Meta `path:"/personal/edit" tags:"系统后台/用户管理" method:"put" summary:"修改个人资料"`
	*SetPersonalReq
	v1.Author
}

type PersonalEditRes struct {
	v1.EmptyRes
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type PersonalResetPwdReq struct {
	g.Meta   `path:"/personal/resetPwd" tags:"系统后台/用户管理" method:"put" summary:"重置个人密码"`
	Password string `p:"password" v:"required|password2#密码不能为空|密码必须包含大小写字母和数字，长度在6~18之间"`
	v1.Author
}

type PersonalResetPwdRes struct {
}

type RefreshTokenReq struct {
	g.Meta `path:"/personal/refreshToken" tags:"系统后台/用户管理" method:"get" summary:"刷新token"`
	v1.Author
}

type RefreshTokenRes struct {
	v1.EmptyRes
	Token    string              `json:"token"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
}
