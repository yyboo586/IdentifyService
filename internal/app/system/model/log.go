package model

import "github.com/yyboo586/common/LogModule"

const (
	_ LogModule.LogModule = iota
	LogModuleDept
	LogModuleUser
)

const (
	_ LogModule.LogAction = iota
	LogActionUserLogin
	LogActionUserLogout
	LogActionUserUnRegister
)

func GetLogActionName(action LogModule.LogAction) string {
	switch action {
	case LogActionUserLogin:
		return "用户登录"
	case LogActionUserLogout:
		return "用户退出"
	case LogActionUserUnRegister:
		return "用户注销"
	}
	return "未知操作"
}
