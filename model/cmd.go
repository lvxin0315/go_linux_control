package model

import "github.com/jinzhu/gorm"

type Cmd struct {
	gorm.Model
	Title string `gorm:"column:title"` //命令标题
	Cmd   string `gorm:"column:cmd"`   //命令内容
	Des   string `gorm:"column:des"`   //命令描述
}

func (Cmd) TableName() string {
	return "cmd"
}
