package machine

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cast"

	"github.com/shirou/gopsutil/disk"

	mem2 "github.com/shirou/gopsutil/mem"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"

	"github.com/shirou/gopsutil/cpu"

	"github.com/labstack/echo"
)

// GetHostName
// 获取系统主机名
func GetHostName() (string, error) {
	hostname, err := os.Hostname()
	return hostname, err
}

// GetCpuSample
func GetCpuSample(c echo.Context) error {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		ubzer.MLog.Error("获取cpu使用率失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取cpu使用率失败", ""))
	}

	cpuPercent := strconv.FormatInt(int64(percent[0]), 10)

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "获取cpu使用率成功", cast.ToInt(cpuPercent)))
}

// GetMemSample
func GetMemSample(c echo.Context) error {
	mem, err := mem2.VirtualMemory()
	if err != nil {
		ubzer.MLog.Error("获取内存使用率失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取cpu使用率失败", ""))
	}
	memPercent := strconv.FormatInt(int64(mem.UsedPercent), 10)

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "获取内存使用率成功", cast.ToInt(memPercent)))
}

// GetDiskSample
func GetDiskSample(c echo.Context) error {
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

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "获取硬盘使用率成功",
		cast.ToInt(diskPercent)))
}
