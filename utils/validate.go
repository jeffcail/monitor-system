package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"

	validate2 "bz.service.cloud.monitoring/pkg/validate"
)

var msg string

// ValidateParam
func ValidateParam(c echo.Context, params interface{}) string {
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
