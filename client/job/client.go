package job

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/c/server-monitoring/client/config"

	"github.com/shirou/gopsutil/disk"

	mem2 "github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"

	"github.com/c/server-monitoring/common/utils"

	"github.com/c/server-monitoring/common/ubzer"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// PushClientCpuPercent
func PushClientCpuPercent() {
	ip := utils.GetIP()
	url := "ws://" + config.Config().GoFileServe + "/client/cpu"
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

// PushClientMemPercent
func PushClientMemPercent() {
	ip := utils.GetIP()
	url := "ws://" + config.Config().GoFileServe + "/client/mem"
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

// PushClientDiskPercent
func PushClientDiskPercent() {
	ip := utils.GetIP()
	url := "ws://" + config.Config().GoFileServe + "/client/disk"
	//url := "ws://192.168.0.159:9092/client/disk"
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
		fmt.Println(err.Error())
		return
	}
	fmt.Println(parts)

	var usePercent float64
	for _, part := range parts {
		fmt.Println("读取", part)
		if part.Mountpoint == "/" {
			diskinfo, err := disk.Usage(part.Mountpoint)
			if err != nil {
				//fmt.Println("读取硬盘", part.Mountpoint, "错误:", err.Error())
				continue
			}
			usePercent = diskinfo.UsedPercent
		}
	}

	diskPercent := strconv.FormatInt(int64(usePercent), 10)
	fmt.Printf("====== diskPercent: %v\n", diskPercent)
	type DiskMessage struct {
		Ip      string `json:"ip"`
		Percent string `json:"percent"`
	}

	ms := &DiskMessage{
		Ip:      ip,
		Percent: diskPercent,
	}
	res, _ := json.Marshal(ms)
	fmt.Printf("====== disk: %v\n", string(res))
	err = con.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		return
	}
}
