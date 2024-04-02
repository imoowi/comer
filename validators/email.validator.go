package validators

import (
	"github.com/go-playground/validator/v10"
)

// 邮箱正则表达式
var EmailRegRuler string = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
var email validator.Func = func(fl validator.FieldLevel) bool {
	return Matched(fl, EmailRegRuler)
}

func init() {
	RegisterValidators(map[string]ValidatorFunc{`email`: ValidatorFunc(email)})
}
