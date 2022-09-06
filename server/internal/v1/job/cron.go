package job

import (
	"fmt"
	"time"

	"bz.service.cloud.monitoring/common/request"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/internal/v1/daos"
)

var serveState int

// checkServeList
func checkServeList() {
	allServe, err := daos.GetAllServe()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("定时检测服务获取服务数据失败"), zap.Error(err))
		return
	}
	for _, v := range allServe {
		checkTime := time.Now()
		res, err := request.Get(v.ServeAddress)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("检测服务: %v 地址: %v ping不通", v.ServeName, v.ServeAddress))
		}
		if string(res) != "pong" {
			serveState = 2
		} else {
			serveState = 1
		}
		v.ServeState = serveState
		v.LastCheckTime = checkTime
		err = daos.UpdateServeById(v)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("检测完服务: %v 更新数据失败", v.ServeName), zap.Error(err))
		}
	}
}

// deleteOperateRecord
func deleteOperateRecord() {
	ts := time.Now().AddDate(0, 0, -30)
	daos.DeleteHistoryRecordLogs(ts)
}
