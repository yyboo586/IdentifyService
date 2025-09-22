package model

const (
	// CtxKey 上下文变量存储键名，前后端系统共享
	CtxKey = "IdentifyServiceContext"
)

type ContextUser struct {
	UserID       string
	UserName     string
	UserNickname string
	OrgID        string
	RoleIDs      []int64
}
