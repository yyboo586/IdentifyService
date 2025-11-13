package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
)

func IsValidAgreementName(name string) bool {
	return name == "用户服务协议" || name == "隐私保护协议"
}

type Agreement struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	MajorVersion string      `json:"major_version"`
	MinorVersion string      `json:"minor_version"`
	PatchVersion string      `json:"patch_version"`
	Content      string      `json:"content"`
	CreatedAt    *gtime.Time `json:"created_at"`
	UpdatedAt    *gtime.Time `json:"updated_at"`
}

// VersionString 返回格式化的版本号字符串，如 "1.0.0"
func (a *Agreement) VersionString() string {
	return fmt.Sprintf("%s.%s.%s", a.MajorVersion, a.MinorVersion, a.PatchVersion)
}

type AgreementCreateInput struct {
	Name         string
	MajorVersion string
	Content      string
}

type AgreementUpdateInput struct {
	ID         int64
	Name       string
	UpdateType string // "minor" 或 "patch"
	Content    string
}

type AgreementListInput struct {
	Name string
	PageReq
}

// CompareVersion 比较两个版本号字符串，返回 -1, 0, 1
func CompareVersion(v1, v2 string) int {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	maxLen := len(v1Parts)
	if len(v2Parts) > maxLen {
		maxLen = len(v2Parts)
	}

	for i := 0; i < maxLen; i++ {
		var v1Val, v2Val int
		if i < len(v1Parts) {
			v1Val, _ = strconv.Atoi(v1Parts[i])
		}
		if i < len(v2Parts) {
			v2Val, _ = strconv.Atoi(v2Parts[i])
		}

		if v1Val < v2Val {
			return -1
		}
		if v1Val > v2Val {
			return 1
		}
	}
	return 0
}

// IncrementVersion 递增版本号
func IncrementVersion(version string, part string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format: %s", version)
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch part {
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		return "", fmt.Errorf("invalid update type: %s", part)
	}

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

type UserAgreement struct {
	ID            int64       `json:"id"`
	UserID        string      `json:"user_id"`
	AgreementID   int64       `json:"agreement_id"`
	AgreementName string      `json:"agreement_name"`
	Agreed        bool        `json:"agreed"`
	CreatedAt     *gtime.Time `json:"created_at"`
}
