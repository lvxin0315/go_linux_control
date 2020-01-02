package model

import "github.com/jinzhu/gorm"

type SendCmd struct {
	gorm.Model
	App    *App   `gorm:"ForeignKey:AppId"` //执行的应用
	AppId  uint   `gorm:"column:app_id"`
	Cmd    *Cmd   `gorm:"ForeignKey:Cmd"` //发送的命令
	CmdId  uint   `gorm:"column:cmd_id"`
	Result string `gorm:"column:result"` //执行命令结果
}

func (SendCmd) TableName() string {
	return "send_cmd"
}
