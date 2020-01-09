package main

import (
	"flag"
	"github.com/lvxin0315/go_linux_control/runner"
	"github.com/sirupsen/logrus"
)

var funcName string

var appName, appRemark string

var cmdTitle, cmdDes, cmdStr string

var appId, cmdId string

var helpContent = `
-- initDB:初始化mysql,配置etc/config.ini;
-- createApp:创建应用,-appName, -appRemark
-- createCmd:创建命令,-cmdTitle, -cmdDes, -cmdStr
-- sendCmd:发送命令,-appId, -cmdId
-- startServer:开启服务端`

func init() {
	flag.StringVar(&appName,
		"appName",
		"appName",
		"应用名称")

	flag.StringVar(&appRemark,
		"appRemark",
		"appRemark",
		"应用备注")

	flag.StringVar(&cmdTitle,
		"cmdTitle",
		"cmdTitle",
		"命令标题")

	flag.StringVar(&cmdDes,
		"cmdDes",
		"cmdDes",
		"命令描述")

	flag.StringVar(&cmdStr,
		"cmdStr",
		"df -h /",
		"命名内容")

	flag.StringVar(&appId,
		"appId",
		"1",
		"应用id")

	flag.StringVar(&cmdId,
		"cmdId",
		"1",
		"命令id")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		logrus.Error("Command is missing, please use help!")
		return
	} else if len(args) > 1 {
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
	case "createCmd":
		runner.CreateCmd(cmdTitle, cmdStr, cmdDes)
	case "sendCmd":
		runner.SendCmd(appId, cmdId)
	case "sendCmdForAllApp":
		runner.SendCmdForAllApp(cmdId)
	default:
		logrus.Error("Command is missing, please use help!")
	}
}

func help() {
	logrus.Info(helpContent)
}
