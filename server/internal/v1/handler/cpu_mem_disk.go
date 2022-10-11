package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bz.service.cloud.monitoring/common/db"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type CpuMemDisk struct {
	Con *websocket.Conn
}

var (
	Cmd    *CpuMemDisk
	upgrad = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ReadCpuMemDisk struct {
	Ip      string `json:"ip"`
	Percent string `json:"percent"`
}

func (this *CpuMemDisk) read() *ReadCpuMemDisk {
	_, content, err := this.Con.ReadMessage()
	if err != nil {
		ubzer.MLog.Error("接收客户端推送的系统信息失败")
	}
	cmdInfo := &ReadCpuMemDisk{}
	_ = json.Unmarshal(content, cmdInfo)
	return cmdInfo
}

// ClientCpu
func ClientCpu(c echo.Context) error {
	conn, err := upgrad.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}
	Cmd = &CpuMemDisk{
		Con: conn,
	}

	info := Cmd.read()
	fmt.Printf("==== ip: %v cpu: %v\n", info.Ip, info.Percent)
	lPushList(info.Ip+"-cpu", info.Percent)
	return nil
}

// ClientMemory
func ClientMemory(c echo.Context) error {
	conn, err := upgrad.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}
	Cmd = &CpuMemDisk{
		Con: conn,
	}
	info := Cmd.read()
	fmt.Printf("==== ip: %v memory: %v\n", info.Ip, info.Percent)
	lPushList(info.Ip+"-mem", info.Percent)
	return nil
}

// ClientDisk
func ClientDisk(c echo.Context) error {
	conn, err := upgrad.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}
	Cmd = &CpuMemDisk{
		Con: conn,
	}
	info := Cmd.read()
	fmt.Printf("==== ip: %v disk: %v\n", info.Ip, info.Percent)
	lPushList(info.Ip+"-disk", info.Percent)
	return nil
}

func lPushList(key, percent string) {
	lens, err := db.Rc.LLen(key).Result()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("计算客户端: %v redis list的长度失败", key), zap.Error(err))
		return
	}
	if lens >= 120 {
		s := db.Rc.LPop(key).String()
		if s != "" {
			err := db.Rc.RPush(key, percent).Err()
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("往redis list 推送客户端: %v 系统cpu、内存、磁盘使用率失败", key), zap.Error(err))
				return
			}
		}
	} else {
		err := db.Rc.RPush(key, percent).Err()
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("往redis list 推送客户端: %v 系统cpu、内存、磁盘使用率失败", key), zap.Error(err))
			return
		}
	}
}
