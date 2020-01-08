package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/go_linux_control/format"
)

type SendCmd struct {
	gorm.Model
	App          *App             `gorm:"ForeignKey:AppId"` //执行的应用
	AppId        uint             `gorm:"column:app_id"`
	Cmd          *Cmd             `gorm:"ForeignKey:Cmd"` //发送的命令
	CmdId        uint             `gorm:"column:cmd_id"`
	Result       string           `gorm:"type:text;column:result"`      //执行命令结果
	ResultToJson string           `gorm:"type:text;column:result_json"` //执行命令结果json
	format       format.CmdFormat //格式化方法
}

func (m *SendCmd) TableName() string {
	return "send_cmd"
}

func (m *SendCmd) BeforeUpdate() {
	if m.format != nil {
		m.ResultToJson = m.format.ToJson(m.Result)
	}
}

func (m *SendCmd) SetFormat(format format.CmdFormat) {
	m.format = format
}
