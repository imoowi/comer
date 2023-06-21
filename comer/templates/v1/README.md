# {{.moduleName}}
## 一、⭐！强制规则
- 数据库选择MySql
- 必须使用迁移：数据库字段的修改必须通过gorm的数据迁移来实现,禁止手动改数据库，数据迁移参考文档[gorm迁移](https://gorm.io/zh_CN/docs/migration.html)
- 必须要有/app/[appname]/migrates/[modelName].migrate.go文件
- 运行项目前，必须在项目根目录执行以下数据迁移命令
```sh
cd project_root
go run . migrate
```

## 二、代码生成器
- 命令如下
```sh
#--app=[模块名，必须输入] --swaggerTags=[接口文档Tags] --handler=[控制器名，默认为模块名] --service=[服务名，默认为控制器名] --model=[模型名1,模型名2,模型名...;默认取第一个生成service]
comer genapp --app=user --swaggerTags='Oauth' --handler=auth --service=user --model=user

```

## 三、本系统脚手架使用步骤
### 步骤一、生成模块代码
```sh
comer genapp --module=user
```
### 步骤二、修改数据表定义文件"xxx.model.go"
```go
//将以下内容

package models

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"column:username;type:varchar(50);not null" binding:"required"`
}

//修改为
package models

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"column:username;type:varchar(50);not null" binding:"required"`  //用户名
	Passwd   string `json:"password" form:"password" gorm:"column:password;type:varchar(255);not null" binding:"required"` //密码
	Salt     string `json:"salt" form:"salt" gorm:"column:salt;type:varchar(6);not null"  `                                //盐
}

```
### 步骤三、数据迁移：生成目标表（只新增，不做删除）
```sh
go run . migrate -c ./config/settings-local.yml
```
### 步骤四、初始化系统
```sh
go run . init -c ./config/settings-local.yml
```
### 步骤五、生成swagger文档
```sh
swag init
```
### 步骤六、查看文档：运行以下命令，然后访问[http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)
```sh
go run . server
#或者
air
```
