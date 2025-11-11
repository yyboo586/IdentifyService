package model

type ScopeAuthDataReq struct {
	MenuId  uint     `p:"menuId"`
	Scope   int      `p:"scope"`
	DeptIds []uint64 `p:"deptIds"`
}

type ScopeAuthData struct {
	Id        uint64  `json:"id"`
	RoleId    uint    `json:"roleId"`
	MenuId    uint    `json:"menuId"`
	DataScope int     `json:"dataScope"`
	DeptIds   []int64 `json:"deptIds"`
}
