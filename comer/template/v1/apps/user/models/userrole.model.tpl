package models

import "{{.moduleName}}/components"

// 用户角色表
type UserRole struct {
	components.GormModel
	UserID uint `json:"user_id" form:"user_id" gorm:"column:user_id;not null;comment:用户id;uniqueIndex:idx_user_role_rel" binding:"required"`
	RoleId uint `json:"role_id" form:"role_id" gorm:"column:role_id;not null;comment:角色id;uniqueIndex:idx_user_role_rel" binding:"required"`
}

type UserRoleAdd struct {
	UserID uint `json:"user_id" form:"user_id"  binding:"required"`
	RoleId uint `json:"role_id" form:"role_id"  binding:"required"`
}
