package model

type TokenInfo struct {
	UserID       string `json:"user_id"`
	UserNickname string `json:"user_nickname"`

	DeviceID string `json:"device_id"`
}
