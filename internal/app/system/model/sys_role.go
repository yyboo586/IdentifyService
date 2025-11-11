package model

import (
	"IdentifyService/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type EffectiveTimeInfo struct {
	EffectiveType int           `json:"effectiveType"`
	WeekDay       []int         `json:"weekDay"`
	DayRange      []*gtime.Time `json:"dayRange"`
	DateRange     []*gtime.Time `json:"dateRange"`
}

type RoleInfoRes struct {
	*entity.SysRole
	*EffectiveTimeInfo
}

const (
	RoleTypeServiceProviderName = "服务提供商" // 服务提供商
	RoleTypeMerchantName        = "展商"    // 商户
)

type RoleStatus int

const (
	RoleStatusDisabled RoleStatus = 0 // 禁用
	RoleStatusNormal   RoleStatus = 1 // 正常
)

type Role struct {
	ID            uint        `json:"id"`
	Pid           uint        `json:"pid"`
	Status        RoleStatus  `json:"status"`
	ListOrder     uint        `json:"list_order"`
	Name          string      `json:"name"`
	Remark        string      `json:"remark"`
	DataScope     uint        `json:"data_scope"` // 1:全部数据权限、2:自定义数据权限、3:本部门数据权限、4:本部门及以下数据权限
	CreatedAt     *gtime.Time `json:"created_at"`
	UpdatedAt     *gtime.Time `json:"updated_at"`
	UserCnt       uint        `json:"user_cnt"`
	CreatedBy     uint64      `json:"created_by"`
	EffectiveTime string      `json:"effective_time"`
}

func ConvertRoleEntity(in *entity.SysRole) *Role {
	return &Role{
		ID:        in.Id,
		Pid:       in.Pid,
		Status:    RoleStatus(in.Status),
		ListOrder: in.ListOrder,
		Name:      in.Name,
		Remark:    in.Remark,
		DataScope: in.DataScope,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
