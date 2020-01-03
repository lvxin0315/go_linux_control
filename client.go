package main

import (
	"flag"
	"fmt"
	"github.com/lvxin0315/go_linux_control/common"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"os/exec"
)

var natsUrl string
var appSecret string

func init() {
	flag.StringVar(&natsUrl,
		"natsUrl",
		"127.0.0.1:4222",
		"nats地址")

	flag.StringVar(&appSecret,
		"appSecret",
		"appSecret",
		"应用秘钥")
}

var nc *nats.Conn

func main() {
	flag.Parse()
	//开始监听nats
	listenNats()
}

func listenNats() {

	natsConn, err := nats.Connect(natsUrl)
	if err != nil {
		panic(err)
	}
	nc = natsConn
	//监听开始
	logrus.Info("nats已经连接")
	nc.Subscribe(fmt.Sprintf("app.%s", appSecret), CmdClientRunner)
	select {}
}

func CmdClientRunner(m *nats.Msg) {
	if len(m.Data) <= 0 {
		logrus.Info("为什么是空的？")
		return
	}
	//执行命令, 并返回内容
	err := nc.Publish(m.Reply, exeSysCommand(m.Data))
	if err != nil {
		logrus.Error("CmdClientRunner nc.Publish: ", err)
	}
}

//执行命令生成临时shell文件，再执行shell
func exeSysCommand(cmdStr []byte) []byte {
	logrus.Info("cmd:", string(cmdStr))
	//生成文件
	shellFileName, err := common.SaveShellFile(cmdStr)
	//执行后删除文件
	defer common.DeleteShellFile(shellFileName)
	if err != nil {
		logrus.Error(err)
		return []byte{}
	}
	//执行内容
	cmd := exec.Command("sh", "-c", shellFileName)
	opBytes, err := cmd.Output()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return opBytes
}
