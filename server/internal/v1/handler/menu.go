package handler

import (
	"net/http"

	_const "github.com/c/server-monitoring/common/const"
	"github.com/c/server-monitoring/common/utils"
	"github.com/c/server-monitoring/server/internal/v1/service"
	"github.com/labstack/echo"
)

// MenusList
func MenusList(c echo.Context) error {
	list := service.MenuList()
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", list))
}
