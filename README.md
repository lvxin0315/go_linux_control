# go_linux_control
自动远程控制linux服务器

------------

- 1.main.go 管理端
- 2.client.go 客户端

------------

### 管理端

- 1.使用gorm + mysql存储
- 2.使用nats进行消息分发和读取
- 3.初始化db：main initDB
- 4.创建：main createApp -appName="test1" -appRemark="remark"
- 5.启动服务：main startServer
- 6.创建命令：main createCmd -cmdTitle="cmd1" -cmdDes="cmd1Des" -cmdStr="du -sh *"
