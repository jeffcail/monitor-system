package daos

import (
	"errors"

	"bz.service.cloud.monitoring/client/models"
	"bz.service.cloud.monitoring/common/db"
)

// SaveHostName
func SaveHostName(hostname, machineCode, ip string) error {
	machine := &models.MonMachine{
		MachineCode: machineCode,
		Hostname:    hostname,
		Ip:          ip,
	}
	affected, err := db.Mysql.Insert(machine)
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("新机器部署，写入数据失败")
	}
	return nil
}
