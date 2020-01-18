package model

import "github.com/jinzhu/gorm"

type AppStatus struct {
	gorm.Model
	AppId  uint `gorm:"column:app_id" json:"app_id"`
	App    *App
	Cpu    string `gorm:"column:cpu" json:"cpu"`
	Mem    string `gorm:"column:mem" json:"mem"`
	Disk   string `gorm:"column:disk" json:"disk"`
	Online bool   `gorm:"column:online" json:"online"`
}

func (AppStatus) TableName() string {
	return "app_status"
}
