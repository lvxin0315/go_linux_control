package service

import (
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/format"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"strings"
)

func GetDiskUseMessage(m *nats.Msg) {
	go SaveDiskUseMessage(m.Subject, m.Data)
}

//保存返回值到db
func SaveDiskUseMessage(sendId string, result []byte) {
	db := db_conn.GetGormDB()
	defer db.Close()
	//保存记录
	sendCmd := new(model.SendCmd)
	db.First(sendCmd, map[string]interface{}{
		"id": strings.ReplaceAll(sendId, "disk_use.", ""),
	})
	if sendCmd.ID <= 0 {
		logrus.Error("消息记录未找到：", sendId)
		return
	}
	sendCmd.Result = string(result)
	sendCmd.SetFormat(new(format.DFHFormat))
	db.Save(sendCmd)
}
