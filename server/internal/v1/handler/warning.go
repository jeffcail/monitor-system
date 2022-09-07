package handler

import (
	"net/http"

	"github.com/spf13/cast"

	"bz.service.cloud.monitoring/server/internal/v1/service"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/params"
	"github.com/labstack/echo"
)

// ServeCheckRecordList
func ServeCheckRecordList(c echo.Context) error {
	count, list, err := service.ServeCheckRecordList()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", utils.Resp.ResponsePagination(
		count, list)))
}

// IgnoreServeCheckRecord
func IgnoreServeCheckRecord(c echo.Context) error {
	param := &params.IgnoreServeCheckRecordParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.IgnoreServeCheckRecord(param.ServeId, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "忽略成功", ""))
}

// MachineCheckList
func MachineCheckList(c echo.Context) error {
	count, list, err := service.MachineCheckList()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", utils.Resp.ResponsePagination(
		count, list)))
}

// IgnoreMachineCheckRecord
func IgnoreMachineCheckRecord(c echo.Context) error {
	param := &params.IgnoreMachineParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.IgnoreMachineCheckRecord(param.Id, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "忽略成功", ""))
}
