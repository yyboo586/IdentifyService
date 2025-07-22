package service

type (
	ISnowID interface {
		GenID() (uint64, error)
	}
)

var (
	localSnowID ISnowID
)

func SnowID() ISnowID {
	if localSnowID == nil {
		panic("implement not found for interface ISnowID, forgot register?")
	}
	return localSnowID
}

func RegisterSnowID(i ISnowID) {
	localSnowID = i
}
