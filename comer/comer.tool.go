package comer

import (
	"embed"
	"fmt"
	"os"
	"runtime"
	"text/template"
)

//go:embed templates/*
var tpl embed.FS

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
func (c *Comer) generateFrameworkDirByName(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(`目录[`, dirName, `]创建成功`)
	} else {
		fmt.Println(`目录[`, dirName, `] 已经存在，无需创建`)
	}
}

func (c *Comer) generateFrameworkFileByMap(fileName string, tplFileName string, tplData any) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("文件["+fileName+"]打开失败", err.Error())
			return
		}
		defer file.Close()

		t, err := template.ParseFiles(tplFileName)
		t2, err := template.ParseFS(tpl, tplFileName)
		fmt.Println(`t2=`, t2, ` tplfilename=`, tplFileName)
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
