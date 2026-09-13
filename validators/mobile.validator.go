package validators

import (
	"github.com/go-playground/validator/v10"
)

// 手机号正则表达式
var MobileRegRuler string = `^1[3456789]{1}\d{9}$`
var mobile validator.Func = func(fl validator.FieldLevel) bool {
	return Matched(fl, MobileRegRuler)
}

func init() {
	RegisterValidators(map[string]ValidatorFunc{`mobile`: ValidatorFunc(mobile)})
}
