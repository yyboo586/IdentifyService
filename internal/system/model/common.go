package model

import (
	"errors"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	DefaultPageSize = 10
)

type Author struct {
	Authorization string `p:"Authorization" v:"required" in:"header" dc:"Bearer {{token}}"`
}

// PageReq 公共请求参数
type PageReq struct {
	PageNum  int `json:"page_num"`  // 当前页码
	PageSize int `json:"page_size"` // 每页数
}

// PageRes 列表公共返回
type PageRes struct {
	CurrentPage int         `json:"current_page"`
	Total       interface{} `json:"total"`
}

// EmptyRes 不响应任何数据
type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

// DB Error
var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrRecordAlreadyExists = errors.New("record already exists")
)

var (
	ErrServerInternal = errors.New("server internal error")
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrForbidden  = errors.New("forbidden")
)

var (
	ErrPermissionDenied = errors.New("permission denied")
)
