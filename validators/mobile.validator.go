package validators

import (
	"github.com/go-playground/validator/v10"
)

// 手机号正则表达式
var MobileRegRuler string = `^1[3456789]{1}\d{9}$`
var mobile validator.Func = func(fl validator.FieldLevel) bool {
	/*
		matched := true
		data, ok := fl.Field().Interface().(string)
		if ok {
			if data != `` {
				reg := regexp.MustCompile(MobileRegRuler)
				matched = reg.MatchString(data)
			}
		}
		return matched
		//*/

	return Matched(fl, MobileRegRuler)
}

func init() {
	RegisterValidators(map[string]ValidatorFunc{`mobile`: ValidatorFunc(mobile)})
}
