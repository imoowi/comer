package components

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

func GenCaptcha(driverType string) (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(70, 130, 4, 0.8, 100)
	// 生成base64图片
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, err = captcha.Generate()
	return
}

func VerifyCaptcha(id string, vcode string) bool {
	return store.Verify(id, vcode, true)
}
