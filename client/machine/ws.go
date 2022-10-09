package machine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/disk"

	mem2 "github.com/shirou/gopsutil/mem"

	_const "bz.service.cloud.monitoring/common/const"
	"github.com/shirou/gopsutil/cpu"
	"github.com/spf13/cast"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"

	"bz.service.cloud.monitoring/common/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type Client struct {
	Ip   string
	Conn *websocket.Conn
	Send chan *SystemMessage
}

type SystemMessage struct {
	Percent string `json:"percent"` // 使用率 cpu memory disk
}

var (
	WsClient *Client
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (c *Client) write(percent string) {
	m := &SystemMessage{
		Percent: percent,
	}
	res, _ := json.Marshal(m)
	err := c.Conn.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error("往服务器推送系统信息失败", zap.Error(err))
	}
}

// WsGetCpuSample
func WsGetCpuSample(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil) // 服务升级
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}

	WsClient = &Client{
		Ip:   utils.GetIP(),
		Conn: conn,
		Send: make(chan *SystemMessage, 100),
	}

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		ubzer.MLog.Error("获取cpu使用率失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取cpu使用率失败", ""))
	}
	cpuPercent := strconv.FormatInt(int64(percent[0]), 10)

	go WsClient.write(cast.ToString(cpuPercent))

	return nil
}

// WsGetMemSample
func WsGetMemSample(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}

	WsClient = &Client{
		Ip:   utils.GetIP(),
		Conn: conn,
		Send: make(chan *SystemMessage, 100),
	}

	mem, err := mem2.VirtualMemory()
	if err != nil {
		ubzer.MLog.Error("获取内存使用率失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取cpu使用率失败", ""))
	}
	memPercent := strconv.FormatInt(int64(mem.UsedPercent), 10)

	go WsClient.write(cast.ToString(memPercent))

	return nil
}

// WsGetDiskSample
func WsGetDiskSample(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}

	WsClient = &Client{
		Ip:   utils.GetIP(),
		Conn: conn,
		Send: make(chan *SystemMessage, 100),
	}

	parts, err := disk.Partitions(true)
	if err != nil {
		ubzer.MLog.Error("获取y硬盘使用率失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取y硬盘使用率失败", ""))
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		ubzer.MLog.Error("获取y硬盘使用率失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取y硬盘使用率失败", ""))
	}
	diskPercent := strconv.FormatInt(int64(diskInfo.UsedPercent), 10)

	go WsClient.write(cast.ToString(diskPercent))

	return nil
}

var exec string

// WsReceiveCom
func WsReceiveCom(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}
	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("从 websocket中读取服务端发送的指令失败"), zap.Error(err))
	}
	exec = string(content)
	err = utils.ExecCommand(exec)
	if err != nil {
		ubzer.MLog.Error("指令执行失败", zap.Error(err))
	}
	return nil
}
