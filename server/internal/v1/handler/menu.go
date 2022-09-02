package handler

import (
	"net/http"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"github.com/labstack/echo"
)

// MenusList
func MenusList(c echo.Context) error {
	list := service.MenuList()
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", list))
}
