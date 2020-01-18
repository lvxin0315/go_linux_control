package service

import (
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/format"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"strings"
)

func GetMemInfoMessage(m *nats.Msg) {
	go SaveMemInfoMessage(m.Subject, m.Data)
}

//保存返回值到db
func SaveMemInfoMessage(sendId string, result []byte) {
	db := db_conn.GetGormDB()
	defer db.Close()
	//保存记录
	sendCmd := new(model.SendCmd)
	db.First(sendCmd, map[string]interface{}{
		"id": strings.ReplaceAll(sendId, "meminfo.", ""),
	})
	if sendCmd.ID <= 0 {
		logrus.Error("消息记录未找到：", sendId)
		return
	}
	sendCmd.Result = string(result)
	sendCmd.SetFormat(new(format.MemInfoFormat))
	db.Save(sendCmd)
}
