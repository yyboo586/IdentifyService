package model

import (
	"IdentifyService/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}

type LinkDeptRes struct {
	DeptId   uint64 `json:"deptId"`
	DeptName string `json:"deptName"`
}

type Dept struct {
	DeptID    uint64 `json:"dept_id"`
	DeptName  string `json:"dept_name"`
	ManagerID string `json:"manager_id"`

	ParentID  int    `json:"parent_id"`
	Ancestors string `json:"ancestors"`
	Leader    []int  `json:"leader"`

	OrderNum int    `json:"order_num"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   int    `json:"status"`

	CreatedBy int64 `json:"created_by"`
	UpdatedBy int64 `json:"updated_by"`

	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
	DeletedAt *gtime.Time `json:"deleted_at"`
}

func ConvertDeptEntityToModel(in *entity.SysDept) (out *Dept) {
	return &Dept{
		DeptID:    in.DeptId,
		DeptName:  in.DeptName,
		ManagerID: in.ManagerID,

		ParentID:  int(in.ParentId),
		Ancestors: in.Ancestors,
		Leader:    in.Leader,

		OrderNum: in.OrderNum,
		Phone:    in.Phone,
		Email:    in.Email,
		Status:   int(in.Status),

		CreatedBy: int64(in.CreatedBy),
		UpdatedBy: int64(in.UpdatedBy),

		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		DeletedAt: in.DeletedAt,
	}
}
