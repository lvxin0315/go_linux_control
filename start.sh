#!/bin/bash

cd /data/go_linux_control || exit
#启动web
nohup ./web_server &
#启动主函数
./main_server startServer
