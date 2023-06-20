package models

import (
	"{{.moduleName}}/components"
)

type ROLE_LEVEL int

const (
	ROLE_LEVEL_NULL   ROLE_LEVEL = iota
	ROLE_LEVEL_SUPER             //超管
	ROLE_LEVEL_NORMAL            //普通用户
)

// 角色表
type Role struct {
	components.GormModel
	Name  string     `json:"name" form:"name" gorm:"column:name;type:varchar(30);not null;comment:角色名" binding:"required"`
	Level ROLE_LEVEL `json:"level" form:"level" gorm:"column:level;type:tinyint(3);comment:角色等级"`
}

type RoleAdd struct {
	Name  string     `json:"name" binding:"required"`  //角色名
	Level ROLE_LEVEL `json:"level" binding:"required"` //角色等级，1：超管；2：普通管理员
}
