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
