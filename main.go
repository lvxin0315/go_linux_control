package main

import (
	"flag"
	"github.com/lvxin0315/go_linux_control/runner"
	"github.com/sirupsen/logrus"
)

var funcName string
var appName string
var appRemark string

func init() {
	flag.StringVar(&appName,
		"appName",
		"appName",
		"应用名称")

	flag.StringVar(&appRemark,
		"appRemark",
		"appRemark",
		"应用备注")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		_ = flag.CommandLine.Parse(args[1:])
	}
	funcName = args[0]
	switch funcName {
	case "initDB":
		runner.InitDB()
	case "createApp":
		runner.CreateApp(appName, appRemark)
	case "startServer":
		runner.StartServer()
	case "help":
		help()
	default:
		logrus.Error("Command is missing")
	}
}

func help() {
	helpContent := `
-- initDB:初始化mysql,配置etc/config.ini;
-- createApp:创建应用,-appName, -appRemark
-- startServer:开启服务端`
	logrus.Info(helpContent)
}
