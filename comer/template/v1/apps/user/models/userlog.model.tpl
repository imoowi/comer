package models

import "{{.moduleName}}/components"

type USER_LOG_TYPE int

const (
	USER_LOG_TYPE_NULL        USER_LOG_TYPE = iota
	USER_LOG_TYPE_USER_LOGIN                //登录
	USER_LOG_TYPE_USER_LOGOUT               //退出
	USER_LOG_TYPE_USER_CHGPWD               //改密

	USER_LOG_TYPE_USER_ADD    //添加用户
	USER_LOG_TYPE_USER_UPDATE //修改用户
	USER_LOG_TYPE_USER_DELETE //删除用户

	USER_LOG_TYPE_ROLE_ADD    //添加角色
	USER_LOG_TYPE_ROLE_UPDATE //修改角色
	USER_LOG_TYPE_ROLE_DELETE //删除角色

)

// 用户行为记录表
// `gorm:"comment:用户行为记录表"`
type UserLog struct {
	components.GormModel
	UserID     uint          `json:"user_id" form:"user_id" gorm:"column:user_id;not null;comment:用户id;index" binding:"required"`                  //用户id
	LogType    USER_LOG_TYPE `json:"log_type" form:"log_type" gorm:"column:log_type;type:tinyint(3);not null;comment:类型;index" binding:"required"` //log类型
	LogContent string        `json:"log_content" form:"log_content" gorm:"column:log_content;comment:记录内容" binding:"required"`                     //log内容
	IP         string        `json:"ip" form:"ip" gorm:"column:ip;comment:ip"`
}