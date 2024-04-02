package validators

import (
	"github.com/go-playground/validator/v10"
)

// 身份证正则表达式
var IDCardRegRuler string = `^[1-9]\d{5}(18|19|20|(3\d))\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
var idcard validator.Func = func(fl validator.FieldLevel) bool {
	return Matched(fl, IDCardRegRuler)
}

func init() {
	RegisterValidators(map[string]ValidatorFunc{`idcard`: ValidatorFunc(idcard)})
}
