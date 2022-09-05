package handler

import (
	"net/http"

	"bz.service.cloud.monitoring/server/internal/v1/params"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"github.com/labstack/echo"
)

// MachineList
func MachineList(c echo.Context) error {
	params := &params.MachineListParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	count, list, err := service.MachineList(params)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取机器列表失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功",
		utils.Resp.ResponsePagination(count, list)))
}

// AllMachine
func AllMachine(c echo.Context) error {
	machines, err := service.AllMachine()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取所有客户端机器失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", machines))
}

//// DeleteMachine
//func DeleteMachine(c echo.Context) error {
//
//}
