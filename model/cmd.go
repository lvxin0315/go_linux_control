package model

import "github.com/jinzhu/gorm"

type Cmd struct {
	gorm.Model
	Title    string `gorm:"column:title"`         //命令标题
	Cmd      string `gorm:"type:text;column:cmd"` //命令内容
	Des      string `gorm:"type:text;column:des"` //命令描述
	IsSystem bool   `gorm:"column:is_system"`     //系统自有命令
	RouteKey string `gorm:"column:route_key"`     //消息接收队列标识
}

func (Cmd) TableName() string {
	return "cmd"
}
