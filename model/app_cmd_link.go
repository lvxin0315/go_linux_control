package model

import "github.com/jinzhu/gorm"

type AppCmdLink struct {
	gorm.Model
	App   *App `gorm:"ForeignKey:AppId"` //执行的应用
	AppId uint `gorm:"column:app_id"`
	Cmd   *Cmd `gorm:"ForeignKey:Cmd"` //发送的命令
	CmdId uint `gorm:"column:cmd_id"`
}

func (AppCmdLink) TableName() string {
	return "app_cmd_link"
}
