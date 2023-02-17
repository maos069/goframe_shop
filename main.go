package main

import (
	_ "goframe_shop/internal/logic"
	_ "goframe_shop/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe_shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
