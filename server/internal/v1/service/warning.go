package service

import (
	"fmt"
	"time"

	_const "github.com/c/server-monitoring/common/const"

	"github.com/c/server-monitoring/common/ubzer"
	"github.com/c/server-monitoring/server/internal/v1/daos"
	"github.com/c/server-monitoring/server/internal/v1/models"
	"go.uber.org/zap"
)

type ServeCheckRecordListRes struct {
	Id            int64  `json:"id"`
	ServeId       int64  `json:"serve_id"`
	ServeName     string `json:"serve_name"`
	ServeAddress  string `json:"serve_address"`
	LastCheckTime string `json:"last_check_time"`
}

// ServeCheckRecordList
func ServeCheckRecordList() (int64, []*ServeCheckRecordListRes, error) {
	data := make([]*ServeCheckRecordListRes, 0)
	count, list, err := daos.ServeCheckRecordList()
	if err != nil {
		ubzer.MLog.Error("获取服务检测报警信息失败", zap.Error(err))
		return 0, nil, err
	}
	for _, v := range list {
		sclr := &ServeCheckRecordListRes{
			Id:            v.Id,
			ServeId:       v.ServeId,
			ServeName:     v.ServeName,
			ServeAddress:  v.ServeAddress,
			LastCheckTime: v.LastCheckTime.Format(_const.Layout),
		}
		data = append(data, sclr)
	}
	return count, data, nil
}

// IgnoreServeCheckRecord
func IgnoreServeCheckRecord(serveId int64, admin *models.MonAdmin, url string, method string) error {
	err := daos.IgnoreServeCheckRecord(serveId)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("服务忽略失败 服务id: %v", serveId), zap.Error(err))
		return err
	}

	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 忽略了服务id为"+
		": %v 的报警信息", admin.Username, time.Now(), serveId))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录%v 在 %v 忽略服务id: %v 的报警信息失败", admin.Username, time.Now(), serveId))
	}
	return nil
}

type MachineCheckRecordRes struct {
	Id          int64  `json:"id"`
	MachineIp   string `json:"machine_ip"`
	MachineName string `json:"machine_name"`
	Category    string `json:"category"`
	Percent     string `json:"percent"`
	Level       string `json:"level"`
	CreatedAt   string `json:"created_at"`
}

// MachineCheckList
func MachineCheckList() (int64, []*MachineCheckRecordRes, error) {
	data := make([]*MachineCheckRecordRes, 0)
	count, list, err := daos.MachineCheckList()
	if err != nil {
		ubzer.MLog.Error("获取机器检测报警信息失败", zap.Error(err))
		return 0, nil, err
	}
	for _, v := range list {
		mcrr := &MachineCheckRecordRes{
			Id:          v.Id,
			MachineIp:   v.MachineIp,
			MachineName: v.MachineName,
			Category:    v.Category,
			Percent:     v.Percent,
			Level:       v.Level,
			CreatedAt:   v.CreatedAt.Format(_const.Layout),
		}
		data = append(data, mcrr)
	}
	return count, data, err
}

// IgnoreMachineCheckRecord
func IgnoreMachineCheckRecord(id int64, admin *models.MonAdmin, url, method string) error {
	err := daos.IgnoreMachineCheckRecord(id)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("机器警报忽略失败 id: %v", id), zap.Error(err))
		return err
	}

	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 忽略了机器"+
		"的报警信息", admin.Username, time.Now()))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录%v 在 %v 忽略机器的报警信息失败", admin.Username, time.Now()))
	}
	return nil
}
