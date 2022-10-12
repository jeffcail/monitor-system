package daos

import (
	"errors"

	"github.com/c/server-monitoring/common/db"
	"github.com/c/server-monitoring/server/internal/v1/models"
	params2 "github.com/c/server-monitoring/server/internal/v1/params"
	"github.com/go-xorm/xorm"
)

// MachineList
func MachineList(params *params2.MachineListParams) (int64, []*models.MonMachine, error) {

	data := make([]*models.MonMachine, 0)
	var query *xorm.Session
	var query2 *xorm.Session
	query = db.Mysql.Limit(params.PageSize, (params.Page-1)*params.PageSize)
	query2 = db.Mysql.NewSession()

	count, _ := query2.Count(&models.MonMachine{})
	query.Desc("id")
	err := query.Find(&data)
	return count, data, err
}

// AllMachine
func AllMachine() ([]*models.MonMachine, error) {
	data := make([]*models.MonMachine, 0)
	err := db.Mysql.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateMachineRemark
func UpdateMachineRemark(param *params2.UpdateMachineRemarkParams) error {
	m := &models.MonMachine{}
	has, err := db.Mysql.Where("machine_code = ?", param.MachineCode).Where("ip = ?", param.Ip).Get(m)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("机器不存在")
	}
	m.Remark = param.Remark
	affected, err := db.Mysql.ID(m.Id).Update(m)
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("修改备注失败")
	}
	return nil
}

// GetMachineInfoFromDbByIp
func GetMachineInfoFromDbByIp(ip string) (*models.MonMachine, error) {
	m := &models.MonMachine{}
	_, err := db.Mysql.Where("ip = ?", ip).Get(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

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

//// DeleteMachine
//func DeleteMachine() {
//
//}
