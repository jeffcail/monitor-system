package utils

import (
	validate2 "github.com/c/monitor-system/common/validate"
	"github.com/go-playground/validator/v10"
)

var msg string

// ValidateParam
func ValidateParam(params interface{}) string {
	translator, validate := validate2.Bv.BindValidate()
	err := validate.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			msg = err.Translate(translator)
		}
		return msg
	}
	return ""
}
