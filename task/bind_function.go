package task

import (
	"context"

	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
	"IdentifyService/internal/mounter"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	mounter.Mount(func(ctx context.Context, s *ghttp.Server) {
		Run()
	})
}

func Run() {
	task1 := &model.TimeTask{
		FuncName: "test1",
		Run:      Test1,
	}
	task2 := &model.TimeTask{
		FuncName: "test2",
		Run:      Test2,
	}
	checkUserOnlineTask := &model.TimeTask{
		FuncName: "checkUserOnline",
		Run:      service.SysUserOnline().CheckUserOnline,
	}
	service.TaskList().AddTask(task1)
	service.TaskList().AddTask(task2)
	service.TaskList().AddTask(checkUserOnlineTask)
	ctx := gctx.New()
	//自动执行已开启的任务
	jobs, err := service.SysJob().GetJobs(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	for _, job := range jobs {
		service.SysJob().JobStart(ctx, job)
	}
}
