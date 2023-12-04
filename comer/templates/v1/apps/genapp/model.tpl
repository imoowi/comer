/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package models

import (
	"{{.ModuleName}}/components"
	"{{.ModuleName}}/utils/request"
)

// {{.ModelName}}表
type {{.ModelName}} struct {
	components.GormModel
	Name string `gorm:"column:name;type:varchar(30);not null;comment:名称"`
}

//新增
type {{.ModelName}}Add struct {
	Name string `json:"name" form:"name"` //名称
}

//完全更新
type {{.ModelName}}Update struct {
	Name string `json:"name" form:"name"` //名称
}

//部分更新
type {{.ModelName}}PatchUpdate struct {
	Name string `json:"name" form:"name"` //名称
}

//查询model
type {{.ModelName}}Query struct {
	request.PageList
	Key string `json:"key" form:"key"`
}