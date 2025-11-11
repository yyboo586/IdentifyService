package main

import (
	_ "IdentifyService/internal/app/boot"
	_ "IdentifyService/internal/app/system/packed"
	"IdentifyService/internal/cmd"
	_ "IdentifyService/library/libValidate"
	_ "IdentifyService/task"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
