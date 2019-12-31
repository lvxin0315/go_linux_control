package db_conn

import (
	"github.com/Unknwon/goconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func GetGormDB() (*gorm.DB, error) {
	cfg, err := goconfig.LoadConfigFile("etc/config.ini")
	if err != nil {
		logrus.Error("goconfig.LoadConfigFile is error:", err)
		return nil, err
	}
	conn, err := cfg.GetValue("gorm", "conn")
	if err != nil {
		logrus.Error("cfg.GetValue is error:", err)
		return nil, err
	}
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		logrus.Info("conn:", conn)
		logrus.Error("GetGormDB is error: ", err)
		return nil, err
	}
	return db, nil
}
