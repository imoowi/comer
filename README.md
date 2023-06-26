# Comer

[![Go](https://github.com/imoowi/comer/actions/workflows/release-tag.yml/badge.svg)](https://github.com/imoowi/comer/actions?query=workflow%3ACI)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/26b6baa851bc426c9bc7dcc9079485b3)](https://app.codacy.com/gh/imoowi/comer/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/imoowi/comer)](https://goreportcard.com/report/github.com/imoowi/comer)
[![codecov](https://codecov.io/gh/imoowi/comer/branch/main/graph/badge.svg)](https://codecov.io/gh/imoowi/comer)

## Comer 是什么？

Comer是一个用go语言写的代码生成工具，能够生成基本的web api框架，同时支持app新增；其中包括gin 、gorm、redis、casbin、auth、captcha等

## 用Comer生成的项目结构
```sh
$ tree
.
|-- README.md
|-- apps //应用集合
|   |-- apps.go //多个应用自动加载文件
|   |-- common //公共模块
|   |   |-- handlers //路由处理方
|   |   |   `-- captcha.handler.go //默认的验证码
|   |   `-- router.go //路由
|   |-- student //通过 comer genapp --app=appName生成
|   |   |-- handlers
|   |   |   `-- student.handler.go
|   |   |-- migrates
|   |   |   `-- student.migrate.go
|   |   |-- models
|   |   |   `-- student.model.go
|   |   |-- repos
|   |   |   `-- student.repo.go
|   |   |-- router.go
|   |   `-- services
|   |       `-- student.service.go
|   |-- swagger //默认包含的swagger文档应用
|   |   `-- router.go
|   `-- user //默认包含用户应用
|       |-- handlers
|       |   `-- auth.handler.go //用户认证
|       |-- migrates //数据迁移文件，会自动生成
|       |   |-- role.migrate.go
|       |   |-- user.migrate.go
|       |   |-- userlog.migrate.go
|       |   `-- userrole.migrate.go
|       |-- models //数据表对应的model，自动生成，表结构通过结构体修改，禁止直接修改数据库里表的结构
|       |   |-- role.model.go //角色
|       |   |-- user.model.go //用户
|       |   |-- userlog.model.go //用户记录
|       |   `-- userrole.model.go //用户角色关系
|       |-- repos //数据提供方
|       |   |-- role.repo.go
|       |   |-- user.repo.go
|       |   |-- userlog.repo.go
|       |   `-- userrole.repo.go
|       |-- router.go //路由
|       `-- services //服务提供方
|           |-- role.service.go
|           |-- user.service.go
|           |-- userlog.service.go
|           `-- userrole.service.go
|-- cmd //由Cobra命令生成
|   |-- init.go //系统初始化
|   |-- migrate.go //数据迁移
|   |-- root.go //主入口
|   `-- server.go api服务
|-- components //组件
|   |-- captcha.go
|   |-- mysql.go
|   `-- redis.go
|-- configs //配置目录
|   |-- casbin.conf
|   `-- settings-local.yml
|-- docs //swagger生成的apidoc
|   |-- docs.go
|   |-- swagger.json
|   `-- swagger.yaml
|-- global //全局文件
|   |-- cache.go
|   |-- casbin.go
|   |-- config.go
|   |-- global.go
|   |-- log.go
|   |-- mysql.go
|   `-- redis.go
|-- go.mod
|-- go.sum
|-- main.go //程序主入口
|-- middlewares //中间件
|   |-- CasbinMiddleware.go //权限控制
|   |-- CrosMiddleware.go //跨域访问
|   |-- JWTAuthMiddleware.go //JWT认证
|   |-- LoggerMiddleware.go //日志
|   |-- RateLimitMiddleware.go //访问频率控制
|   |-- VcodeMiddleware.go //验证码中间件
|   |-- middleware.go
|   `-- token //jwttoken
|       `-- jwttoken.go
|-- router
|   `-- router.go //路由定义
|-- runtime //运行时
|   `-- log2023062615.log
`-- utils //工具箱
    |-- copy
    |   `-- copy.go
    |-- format
    |   `-- format.go
    |-- logger.go
    |-- maker
    |   `-- maker.go
    |-- myfile
    |   `-- myfile.go
    |-- mytime
    |   |-- mytime.go
    |   |-- translater.go
    |   `-- week.go
    |-- office
    |   `-- excel.go
    |-- password
    |   `-- password.go
    |-- request
    |   |-- http.go
    |   `-- pages.go
    |-- response
    |   |-- pages.go
    |   `-- response.go
    |-- slice
    |   `-- slice.go
    `-- utils.go

36 directories, 77 files
```
## 安装

```go
go install github.com/imoowi/comer@latest
```

## 使用

- 创建项目

```sh
comer --module=github.com/imoowi/comer-example --path=comer-example
```
如下:

```sh
$ comer --path=comer-example --module=github.com/imoowi/comer-example
2023/06/26 14:46:34 go.mod not exists
Comer version  v1.1.6

_________
\_   ___ \   ____    _____    ____  _______
/    \  \/  /  _ \  /     \ _/ __ \ \_  __ \
\     \____(  <_> )|  Y Y  \\  ___/  |  | \/
 \______  / \____/ |__|_|  / \___  > |__|
                \/               \/      \/ v1.1.6, built with go1.20.2
dir [ comer-example/apps ] created
dir [ comer-example/cmd ] created
dir [ comer-example/components ] created
dir [ comer-example/apps/common ] created
dir [ comer-example/apps/common/handlers ] created
dir [ comer-example/apps/swagger ] created
dir [ comer-example/apps/user/handlers ] created
dir [ comer-example/apps/user/migrates ] created
dir [ comer-example/apps/user/models ] created
dir [ comer-example/apps/user/repos ] created
dir [ comer-example/apps/user/services ] created
dir [ comer-example/configs ] created
dir [ comer-example/global ] created
dir [ comer-example/middlewares ] created
dir [ comer-example/middlewares/token ] created
dir [ comer-example/router ] created
dir [ comer-example/runtime ] created
dir [ comer-example/utils ] created
dir [ comer-example/utils/copy ] created
dir [ comer-example/utils/format ] created
dir [ comer-example/utils/maker ] created
dir [ comer-example/utils/myfile ] created
dir [ comer-example/utils/mytime ] created
dir [ comer-example/utils/office ] created
dir [ comer-example/utils/password ] created
dir [ comer-example/utils/request ] created
dir [ comer-example/utils/response ] created
dir [ comer-example/utils/slice ] created
dir [ comer-example/.vscode ] created
file [ comer-example/global/global.go ] created
file [ comer-example/global/cache.go ] created
file [ comer-example/middlewares/LoggerMiddleware.go ] created
file [ comer-example/utils/myfile/myfile.go ] created
file [ comer-example/utils/mytime/mytime.go ] created
file [ comer-example/apps/user/handlers/auth.handler.go ] created
file [ comer-example/components/mysql.go ] created
file [ comer-example/configs/settings-local.yml ] created
file [ comer-example/apps/user/models/userlog.model.go ] created
file [ comer-example/cmd/migrate.go ] created
file [ comer-example/utils/mytime/translater.go ] created
file [ comer-example/go.mod ] created
file [ comer-example/.vscode/launch.json ] created
file [ comer-example/apps/user/router.go ] created
file [ comer-example/apps/user/services/user.service.go ] created
file [ comer-example/middlewares/middleware.go ] created
file [ comer-example/apps/common/router.go ] created
file [ comer-example/apps/user/models/userrole.model.go ] created
file [ comer-example/README.md ] created
file [ comer-example/global/log.go ] created
file [ comer-example/utils/response/pages.go ] created
file [ comer-example/apps/apps.go ] created
file [ comer-example/components/captcha.go ] created
file [ comer-example/middlewares/CrosMiddleware.go ] created
file [ comer-example/apps/user/migrates/userlog.migrate.go ] created
file [ comer-example/.air.toml ] created
file [ comer-example/utils/mytime/week.go ] created
file [ comer-example/middlewares/RateLimitMiddleware.go ] created
file [ comer-example/apps/user/models/role.model.go ] created
file [ comer-example/configs/casbin.conf ] created
file [ comer-example/global/casbin.go ] created
file [ comer-example/router/router.go ] created
file [ comer-example/utils/maker/maker.go ] created
file [ comer-example/apps/user/models/user.model.go ] created
file [ comer-example/apps/user/repos/role.repo.go ] created
file [ comer-example/.vscode/settings.json ] created
file [ comer-example/cmd/init.go ] created
file [ comer-example/utils/request/http.go ] created
file [ comer-example/apps/common/handlers/captcha.handler.go ] created
file [ comer-example/main.go ] created
file [ comer-example/middlewares/JWTAuthMiddleware.go ] created
file [ comer-example/utils/request/pages.go ] created
file [ comer-example/utils/response/response.go ] created
file [ comer-example/utils/logger.go ] created
file [ comer-example/utils/utils.go ] created
file [ comer-example/apps/user/repos/userlog.repo.go ] created
file [ comer-example/apps/user/repos/userrole.repo.go ] created
file [ comer-example/global/config.go ] created
file [ comer-example/utils/copy/copy.go ] created
file [ comer-example/apps/user/services/role.service.go ] created
file [ comer-example/apps/user/services/userrole.service.go ] created
file [ comer-example/cmd/server.go ] created
file [ comer-example/components/redis.go ] created
file [ comer-example/apps/user/repos/user.repo.go ] created
file [ comer-example/middlewares/VcodeMiddleware.go ] created
file [ comer-example/utils/password/password.go ] created
file [ comer-example/utils/office/excel.go ] created
file [ comer-example/apps/user/migrates/userrole.migrate.go ] created
file [ comer-example/apps/user/migrates/role.migrate.go ] created
file [ comer-example/apps/user/migrates/user.migrate.go ] created
file [ comer-example/apps/user/services/userlog.service.go ] created
file [ comer-example/global/redis.go ] created
file [ comer-example/apps/swagger/router.go ] created
file [ comer-example/middlewares/token/jwttoken.go ] created
file [ comer-example/utils/format/format.go ] created
file [ comer-example/cmd/root.go ] created
file [ comer-example/middlewares/CasbinMiddleware.go ] created
file [ comer-example/global/mysql.go ] created
file [ comer-example/utils/slice/slice.go ] created
Do next:
1、cd  comer-example
2、change file（comer-example/configs/settings-local.yml）mysql and redis config
3、comer genapp --app=appName
4、air OR swag init && go mod tidy && go run . server

```

- 给项目添加app

```sh
cd comer-example
comer genapp --app=appName
#或者
comer genapp --app=user --swaggerTags='Oauth' --handler=auth --service=user --model=user,role
```
例如：
```sh
cd comer-example
$ comer genapp --app=student
Comer version  v1.1.6

_________
\_   ___ \   ____    _____    ____  _______
/    \  \/  /  _ \  /     \ _/ __ \ \_  __ \
\     \____(  <_> )|  Y Y  \\  ___/  |  | \/
 \______  / \____/ |__|_|  / \___  > |__|
                \/               \/      \/ v1.1.6, built with go1.20.2
dir [ ./apps ] existed
dir [ ./apps/student/handlers ] created
dir [ ./apps/student/migrates ] created
dir [ ./apps/student/models ] created
dir [ ./apps/student/repos ] created
dir [ ./apps/student/services ] created
file [ ./apps/student/models/student.model.go ] created
file [ ./apps/student/repos/student.repo.go ] created
file [ ./apps/student/services/student.service.go ] created
file [ ./apps/apps.go ] already exists
file [ ./apps/student/router.go ] created
file [ ./apps/student/handlers/student.handler.go ] created
file [ ./apps/student/migrates/student.migrate.go ] created
file [ ./apps/apps.go ] already exists
file [ ./apps/student/router.go ] already exists
comer genapp end.

```

- 安装swag

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

- 生成swagger文档

```sh
swag init
```

- 安装air

```sh
go install github.com/cosmtrek/air@latest
```

- 修改数据库配置(./configs/settings-local.yml)

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
  dsn: root:password@tcp(127.0.0.1:3306)/comer_project?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
  casbin: root:password@tcp(127.0.0.1:3306)/comer_project
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
  prefix: "comer_project:cache"

```

- 数据迁移:生成基本的数据库表

```sh
go run . migrate
```
例如：
```sh
$ go run . migrate
2023/06/26 14:57:32 Using config file: C:\Users\simpl\dev\golang\imoowi\comer-example\configs\settings-local.yml
2023/06/26 14:57:32 migrate start.
Connected to MySql!

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/student/migrates/student.migrate.go:26
[0.613ms] [rows:-] SELECT DATABASE()

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/student/migrates/student.migrate.go:26
[15.252ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'comer_project%' ORDER BY SCHEMA_NAME='comer_project' DESC,SCHEMA_NAME limit 1

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/student/migrates/student.migrate.go:26
[3.550ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'comer_project' AND table_name = 'students' AND table_type = 'BASE TABLE'

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/student/migrates/student.migrate.go:26
[29.227ms] [rows:0] CREATE TABLE `students` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`name` varchar(30) NOT NULL COMMENT '名称',PRIMARY KEY (`id`),INDEX `idx_students_deleted_at` (`deleted_at`))ENGINE=InnoDB,COMMENT='Student表'

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/role.migrate.go:23
[1.786ms] [rows:-] SELECT DATABASE()

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/role.migrate.go:23
[51.542ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'comer_project%' ORDER BY SCHEMA_NAME='comer_project' DESC,SCHEMA_NAME limit 1

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/role.migrate.go:23
[2.822ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'comer_project' AND table_name = 'roles' AND table_type = 'BASE TABLE'

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/role.migrate.go:23
[35.863ms] [rows:0] CREATE TABLE `roles` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`name` varchar(30) NOT NULL COMMENT '角色名',`level` tinyint(3) COMMENT '角色等级',PRIMARY KEY (`id`),INDEX `idx_roles_deleted_at` (`deleted_at`))ENGINE=InnoDB,COMMENT='角色表'

2023/06/26 14:57:32 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/user.migrate.go:23
[2.025ms] [rows:-] SELECT DATABASE()

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/user.migrate.go:23
[17.792ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'comer_project%' ORDER BY SCHEMA_NAME='comer_project' DESC,SCHEMA_NAME limit 1

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/user.migrate.go:23
[2.729ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'comer_project' AND table_name = 'users' AND table_type = 'BASE TABLE'

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/user.migrate.go:23
[38.228ms] [rows:0] CREATE TABLE `users` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`username` varchar(50) NOT NULL COMMENT '用户名',`password` varchar(255) NOT NULL,`salt` varchar(6) NOT NULL,`mp_openid` varchar(50),PRIMARY KEY (`id`),INDEX `idx_users_deleted_at` (`deleted_at`),UNIQUE INDEX `users_username` (`username`))ENGINE=InnoDB,COMMENT='用户表'

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userlog.migrate.go:23
[1.397ms] [rows:-] SELECT DATABASE()

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userlog.migrate.go:23
[18.770ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'comer_project%' ORDER BY SCHEMA_NAME='comer_project' DESC,SCHEMA_NAME limit 1

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userlog.migrate.go:23
[13.336ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'comer_project' AND table_name = 'user_logs' AND table_type = 'BASE TABLE'

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userlog.migrate.go:23
[36.882ms] [rows:0] CREATE TABLE `user_logs` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`user_id` bigint unsigned NOT NULL COMMENT '用户id',`log_type` tinyint(3) NOT NULL COMMENT '类型',`log_content` longtext COMMENT '记录内容',`ip` longtext COMMENT 'ip',PRIMARY KEY (`id`),INDEX `idx_user_logs_deleted_at` (`deleted_at`),INDEX `idx_user_logs_user_id` (`user_id`),INDEX `idx_user_logs_log_type` (`log_type`))ENGINE=InnoDB,COMMENT='用户行为记录表'

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userrole.migrate.go:23
[1.044ms] [rows:-] SELECT DATABASE()

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userrole.migrate.go:23
[17.370ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'comer_project%' ORDER BY SCHEMA_NAME='comer_project' DESC,SCHEMA_NAME limit 1

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userrole.migrate.go:23
[2.150ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'comer_project' AND table_name = 'user_roles' AND table_type = 'BASE TABLE'

2023/06/26 14:57:33 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/migrates/userrole.migrate.go:23
[36.904ms] [rows:0] CREATE TABLE `user_roles` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`user_id` bigint unsigned NOT NULL COMMENT '用户id',`role_id` bigint unsigned NOT NULL COMMENT '角色id',PRIMARY KEY (`id`),INDEX `idx_user_roles_deleted_at` (`deleted_at`),UNIQUE INDEX `idx_user_role_rel` (`user_id`,`role_id`))ENGINE=InnoDB,COMMENT='用户角色关系表'
2023/06/26 14:57:33 migrate end.


```
- 初始化数据库

```sh
go run . init
```
例如：
```sh
$ go run . init
2023/06/26 15:00:26 Using config file: C:\Users\simpl\dev\golang\imoowi\comer-example\configs\settings-local.yml
init called
2023/06/26 15:00:26 init start.
Connected to MySql!

2023/06/26 15:00:26 C:/Users/simpl/go/pkg/mod/github.com/casbin/gorm-adapter/v3@v3.18.0/adapter.go:413 SLOW SQL >= 200ms
[222.496ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'comer_project%' ORDER BY SCHEMA_NAME='comer_project' DESC,SCHEMA_NAME limit 1
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)


2023/06/26 15:00:26 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/repos/role.repo.go:57 record not found
[10.772ms] [rows:0] SELECT * FROM `roles` WHERE name='超级管理员' AND `roles`.`deleted_at` IS NULL ORDER BY `roles`.`id` LIMIT 1

2023/06/26 15:00:26 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/repos/role.repo.go:57 record not found
[2.152ms] [rows:0] SELECT * FROM `roles` WHERE name='超级管理员' AND `roles`.`deleted_at` IS NULL ORDER BY `roles`.`id` LIMIT 1

2023/06/26 15:00:26 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/repos/role.repo.go:62
[6.441ms] [rows:1] INSERT INTO `roles` (`created_at`,`updated_at`,`deleted_at`,`name`,`level`) VALUES ('2023-06-26 15:00:26.591','2023-06-26 15:00:26.591',NULL,'超级管理员','1')

2023/06/26 15:00:26 C:/Users/simpl/dev/golang/imoowi/comer-example/apps/user/repos/user.repo.go:52 record not found
[7.794ms] [rows:0] SELECT * FROM `users` WHERE id=0 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
2023/06/26 15:00:59 init end.
```

- 运行项目

```sh
air
#或者
go mod tidy
go run . server
```
例如：
```sh
$ air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ , built with Go

watching .
watching apps
watching apps\common
watching apps\common\handlers
watching apps\student
watching apps\student\handlers
watching apps\student\migrates
watching apps\student\models
watching apps\student\repos
watching apps\student\services
watching apps\swagger
watching apps\user
watching apps\user\handlers
watching apps\user\migrates
watching apps\user\models
watching apps\user\repos
watching apps\user\services
watching cmd
watching components
watching configs
watching docs
watching global
watching middlewares
watching middlewares\token
watching router
!exclude runtime
watching utils
watching utils\copy
watching utils\format
watching utils\maker
watching utils\myfile
watching utils\mytime
watching utils\office
watching utils\password
watching utils\request
watching utils\response
watching utils\slice
building...
2023/06/26 14:55:39 Generate swagger docs....
2023/06/26 14:55:39 Generate general API Info, search dir:./
2023/06/26 14:55:41 Generating response.PageList
2023/06/26 14:55:41 Generating response.Pages
2023/06/26 14:55:41 Generating models.UserLogin
2023/06/26 14:55:41 Generating models.UserChgPwd
2023/06/26 14:55:41 create docs.go at docs/docs.go
2023/06/26 14:55:41 create swagger.json at docs/swagger.json
2023/06/26 14:55:41 create swagger.yaml at docs/swagger.yaml
docs\docs.go has changed
running...
Connected to MySql!
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/student/students     --> github.com/imoowi/comer-example/apps/student/handlers.StudentPageList (8 handlers)
[GIN-debug] GET    /api/student/students/:id --> github.com/imoowi/comer-example/apps/student/handlers.StudentOne (8 handlers)
[GIN-debug] POST   /api/student/students     --> github.com/imoowi/comer-example/apps/student/handlers.StudentAdd (8 handlers)
[GIN-debug] PUT    /api/student/students/:id --> github.com/imoowi/comer-example/apps/student/handlers.StudentUpdate (8 handlers)
[GIN-debug] DELETE /api/student/students/:id --> github.com/imoowi/comer-example/apps/student/handlers.StudentDel (8 handlers)
[GIN-debug] GET    /api/common/captcha       --> github.com/imoowi/comer-example/apps/common/handlers.Captcha (6 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (6 handlers)
[GIN-debug] POST   /api/auth/login           --> github.com/imoowi/comer-example/apps/user/handlers.AuthLogin (7 handlers)
[GIN-debug] GET    /api/auth/logout          --> github.com/imoowi/comer-example/apps/user/handlers.AuthLogout (7 handlers)
[GIN-debug] POST   /api/auth/chgpwd          --> github.com/imoowi/comer-example/apps/user/handlers.AuthChgPwd (8 handlers)
[GIN-debug] GET    /api/casbins/allapi       --> github.com/imoowi/comer-example/router.InitRouter.func1 (6 handlers)
server port:  8000
API document address http://localhost:8000/swagger/index.html

```
- 访问接口文件：[http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)
![](assets/comer-swagger.png)
