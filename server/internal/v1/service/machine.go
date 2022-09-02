package service

import (
	"errors"
	"fmt"

	_const "bz.service.cloud.monitoring/common/const"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/internal/v1/daos"
)

type MachineListResult struct {
	ID          int64  `json:"id"`
	MachineCode string `json:"machine_code"`
	Ip          string `json:"ip"`
	HostName    string `json:"host_name"`
	Remark      string `json:"remark"`
	CreatedAt   string `json:"created_at"`
}

// MachineList
func MachineList() ([]*MachineListResult, error) {
	list, err := daos.MachineList()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取机器列表失败"), zap.Error(err))
		return nil, errors.New("获取机器列表失败")
	}

	data := make([]*MachineListResult, 0)
	mlr := &MachineListResult{}

	for _, v := range list {
		mlr.ID = v.Id
		mlr.MachineCode = v.MachineCode
		mlr.Ip = v.Ip
		mlr.HostName = v.Hostname
		mlr.Remark = v.Remark
		mlr.CreatedAt = v.CreatedAt.Format(_const.Layout)
		data = append(data, mlr)
	}
	return data, nil
}
