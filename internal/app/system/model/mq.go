package model

type MQEventType int

const (
	_                        MQEventType = iota
	MQEventForceDeviceLogout             // 强制设备下线
)

type ForceDeviceLogout struct {
	Title      string `json:"title"`
	UserID     string `json:"user_id"`
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	ForceTime  string `json:"force_time"`
}
