/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package components

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

type POSTVcode struct {
	Id    string `json:"captcha_id" form:"captcha_id" binding:"required"`     //验证码id
	Vcode string `json:"captcha_code" form:"captcha_code" binding:"required"` //验证码字符串
}

func GenCaptcha(driverType string) (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(70, 130, 4, 0.8, 100)
	// 生成base64图片
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, _, err = captcha.Generate()
	return
}

func VerifyCaptcha(id string, vcode string) bool {
	return store.Verify(id, vcode, true)
}
