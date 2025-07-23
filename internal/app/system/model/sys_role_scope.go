package model

type ScopeAuthDataReq struct {
	MenuId  int64    `p:"menuId" dc:"菜单ID"`
	Scope   int      `p:"scope" dc:"数据范围"`
	DeptIds []uint64 `p:"deptIds" dc:"部门ID列表"`
}

type ScopeAuthData struct {
	ID        uint64  `json:"id"`
	RoleId    int64   `json:"roleId"`
	MenuId    int64   `json:"menuId"`
	DataScope int     `json:"dataScope"`
	DeptIds   []int64 `json:"deptIds"`
}
