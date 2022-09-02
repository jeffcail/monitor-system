package handler

import (
	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// controller层
// 管理员创建
func AdminRegister(e echo.Context) error {
	//实例化一个对象接受信息
	params := &params.AdminParam{}
	err := e.Bind(params)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return e.JSON(http.StatusBadRequest, "参数错误")
	}

	res, _ := service.AdminRegister(params, e.Request().URL.Path, e.Request().Method)
	if !res {

		return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Fail, "用户创建失败", ""))
	}
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "管理员用户创建成功！", ""))
}

// 查询
func SelectAdmin(e echo.Context) error {
	// 接参
	params := &params.SelAdminParam{}
	err := e.Bind(params)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return e.JSON(http.StatusBadRequest, "参数错误")
	}
	// 请求service
	total, list, err := int64(0), make([]*string, 0), errors.New("")
	// 返回
	data := utils.Resp.ResponsePagination(total, list)
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "管理员信息更改成功！", data))
}
