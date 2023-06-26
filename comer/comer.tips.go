package comer

import "fmt"

func (c *Comer) showTips() string {
	fmt.Println(`Do next:`)
	fmt.Println(`1、cd  ` + c.path)
	fmt.Println(`2、change file（` + c.path + `/configs/settings-local.yml）mysql and redis config`)
	fmt.Println(`3、comer genapp --app=appName`)
	fmt.Println(`4、air OR swag init && go mod tidy && go run . server`)
	return `showTips called`
}
