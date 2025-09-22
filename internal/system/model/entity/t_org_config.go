package entity

import "github.com/gogf/gf/v2/os/gtime"

type OrgConfig struct {
	ID        string      `orm:"id"        description:"组织配置id"`
	OrgID     string      `orm:"org_id"     description:"组织ID"`
	Key       string      `orm:"key"       description:"配置键"`
	Value     string      `orm:"value"     description:"配置值"`
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"修改时间"`
}
