package handler

import (
	"net/http"

	_const "bz.service.cloud.monitoring/common/const"

	"github.com/spf13/cast"

	params2 "bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"bz.service.cloud.monitoring/server/utils"
	"github.com/labstack/echo"
)

// Login
func Login(c echo.Context) error {
	params := &params2.AdminLoginParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	res, err := service.AdminLogin(params, c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, cast.ToString(err), ""))
	}

	type loginResult struct {
		Token    string
		Username string
	}
	lr := &loginResult{
		Token:    res,
		Username: params.Username,
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "登陆成功", lr))
}
