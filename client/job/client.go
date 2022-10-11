package job

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/disk"

	mem2 "github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"

	"bz.service.cloud.monitoring/common/utils"

	"bz.service.cloud.monitoring/common/ubzer"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// 每30秒向服务端推送 客户端 cpu使用率
// PushClientCpuPercent
func PushClientCpuPercent() {
	ip := utils.GetIP()
	//d := "ws://" + config.Config().GoFileServe + "/client/cpu"
	url := "ws://192.168.0.159:9092/client/cpu"
	dl := websocket.Dialer{}
	con, _, err := dl.Dial(url, nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 连接服务端websocket失败 ============="), zap.Error(err))
		return
	}
	if con == nil {
		return
	}
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		ubzer.MLog.Error("获取cpu使用率失败", zap.Error(err))
		return
	}
	cpuPercent := strconv.FormatInt(int64(percent[0]), 10)
	type CpuMessage struct {
		Ip      string `json:"ip"`
		Percent string `json:"percent"`
	}
	m := &CpuMessage{
		Ip:      ip,
		Percent: cpuPercent,
	}
	res, _ := json.Marshal(m)
	err = con.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 往服务端推送客户端cpu使用率失败  ============="), zap.Error(err))
		return
	}
}

// 每30秒向服务端推送 客户端 memory使用率
// PushClientMemPercent
func PushClientMemPercent() {
	ip := utils.GetIP()
	//d := "ws://" + config.Config().GoFileServe + "/client/mem"
	url := "ws://192.168.0.159:9092/client/mem"
	dl := websocket.Dialer{}
	con, _, err := dl.Dial(url, nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 连接服务端websocket失败 ============="), zap.Error(err))
		return
	}
	if con == nil {
		return
	}
	mem, err := mem2.VirtualMemory()
	if err != nil {
		ubzer.MLog.Error("获取内存使用率失败", zap.Error(err))
		return
	}
	memPercent := strconv.FormatInt(int64(mem.UsedPercent), 10)
	type MemMessage struct {
		Ip      string `json:"ip"`
		Percent string `json:"percent"`
	}
	m := &MemMessage{
		Ip:      ip,
		Percent: memPercent,
	}
	res, _ := json.Marshal(m)
	err = con.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 往服务端推送客户端内存使用率失败  ============="), zap.Error(err))
		return
	}
}

// 每30秒向服务端推送 客户端 disk使用率
// PushClientMemPercent
func PushClientDiskPercent() {
	ip := utils.GetIP()
	//d := "ws://" + config.Config().GoFileServe + "/client/disk"
	url := "ws://192.168.0.159:9092/client/disk"
	dl := websocket.Dialer{}
	con, _, err := dl.Dial(url, nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 连接服务端websocket失败 ============="), zap.Error(err))
		return
	}
	if con == nil {
		return
	}
	parts, err := disk.Partitions(true)
	if err != nil {
		ubzer.MLog.Error("获取y硬盘使用率失败", zap.Error(err))
		return
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		ubzer.MLog.Error("获取y硬盘使用率失败", zap.Error(err))
		return
	}
	diskPercent := strconv.FormatInt(int64(diskInfo.UsedPercent), 10)
	type DiskMessage struct {
		Ip      string `json:"ip"`
		Percent string `json:"percent"`
	}
	m := &DiskMessage{
		Ip:      ip,
		Percent: diskPercent,
	}
	res, _ := json.Marshal(m)
	err = con.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		return
	}
}
