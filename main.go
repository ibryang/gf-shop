package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "shop/internal/logic"
	_ "shop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
