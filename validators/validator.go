package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidatorFunc validator.Func

var Validators = map[string]ValidatorFunc{}

func RegisterValidators(v map[string]ValidatorFunc) {
	for k, val := range v {
		Validators[k] = val
	}
}
func Matched(fl validator.FieldLevel, regRuler string) bool {
	matched := false
	data, ok := fl.Field().Interface().(string)
	if ok {
		if data != `` {
			reg := regexp.MustCompile(regRuler)
			matched = reg.MatchString(data)
		}
	}
	return matched
}
func InitValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for k, val := range Validators {
			v.RegisterValidation(k, validator.Func(val))
		}
	}
}
