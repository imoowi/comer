package comer

import "fmt"

func (c *Comer) showTips() {
	fmt.Printf("接下来，请执行以下操作:\n1、修改配置文件（%s/configs/settings-local.yml）数据库和redis的配置\n2、cd %s \n3、air\n", c.path, c.path)
}
