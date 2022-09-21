package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bz.service.cloud.monitoring/server/config"

	"github.com/gorilla/websocket"

	"bz.service.cloud.monitoring/server/internal/v1/daos"

	"github.com/spf13/cast"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"

	"go.uber.org/zap"

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

	//bytes, err := request.Get("http://" + params.Ip + ":9093/c/sys/cpu")
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的cpu使用率失败", params.Ip), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	//}
	//res := &utils.Result{}

	// 由原来的http 升级为websocket
	dl := websocket.Dialer{}
	d := "ws://" + params.Ip + config.Config().ClientHttpBind + "/c/sys/cpu"
	conn, _, err := dl.Dial(d, nil)

	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的cpu使用率, websocket连接客户端失败", params.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}

	type Result struct {
		Percent string
	}

	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的cpu使用率, 从websocket中读取客户端推送信息失败", params.Ip),
			zap.Error(err))
	}
	res := &Result{}
	_ = json.Unmarshal(content, res)
	if res.Percent != "" {
		if utils.LessThenAndEqual("60", res.Percent) {
			err := daos.CreateMachineCheckRecord(params.Ip, "cpu", cast.ToInt(res.Percent))
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("记录服务器报警信息失败"))
			}
		}
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", cast.ToInt(res.Percent)))
}

// ClientMemPercent
func ClientMemPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	//bytes, err := request.Get("http://" + params.Ip + ":9093/c/sys/mem")
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的内存使用率失败", params.Ip), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	//}
	//res := &utils.Result{}

	// 由原来的http 升级为websocket
	dl := websocket.Dialer{}
	d := "ws://" + params.Ip + config.Config().ClientHttpBind + "/c/sys/mem"
	conn, _, err := dl.Dial(d, nil)

	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的内存使用率, websocket连接客户端失败", params.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}
	type Result struct {
		Percent string
	}
	defer conn.Close()

	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的内存使用率, 从websocket中读取客户端推送信息失败", params.Ip),
			zap.Error(err))
	}
	res := &Result{}
	_ = json.Unmarshal(content, res)
	if res.Percent != "" {
		if utils.LessThenAndEqual("60", res.Percent) {
			err := daos.CreateMachineCheckRecord(params.Ip, "mem", cast.ToInt(res.Percent))
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("记录服务器报警信息失败"))
			}
		}
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", cast.ToInt(res.Percent)))
}

// ClientDiskPercent
func ClientDiskPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	//bytes, err := request.Get("http://" + params.Ip + ":9093/c/sys/disk")
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的硬盘使用率失败", params.Ip), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	//}
	//res := &utils.Result{}

	// 由原来的http 升级为websocket
	dl := websocket.Dialer{}
	d := "ws://" + params.Ip + config.Config().ClientHttpBind + "/c/sys/disk"
	conn, _, err := dl.Dial(d, nil)

	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的磁盘使用率, websocket连接客户端失败", params.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}
	type Result struct {
		Percent string
	}

	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的磁盘使用率, 从websocket中读取客户端推送信息失败", params.Ip),
			zap.Error(err))
	}
	res := &Result{}
	_ = json.Unmarshal(content, res)
	if res.Percent != "" {
		if utils.LessThenAndEqual("60", res.Percent) {
			err := daos.CreateMachineCheckRecord(params.Ip, "cpu", cast.ToInt(res.Percent))
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("记录服务器报警信息失败"))
			}
		}
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", cast.ToInt(res.Percent)))
}
