/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package models

import "{{.ModuleName}}/components"

// {{.ModelName}}表
type {{.ModelName}} struct {
	components.GormModel
	Name string `gorm:"column:name;type:varchar(30);not null;comment:名称"`
}
