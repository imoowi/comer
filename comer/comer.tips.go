/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import "fmt"

func (c *Comer) showTips() string {
	fmt.Println(`Do next:`)
	fmt.Println(`1、cd  ` + c.path)
	fmt.Println(`2、change file（` + c.path + `/configs/settings-local.yml）mysql and redis config`)
	fmt.Println(`3、comer add -a=appName -c=handlerName -w=swaggerTagsName -s=serviceName -m=modelName1`)
	fmt.Println(`4、air OR swag init && go mod tidy && go run . server`)
	return `showTips called`
}
