package job

import (
	"fmt"
	"net/http"
	"time"

	"github.com/c/server-monitoring/server/internal/v1/models"

	"go.uber.org/zap"

	"github.com/c/server-monitoring/common/ubzer"
	"github.com/c/server-monitoring/server/internal/v1/daos"
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
		res, err := http.Head(v.ServeAddress)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("检测服务: %v 地址: %v ping不通", v.ServeName, v.ServeAddress))
		}
		if res == nil || res.StatusCode != 200 {
			serveState = 2
			scr := &models.MonServeCheckRecord{
				ServeId:       v.Id,
				ServeName:     v.ServeName,
				ServeAddress:  v.ServeAddress,
				LastCheckTime: v.LastCheckTime,
			}
			err := daos.CreateCheckServeRecord(scr)
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("检测服务 %v : 服务地址: %v 异常记录到异常记录表失败", v.ServeName, v.ServeAddress))
			}
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
