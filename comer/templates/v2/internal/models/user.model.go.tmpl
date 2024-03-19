package models

import "github.com/imoowi/comer/components"

type User struct {
	components.GormModel
	UserBase
}
type UserBase struct {
	Username string `json:"username" form:"username" gorm:"column:username;type:varchar(50);not null;comment:用户名;uniqueIndex:users_username" binding:"required"` //用户名
	Passwd   string `json:"password" form:"password" gorm:"column:password;type:varchar(255);not null" binding:"required"`                                       //密码
	Salt     string `json:"salt" form:"salt" gorm:"column:salt;type:varchar(6);not null"  `                                                                      //盐
	MpOpenid string `json:"mp_openid" form:"mp_openid" gorm:"column:mp_openid;type:varchar(50)"  `
}

// IModel.GetID实现
func (m *User) GetID() uint {
	return m.ID
}
func (m *User) SetId(id uint) {
	m.ID = id
}

// IModel.TableName实现
func (m *User) TableName() string {
	return `user` + `s`
}

// 登录表单
type UserLogin struct {
	Username string `json:"username" form:"username" binding:"required"` //用户名
	Passwd   string `json:"password" form:"password" binding:"required"` //密码
	components.POSTVcode
}

// 添加用户
type UserAdd struct {
	RoleId   uint   `json:"role_id" form:"role_id" binding:"required"`   //角色id
	Username string `json:"username" form:"username" binding:"required"` //用户名
	Passwd   string `json:"password" form:"password" binding:"required"` //密码
}

// 改密
type UserChgPwd struct {
	UserId     uint   `json:"-" gorm:"id"`                    //用户id
	OriginPwd  string `json:"origin_pwd" binding:"required"`  //原始密码
	NewPwd     string `json:"new_pwd" binding:"required"`     //新密码
	ConfirmPwd string `json:"confirm_pwd" binding:"required"` //确认密码
	components.POSTVcode
}
