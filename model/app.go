package model

import "github.com/jinzhu/gorm"

type App struct {
	gorm.Model
	Name   string `gorm:"column:name"`            //应用名称
	Remark string `gorm:"column:remark"`          //应用备注
	Secret string `gorm:"column:secret;size:255"` //32位秘钥
}

func (App) TableName() string {
	return "app"
}
