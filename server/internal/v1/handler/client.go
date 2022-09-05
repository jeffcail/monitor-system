package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/request"
	"bz.service.cloud.monitoring/common/ubzer"
	params2 "bz.service.cloud.monitoring/server/internal/v1/params"
	"github.com/labstack/echo"
)

// ClientCpuPercent
func ClientCpuPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	bytes, err := request.Get("http://" + params.Ip + ":9093/c/sys/cpu")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的cpu使用率失败", params.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}

	res := &utils.Result{}
	_ = json.Unmarshal(bytes, res)
	if res.Code == 2000 {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", res.Data))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
}

// ClientMemPercent
func ClientMemPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	bytes, err := request.Get("http://" + params.Ip + ":9093/c/sys/mem")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的内存使用率失败", params.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}
	res := &utils.Result{}
	_ = json.Unmarshal(bytes, res)
	if res.Code == 2000 {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", res.Data))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
}

// ClientDiskPercent
func ClientDiskPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	bytes, err := request.Get("http://" + params.Ip + ":9093/c/sys/disk")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的硬盘使用率失败", params.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}
	res := &utils.Result{}
	_ = json.Unmarshal(bytes, res)
	if res.Code == 2000 {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", res.Data))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
}
