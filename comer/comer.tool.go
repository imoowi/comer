package comer

import (
	"fmt"
	"os"
	"runtime"
	"text/template"
)

func (c *Comer) Version() string {
	c.version = `v1.0`
	fmt.Println(`Comer version `, c.version)
	return c.version
}
func (c *Comer) goVersion() string {
	return runtime.Version()
}
func NewComer() *Comer {
	return &Comer{}
}
func (c *Comer) generateFrameworkDirByName(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err := os.MkdirAll(name, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(`目录[`, name, `]创建成功`)
	} else {
		fmt.Println(`目录[`, name, `] 已经存在，无需创建`)
	}
}

func (c *Comer) generateFrameworkFileByMap(fileName string, tplName string, tplData any) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("文件["+fileName+"]打开失败", err.Error())
			return
		}
		defer file.Close()

		t, err := template.ParseFiles(tplName)
		if err != nil {
			fmt.Println(`err:`, err.Error())
			return
		}
		err = t.Execute(file, tplData)
		if err != nil {
			fmt.Println(`err=`, err.Error())
		}
		fmt.Println(`文件[`, fileName, `] 创建成功`)
	} else {
		fmt.Println(`文件[`, fileName, `] 已经存在，无需创建`)
	}
}
