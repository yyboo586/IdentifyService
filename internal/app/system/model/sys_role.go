/*
* @desc:角色
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:11
 */

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
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
