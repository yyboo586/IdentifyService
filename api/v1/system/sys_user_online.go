package system

import (
	v1 "IdentifyService/api/v1"
	"IdentifyService/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// SysUserOnlineSearchReq 列表搜索参数
type SysUserOnlineSearchReq struct {
	g.Meta   `path:"/online/list" tags:"系统后台/在线用户管理" method:"get" summary:"列表"`
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	v1.PageReq
	v1.Author
}

// SysUserOnlineSearchRes 列表结果
type SysUserOnlineSearchRes struct {
	g.Meta `mime:"application/json"`
	v1.ListRes
	List []*entity.SysUserOnline `json:"list"`
}

type SysUserOnlineForceLogoutReq struct {
	g.Meta `path:"/online/forceLogout" tags:"系统后台/在线用户管理" method:"delete" summary:"强制用户退出登录"`
	v1.Author
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type SysUserOnlineForceLogoutRes struct {
	v1.EmptyRes
}
