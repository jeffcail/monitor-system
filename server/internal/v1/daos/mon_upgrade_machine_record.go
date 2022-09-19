package daos

import (
	"errors"

	"github.com/go-xorm/xorm"

	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	"bz.service.cloud.monitoring/server/internal/v1/params"
)

// AddUpgradeClientServeRecord
func AddUpgradeClientServeRecord(param *params.UpgradeClientMachineParams, packageName string, md5_sum string) error {
	m := &models.MonUpgradeMachineRecord{
		MachineCode:     param.MachineCode,
		MachineIp:       param.MachineIp,
		MachineHostname: param.MachineHostname,
		MachineRemark:   param.MachineRemark,
		PackageName:     packageName,
		UpgradeVersion:  param.UpgradeVersion,
		Md5Sum:          md5_sum,
	}

	i, err := db.Mysql.Insert(m)
	if err != nil {
		return err
	}
	if i != 1 {
		return errors.New("升级客户端服务失败")
	}
	return nil
}

// CheckClientUpgradeVersion
func CheckClientUpgradeVersion(machineIp string) (*models.MonUpgradeMachineRecord, error) {
	m := &models.MonUpgradeMachineRecord{}
	_, err := db.Mysql.Where("machine_ip = ?", machineIp).Desc("id").Get(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// UpgradeClientMachineRecord
func UpgradeClientMachineRecord(param *params.UpgradeClientMachineRecord) (int64, []*models.MonUpgradeMachineRecord, error) {
	data := make([]*models.MonUpgradeMachineRecord, 0)
	var query *xorm.Session
	var query2 *xorm.Session
	query = db.Mysql.Limit(param.PageSize, (param.Page-1)*param.PageSize)
	query2 = db.Mysql.NewSession()

	count, _ := query2.Where("machine_ip = ?", param.MachineIp).Count(&models.MonUpgradeMachineRecord{})
	query.Desc("id")
	err := query.Where("machine_ip = ?", param.MachineIp).Find(&data)
	return count, data, err
}
