package validators

import (
	"github.com/go-playground/validator/v10"
)

// 中文正则表达式
var EnglishRegRuler string = `^[a-zA-Z0-9_\-]+$`
var english validator.Func = func(fl validator.FieldLevel) bool {
	return Matched(fl, EnglishRegRuler)
}

func init() {
	RegisterValidators(map[string]ValidatorFunc{`english`: ValidatorFunc(english)})
}
