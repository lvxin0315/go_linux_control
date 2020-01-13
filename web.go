package main

import (
	"github.com/lvxin0315/go_linux_control/runner"
	_ "github.com/lvxin0315/go_linux_control/web/controller"
)

func main() {
	runner.StartWebServer()
}
