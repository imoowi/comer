package comer

import "fmt"

func (c *Comer) showTips() {
	fmt.Println(`接下来，请执行以下操作:`)
	fmt.Println(`1、修改配置文件（` + c.path + `/configs/settings-local.yml）数据库和redis的配置`)
	fmt.Println(`2、cd  ` + c.path)
	fmt.Println(`3、air 或者 swag init && go mod tidy && go run .`)
}
