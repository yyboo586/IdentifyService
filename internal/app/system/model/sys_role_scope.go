package model

type ScopeAuthDataReq struct {
	MenuId  uint     `p:"menuId" dc:"菜单ID"`
	Scope   int      `p:"scope" dc:"数据范围"`
	DeptIds []uint64 `p:"deptIds" dc:"部门ID列表"`
}

type ScopeAuthData struct {
	Id        uint64  `json:"id"`
	RoleId    uint    `json:"roleId"`
	MenuId    uint    `json:"menuId"`
	DataScope int     `json:"dataScope"`
	DeptIds   []int64 `json:"deptIds"`
}
