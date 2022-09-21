package machine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"bz.service.cloud.monitoring/client/config"
	"bz.service.cloud.monitoring/client/models"
	"bz.service.cloud.monitoring/common/db"

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
	Send chan *Message
}

type Message struct {
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
	m := &Message{
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
		Send: make(chan *Message, 100),
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
		Send: make(chan *Message, 100),
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
		Send: make(chan *Message, 100),
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

// WsUpgradeClient
func WsUpgradeClient(c echo.Context) (err error) {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return
	}
	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("客户端升级,从 websocket中读取服务端发送的升级指令失败"), zap.Error(err))
	}
	con := string(content)
	if con == "ok" {
		ubzer.MLog.Info(fmt.Sprintf("收到的升级指令: %v", con))
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 ============="))
		ip := utils.GetIP()
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 拿到ip: %v =============", ip))

		m := &models.MonUpgradeMachineRecord{}
		_, err := db.Mysql.Where("machine_ip = ?", ip).Desc("id").Limit(1).Get(m)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("定时从数据库中检测客户端最新版本失败 ip: %v", ip), zap.Error(err))
			return err
		}
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 拿到最新的升级记录: %v =============", m))
		src, err := os.Open("version.txt")
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("定时检测更新客户端版本 打开版本文件读取版本号失败"), zap.Error(err))
			return err
		}
		defer src.Close()

		content, err := io.ReadAll(src)
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 当前的版本为: %v =============", string(content)))

		if utils.LessThenAndEqual(m.UpgradeVersion, string(content)) {
			ubzer.MLog.Error(fmt.Sprintf("============ 要升级的版本小于: %v 当前的版本: %v =============", string(content),
				m.UpgradeVersion))
			return err
		}
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 要升级的版本为: %v =============", m.UpgradeVersion))

		file, err := os.OpenFile("version.txt", os.O_RDWR|os.O_TRUNC, 0766)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("客户端版本升级打开版本号文件失败"), zap.Error(err))
		}
		defer file.Close()

		str := m.UpgradeVersion
		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(str)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("版本号写入缓存失败"), zap.Error(err))
		}
		err = writer.Flush()
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("版本号写入文件失败"), zap.Error(err))
		}

		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 版本写入版本控制文件成功 ============="))

		pn := m.PackageName + ":" + m.UpgradeVersion
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 要升级的包名: %v =============", pn))

		exec := "wget " + config.Config().GoFileServe + "/api/dl?file=" + pn
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 下载要升级的包: %v =============", exec))

		err = utils.ExecCommand(exec)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("定时检测更新客户端版本 下载对应的版本失败"), zap.Error(err))
		}

		utils.ExecCommand("mv dl?file=" + pn + " " + pn)
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 重命名要升级的包 重命名之后包的名字: %v =============", pn))

		utils.ExecCommand("chmod +x " + pn)
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 给包赋予可执行的权限成功，开始杀死原进程: %v =============", pn))

		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 准备启动客户端程序 ============="))
		err = utils.ExecCommand("/etc/systemd/system/systemctl stop client-monitor.service")
		err = utils.ExecCommand("./" + pn + " uninstall")
		err = utils.ExecCommand("./" + pn + " install")
		err = utils.ExecCommand("/etc/systemd/system/systemctl start client-monitor.service")
		if err != nil {
			ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 启动客户端程序失败 ============="), zap.Error(err))
		} else {
			ubzer.MLog.Info(fmt.Sprintf("定时检测更新客户端版本成功, 原版本号为: %v 更新成功的版本号为: %v", string(content),
				m.UpgradeVersion))
		}
	}
	return
}
