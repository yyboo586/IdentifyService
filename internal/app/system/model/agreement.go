package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
)

const (
	AgreementStatusDraft     = 0
	AgreementStatusPublished = 1
	AgreementStatusArchived  = 2
)

func IsValidAgreementName(name string) bool {
	return name == "用户服务协议" || name == "隐私保护协议" || name == "支付协议"
}

type Agreement struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	MajorVersion int         `json:"major_version"`
	MinorVersion int         `json:"minor_version"`
	PatchVersion int         `json:"patch_version"`
	VersionCode  int         `json:"version_code"`
	Status       int         `json:"status"`
	Content      string      `json:"content"`
	PublishedAt  *gtime.Time `json:"published_at"`
	CreatedAt    *gtime.Time `json:"created_at"`
	UpdatedAt    *gtime.Time `json:"updated_at"`
}

// VersionString 返回格式化的版本号字符串，如 "1.0.0"
func (a *Agreement) VersionString() string {
	return fmt.Sprintf("%d.%d.%d", a.MajorVersion, a.MinorVersion, a.PatchVersion)
}

type AgreementCreateInput struct {
	Name       string
	Major      int
	Minor      int
	Patch      int
	Content    string
	PublishNow bool
}

type AgreementUpdateInput struct {
	ID         int64
	Content    string
	PublishNow bool
	Status     *int
}

type AgreementListInput struct {
	Name   string
	Status *int
	PageReq
}

func ParseVersion(version string) (int, int, int, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid version format: %s", version)
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, err
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, err
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return major, minor, patch, nil
}

func BuildVersionCode(major, minor, patch int) int {
	return major*10000 + minor*100 + patch
}

type UserAgreement struct {
	ID            int64       `json:"id"`
	UserID        string      `json:"user_id"`
	AgreementID   int64       `json:"agreement_id"`
	AgreementName string      `json:"agreement_name"`
	VersionCode   int         `json:"version_code"`
	Agreed        bool        `json:"agreed"`
	CreatedAt     *gtime.Time `json:"created_at"`
}
