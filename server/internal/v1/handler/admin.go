package handler

import (
	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"bz.service.cloud.monitoring/server/utils"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// controller层
// 管理员注册
func AdminRegister(e echo.Context) error {
	//实例化一个对象用户接受注册信息
	params := &params.AdminParam{}
	err := e.Bind(params)
	if err != nil {
		fmt.Errorf(err.Error())
		return e.JSON(http.StatusBadRequest, "参数错误")
	}

	res, _ := service.AdminRegister(params, e.Request().URL.Path, e.Request().Method)
	if !res {
		return e.JSON(http.StatusBadRequest, "")
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Fail, "用户注册失败", ""))
	}
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "管理员用户注册成功！", ""))
}
