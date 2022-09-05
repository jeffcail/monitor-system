package daos

import (
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	params2 "bz.service.cloud.monitoring/server/internal/v1/params"
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

//// DeleteMachine
//func DeleteMachine() {
//
//}
