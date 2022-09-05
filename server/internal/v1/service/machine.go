package service

import (
	"errors"
	"fmt"

	"bz.service.cloud.monitoring/server/internal/v1/models"

	params2 "bz.service.cloud.monitoring/server/internal/v1/params"

	_const "bz.service.cloud.monitoring/common/const"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/internal/v1/daos"
)

type MachineListResult struct {
	ID          int64  `json:"id"`
	MachineCode string `json:"machine_code"`
	Ip          string `json:"ip"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	HostName    string `json:"host_name"`
	Remark      string `json:"remark"`
	CreatedAt   string `json:"created_at"`
}

// MachineList
func MachineList(params *params2.MachineListParams) (int64, []*MachineListResult, error) {
	count, list, err := daos.MachineList(params)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取机器列表失败"), zap.Error(err))
		return 0, nil, errors.New("获取机器列表失败")
	}

	data := make([]*MachineListResult, 0)
	mlr := &MachineListResult{}

	for _, v := range list {
		mlr.ID = v.Id
		mlr.MachineCode = v.MachineCode
		mlr.Ip = v.Ip
		mlr.Name = v.Name
		mlr.Password = v.Password
		mlr.HostName = v.Hostname
		mlr.Remark = v.Remark
		mlr.CreatedAt = v.CreatedAt.Format(_const.Layout)
		data = append(data, mlr)
	}
	return count, data, nil
}

// AllMachine
func AllMachine() ([]*models.MonMachine, error) {
	machine, err := daos.AllMachine()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取所有服务器失败"), zap.Error(err))
		return nil, err
	}
	return machine, nil
}

//// DeleteMachine
//func DeleteMachine() {
//
//}
