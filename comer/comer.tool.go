/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"os"
	"text/template"
)

func NewComer() *Comer {
	return &Comer{}
}
func (c *Comer) generateDirByName(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(`dir [`, dirName, `] created`)
	} else {
		fmt.Println(`dir [`, dirName, `] existed`)
	}
}

func (c *Comer) generateFileByMap(fileName string, tplFileName string, tplData any, customeTpl bool) {
	// fmt.Println(`fileName=`, fileName, `tplFileName=`, tplFileName)
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("file ["+fileName+"] open failed", err.Error())
			return
		}
		defer file.Close()
		var t *template.Template
		if customeTpl {
			t, err = template.ParseFiles(tplFileName)
		} else {
			t, err = template.ParseFS(tplLocal, tplFileName)
		}
		// t, err := template.ParseFiles(tplFileName)

		if err != nil {
			fmt.Println(`err:`, err.Error())
			return
		}
		err = t.Execute(file, tplData)
		if err != nil {
			fmt.Println(`err=`, err.Error())
		}
		fmt.Println(`file [`, fileName, `] created`)
	} else {
		fmt.Println(`file [`, fileName, `] already exists`)
	}
}

func (c *Comer) showLogo() {

	fmt.Printf(`
_________                                   
\_   ___ \   ____    _____    ____  _______ 
/    \  \/  /  _ \  /     \ _/ __ \ \_  __ \
\     \____(  <_> )|  Y Y  \\  ___/  |  | \/
 \______  / \____/ |__|_|  / \___  > |__|   
		\/               \/      \/ %s, built with %s
`, c.Version(), c.goVersion())
}
