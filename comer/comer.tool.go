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

func (c *Comer) generateDirByName(dirName string) error {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dirName, 0755); err != nil {
			return err
		}
		fmt.Println(`dir [`, dirName, `] created`)
	} else if err != nil {
		return err
	} else {
		fmt.Println(`dir [`, dirName, `] existed`)
	}
	return nil
}

func (c *Comer) generateFileByMap(fileName string, tplFileName string, tplData any, customeTpl bool) error {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Println(`file [`, fileName, `] already exists`)
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("file [%s] open failed: %w", fileName, err)
	}
	defer file.Close()

	var t *template.Template
	if customeTpl {
		t, err = template.ParseFiles(tplFileName)
	} else {
		t, err = template.ParseFS(tplLocal, tplFileName)
	}
	if err != nil {
		return fmt.Errorf("parse template [%s] failed: %w", tplFileName, err)
	}
	if err = t.Execute(file, tplData); err != nil {
		return fmt.Errorf("execute template [%s] failed: %w", tplFileName, err)
	}
	fmt.Println(`file [`, fileName, `] created`)
	return nil
}

func (c *Comer) showLogo() {
	fmt.Println("Comer version ", c.Version())

	fmt.Printf(`
_________
\_   ___ \   ____    _____    ____  _______
/    \  \/  /  _ \  /     \ _/ __ \ \_  __ \
\     \____(  <_> )|  Y Y  \\  ___/  |  | \/
 \______  / \____/ |__|_|  / \___  > |__|
		\/               \/      \/ %s, built with %s
`, c.Version(), c.goVersion())
}
