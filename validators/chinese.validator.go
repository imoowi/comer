package validators

import (
	"github.com/go-playground/validator/v10"
)

// 中文正则表达式
var ChineseRegRuler string = "^[\u4e00-\u9fa5]+$"
var chinese validator.Func = func(fl validator.FieldLevel) bool {
	return Matched(fl, ChineseRegRuler)
}

func init() {
	RegisterValidators(map[string]ValidatorFunc{`chinese`: ValidatorFunc(chinese)})
}
