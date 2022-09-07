package handler

import (
	"net/http"

	"github.com/spf13/cast"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"

	"bz.service.cloud.monitoring/common/request"

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

// SendCom
func SendCom(c echo.Context) error {
	param := &params.SendComParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	h := make(map[string]string)
	p := make(map[string]interface{})
	p["content"] = param.Content
	bytes, err := request.GetParams("http://"+param.Ip+":9093/c/machine/receive/com", h, p)
	if err != nil {
		ubzer.MLog.Error("向客户端发送指令请求异常", zap.Error(err))
	}
	if string(bytes) != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(bytes), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "指令发送成功", ""))
}

//// DeleteMachine
//func DeleteMachine(c echo.Context) error {
//
//}
