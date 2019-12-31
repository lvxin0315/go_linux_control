package main

import (
	"flag"
	"fmt"
	"github.com/lvxin0315/go_linux_control/common"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var funcName string
var appName string
var appRemark string

func init() {
	flag.StringVar(&funcName,
		"funcName",
		"initDB",
		"操作名称")

	flag.StringVar(&appName,
		"appName",
		"appName",
		"应用名称")

	flag.StringVar(&appRemark,
		"appRemark",
		"appRemark",
		"应用备注")

}

func main() {
	flag.Parse()
	switch funcName {
	case "initDB":
		initDB()
	case "createApp":
		createApp()
	case "startServer":
		startServer()
	}

}

//初始化数据库结构
func initDB() {
	if fileStat, _ := common.PathExists("./tmp/db_init.lock"); fileStat {
		logrus.Info("非第一次启动")
		return
	}
	//初始化db
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	//建表
	tableList := []interface{}{
		&model.App{},
		&model.Cmd{},
		&model.SendCmd{}}
	//先删除再创建
	db.DropTableIfExists(tableList...)
	db.CreateTable(tableList...)
	//建立lock文件
	err = ioutil.WriteFile("./tmp/db_init.lock", []byte("go_linux_control"), 0666) //写入文件(字节数组)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}

//创建app
func createApp() {
	//生成
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	app := new(model.App)
	app.Secret = common.UniqueId()
	app.Name = appName
	db.Create(app)
	if app.ID <= 0 {
		panic("createApp error")
	}
	//成功输出 Secret
	fmt.Println(fmt.Sprintf("createApp success, Secret is 「%s」", app.Secret))
}

//启动监听服务
func startServer() {

}
