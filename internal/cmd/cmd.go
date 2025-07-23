package cmd

import (
	"context"

	"IdentifyService/internal/mounter"
	"IdentifyService/internal/router"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
			s := g.Server()
			//调用注册已挂载相关组件
			mounter.DoMount(ctx, s)
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.R.BindController(ctx, group)
			})
			s.SetOpenApiPath("/api.json")
			s.SetSwaggerPath("/swagger")
			s.Run()
			return nil
		},
	}
)
