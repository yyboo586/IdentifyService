package main

import (
	"IdentifyService/internal/app/system/logics"
	"IdentifyService/internal/app/system/logics/sysJob"
	"IdentifyService/internal/app/system/logics/sysJobLog"
	"IdentifyService/internal/app/system/logics/sysNotice"
	"IdentifyService/internal/app/system/logics/sysNoticeRead"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/internal/router"
	_ "IdentifyService/library/libValidate"
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	g.Log().SetTimeFormat("2006-01-02 15:04:05")

	s := g.Server()
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")

	service.RegisterOrg(logics.NewOrg())
	service.RegisterUser(logics.NewUser())
	service.RegisterRole(logics.NewRole())
	service.RegisterAuthRule(logics.NewAuthRule())
	service.RegisterLog(logics.NewLog())
	service.RegisterTokenService()
	service.RegisterContextService()

	service.RegisterSysJob(sysJob.New())
	service.RegisterSysJobLog(sysJobLog.New())
	service.RegisterSysNotice(sysNotice.New())
	service.RegisterSysNoticeRead(sysNoticeRead.New())
	service.RegisterMQService()
	service.RegisterMessagePushService()

	s.Group("/", func(group *ghttp.RouterGroup) {
		router.R.BindController(context.Background(), group)
		router.R.BindWebsocketModuleController(context.Background(), group)
	})

	s.Run()
}

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Address string `json:"address"`

	OpenApiPath string `yaml:"openapiPath"`
	SwaggerPath string `yaml:"swaggerPath"`
}

type LoggerConfig struct {
	Level      string `yaml:"level"`
	Stdout     string `yaml:"stdout"`
	TimeFormat string `yaml:"timeFormat"`

	Path string `yaml:"path"`
	File string `yaml:"file"`

	ErrorStack       bool   `yaml:"errorStack"`
	ErrorLogEnabled  bool   `yaml:"errorLogEnabled"`
	ErrorLogPattern  string `yaml:"errorLogPattern"`
	AccessLogEnabled bool   `yaml:"accessLogEnabled"`
	AccessLogPattern string `yaml:"accessLogPattern"`
}
