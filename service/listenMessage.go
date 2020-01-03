package service

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"strings"
)

func GetMessage(m *nats.Msg) {
	go SaveMessage(m.Subject, m.Data)
}

//保存返回值到db
func SaveMessage(sendId string, result []byte) {
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	//保存记录
	sendCmd := new(model.SendCmd)
	db.First(sendCmd, map[string]interface{}{
		"id": strings.ReplaceAll(sendId, "cmd.", ""),
	})
	if sendCmd.ID <= 0 {
		logrus.Error("消息记录未找到：", sendId)
		return
	}
	sendCmd.Result = string(result)
	db.Save(sendCmd)
}

//发布消息，先使用短连接方式发送消息
func SendCmdMessage(app *model.App, cmd *model.Cmd) {
	//生成send_cmd 记录
	sendCmd := new(model.SendCmd)
	sendCmd.AppId = app.ID
	sendCmd.CmdId = cmd.ID
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		return
	}
	db.Create(sendCmd)
	if sendCmd.ID <= 0 {
		logrus.Error("sendCmd保存失败")
		return
	}
	cfg, err := goconfig.LoadConfigFile("etc/config.ini")
	if err != nil {
		logrus.Error("goconfig.LoadConfigFile 「config.ini」 is error:", err)
		return
	}
	conn, err := cfg.GetValue("nats", "host")
	if err != nil {
		logrus.Error("cfg.GetValue 「config.ini」 is error:", err)
		return
	}
	nc, err := nats.Connect(conn)
	defer nc.Close()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("send cmd:", cmd.Cmd)
	err = nc.PublishRequest(
		fmt.Sprintf("app.%s", app.Secret),
		fmt.Sprintf("cmd.%d", sendCmd.ID),
		[]byte(cmd.Cmd))
	if err != nil {
		logrus.Error(err)
		return
	}
}
