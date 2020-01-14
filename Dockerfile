FROM golang

EXPOSE 8088

ENV TZ Asia/Shanghai
ENV GO111MODULE on

RUN echo 'Asia/Shanghai' >/etc/timezone

RUN mkdir -p /data/go_linux_control
RUN chmod -R 777 /data/go_linux_control

WORKDIR /data/go_linux_control

ADD . .

RUN chmod +x ./start.sh

#编译主服务端
RUN go build -o main_server main.go
#编译web服务端
RUN go build -o web_server web.go

##减少无用文件
#RUN find . -name "*.go"  | xargs rm -f

CMD ./start.sh