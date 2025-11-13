package model

import "github.com/gogf/gf/v2/frame/g"

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}

type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

type PageReq struct {
	Page int `p:"page" dc:"当前页码,默认1"`
	Size int `p:"size" dc:"每页条数,默认10"`
}

type PageRes struct {
	Total       int `json:"total" dc:"总条数"`
	CurrentPage int `json:"current_page" dc:"当前页码"`
}
