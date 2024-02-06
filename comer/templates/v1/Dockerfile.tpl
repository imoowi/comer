#FROM 基础镜像
FROM golang:1.21-alpine as builder

#LABEL 指令用来给镜像添加一些元数据（metadata），以键值对的形式
LABEL maintainer="imoowi"

#设置容器环境变量
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct

#为 RUN、CMD、ENTRYPOINT、COPY 和 ADD 设置工作目录，就是切换目录
WORKDIR /go/release

#COPY 拷贝文件或目录到容器中，跟ADD类似，但不具备自动下载或解压的功能
COPY . .

#RUN 构建镜像时运行的指令
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add tzdata

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.10 && swag init
RUN CGO_ENABLED=0 GOOS=linux go build -p 1 -ldflags="-w -s" -a -installsuffix cgo -o {{.moduleProjectName}} .

FROM alpine:latest

COPY --from=builder /go/release/{{.moduleProjectName}} /

COPY --from=builder /go/release/configs /configs

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

#EXPOSE 声明容器的服务端口（仅仅是声明）
EXPOSE 8000

#数据迁移
RUN /{{.moduleProjectName}} migrate -c /configs/settings-local.yml

#初始化数据库
RUN /{{.moduleProjectName}} init -c /configs/settings-local.yml

#CMD 运行容器时执行的shell环境
CMD ["/{{.moduleProjectName}}","server","-c", "/configs/settings-local.yml"]
