package snowIDGen

import (
	"IdentifyService/internal/app/common/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sony/sonyflake"
)

var machineID uint16 = 1

func init() {
	service.RegisterSnowID(New())
}

func New() service.ISnowID {
	return &sSnowID{
		sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: gtime.NewFromStr("2010-05-01").Time,
			MachineID: GetMachineId,
		}),
	}
}

type sSnowID struct {
	*sonyflake.Sonyflake
}

func (s *sSnowID) GenID() (uint64, error) {
	return s.NextID()
}

func GetMachineId() (uint16, error) {
	return machineID, nil
}
