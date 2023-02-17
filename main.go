package main

import (
	_ "goframe_shop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe_shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
