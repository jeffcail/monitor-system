package handler

import (
	"net/http"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"github.com/labstack/echo"
)

// MachineList
func MachineList(c echo.Context) error {
	list, err := service.MachineList()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取机器列表失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", list))
}
