package handler

import (
	"net/http"

	"github.com/spf13/cast"

	"bz.service.cloud.monitoring/common"
	params2 "bz.service.cloud.monitoring/internal/v1/params"
	"bz.service.cloud.monitoring/internal/v1/service"
	"bz.service.cloud.monitoring/utils"
	"github.com/labstack/echo"
)

// Login
func Login(c echo.Context) error {
	params := &params2.AdminLoginParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, common.Fail, msg, ""))
	}
	res, err := service.AdminLogin(params, c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, common.Success, cast.ToString(err), ""))
	}

	type loginResult struct {
		Token    string
		Username string
	}
	lr := &loginResult{
		Token:    res,
		Username: params.Username,
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, common.Success, "登陆成功", lr))
}
