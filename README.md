# go_linux_control
自动远程控制linux服务器

------------

- 1.main.go 管理端
- 2.client.go 客户端
- 3.web.go  web端

------------

### 管理端

- 1.使用gorm + mysql存储
- 2.使用nats进行消息分发和读取
- 3.初始化db：main initDB
- 4.创建：main createApp -appName="test1" -appRemark="remark"
- 5.启动服务：main startServer
- 6.创建命令：main createCmd -cmdTitle="cmd1" -cmdDes="cmd1Des" -cmdStr="du -sh *"
- 7.发送命令：main sendCmd -appId="1" -cmdId="2"
- 8.全部发送命令：main sendCmdForAllApp -cmdId="2"


### 客户端
- 1.开启服务 client -natsUrl="" -appSecret=""


### 笔记
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

go build -ldflags "-X 'main.NatsUrl=192.168.0.230:4222' -X 'main.AppSecret=ss'" -o linux_control_client client.go
