/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package components

import (
	"image/color"
	"strings"

	"github.com/mojocn/base64Captcha"
)

var store base64Captcha.Store = base64Captcha.DefaultMemStore

type POSTVcode struct {
	Id    string `json:"captcha_id" form:"captcha_id" binding:"required"`     //验证码id
	Vcode string `json:"captcha_code" form:"captcha_code" binding:"required"` //验证码字符串
}

// SetCaptchaStore 设置验证码存储（默认内存存储），可传入 RedisCaptchaStore。
func SetCaptchaStore(s base64Captcha.Store) {
	store = s
}

// GenCaptcha 根据 driverType 生成验证码；未知类型回退到数字验证码。
func GenCaptcha(driverType string) (id string, b64s string, err error) {
	captcha := base64Captcha.NewCaptcha(captchaDriver(driverType), store)
	id, b64s, _, err = captcha.Generate()
	return
}

func VerifyCaptcha(id string, vcode string) bool {
	return store.Verify(id, vcode, true)
}

// captchaDriver 按 driverType 返回对应的验证码 driver。
func captchaDriver(driverType string) base64Captcha.Driver {
	switch strings.ToLower(strings.TrimSpace(driverType)) {
	case `string`:
		return base64Captcha.NewDriverString(80, 240, 0, 0, 4, `1234567890abcdefghijklmnopqrstuvwxyz`, randLightColor(), base64Captcha.DefaultEmbeddedFonts, []string{`RitaSmith.ttf`})
	case `math`:
		return base64Captcha.NewDriverMath(80, 240, 0, 0, randLightColor(), base64Captcha.DefaultEmbeddedFonts, []string{`RitaSmith.ttf`})
	case `chinese`:
		return base64Captcha.NewDriverChinese(80, 240, 0, 0, 4, `你好世界吃葡萄不吐葡萄皮`, randLightColor(), base64Captcha.DefaultEmbeddedFonts, []string{`wqy-microhei.ttc`})
	case `audio`:
		return base64Captcha.NewDriverAudio(6, `en`)
	default:
		return base64Captcha.NewDriverDigit(70, 130, 4, 0.8, 100)
	}
}

// randLightColor 返回随机浅色背景。
func randLightColor() *color.RGBA {
	c := base64Captcha.RandLightColor()
	return &c
}
