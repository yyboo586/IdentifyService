package model

import (
	"IdentifyService/internal/app/system/model/entity"
)

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}

type LinkDeptRes struct {
	DeptId   uint64 `json:"deptId"`
	DeptName string `json:"deptName"`
}
