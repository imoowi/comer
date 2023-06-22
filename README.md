# Comer

[![Go](https://github.com/imoowi/comer/actions/workflows/release-tag.yml/badge.svg)](https://github.com/imoowi/comer/actions?query=workflow%3ACI) 
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/26b6baa851bc426c9bc7dcc9079485b3)](https://app.codacy.com/gh/imoowi/comer/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/imoowi/comer)](https://goreportcard.com/report/github.com/imoowi/comer)

## Comer 是什么？

Comer是一个用go语言写的代码生成工具，能够生成基本的web api框架，同时支持app新增。

## 安装

```go
go install github.com/imoowi/comer@latest
```

## 使用

-    创建项目

```sh
comer --module=moduleName --path=projectDir
```

-    给项目添加app

```sh
cd projectDir
comer genapp --app=appName
#或者
comer genapp --app=user --swaggerTags='Oauth' --handler=auth --service=user --model=user,role
```

-    安装swag

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

-    生成swagger文档

```sh
swag init
```

-    安装air

```sh
go install github.com/cosmtrek/air@latest
```

-    修改数据库配置(./configs/settings-local.yml)

```yml
application: # dev开发环境 test测试环境 prod线上环境
  mode: dev
  name: comerProject # 服务名称
server:
  host: 0.0.0.0 # 服务器ip，默认使用 0.0.0.0
  port: 8000 # 服务端口号
  readtimeout: 60 # 读超时时间
  writertimeout: 60 # 写超时时间
logger:
  path: runtime/logs/log # 日志存放路径
  stdout: "" # 日志输出，file：文件，default：命令行，其他：命令行
  level: trace # 日志等级, trace, debug, info, warn, error, fatal
  maxAge: 168h # 日志最长保存时间，7天, ns、us、ms、s、m、h
  rotationTime: 24h # 日志切割级别
ratelimit:
  # 每秒放多少个令牌
  cap: 1000
  # 每秒取多少个令牌
  quantum: 1000
jwt:
  secret: comerProject-admin # token 密钥，生产环境时及的修改
  timeout: 2000h0m0s # token 过期时间 格式：0h0m0s
  refresh_token_timeout: 0h5m0s # token 过期时间减去的时间，用于刷新token
mysql:
  dsn: root:password@tcp(127.0.0.1:3306)/comer_project_db_YeaVWT?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
  casbin: root:password@tcp(127.0.0.1:3306)/comer_project_db_YeaVWT
# influxdb:
#   addr: http://127.0.0.1:8086
#   token: [token string]
#   org: [orgnization name]
#   bucket: [bucket name]
#   testSwitchOn: false
redis:
  addr: com.redis.host:6379
  password: "password"
  db: 0
cache:
  driver: redis
  prefix: "comer_project_db_YeaVWT:cache"

```

-    数据迁移:生成基本的数据库表

```sh
go run . migrate
```

-    运行项目

```sh
air
#或者
go mod tidy
go run . server
```

-    访问接口文件：[http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)
![](assets/comer-swagger.png)

## 接下来做什么？
听取建议并采纳优秀的解决方案