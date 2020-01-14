package runner

import (
	"github.com/Unknwon/goconfig"
	"github.com/lvxin0315/go_linux_control/service"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
)

//启动监听服务
func StartServer() {
	cfg, err := goconfig.LoadConfigFile("etc/config.ini")
	if err != nil {
		logrus.Error("goconfig.LoadConfigFile 「config.ini」 is error:", err)
		panic(err)
	}
	conn, err := cfg.GetValue("nats", "host")
	if err != nil {
		logrus.Error("cfg.GetValue 「config.ini」 is error:", err)
		panic(err)
	}
	nc, err := nats.Connect(conn)
	if err != nil {
		panic(err)
	}
	//监听开始
	logrus.Info("nats已经连接")

	//系统定制消息接收
	nc.Subscribe("disk_use.*", service.GetDiskUseMessage)
	//自定义消息接收
	nc.Subscribe("cmd.*", service.GetMessage)

	select {}
}
