package daos

import (
	"errors"

	"github.com/go-xorm/xorm"

	"github.com/c/server-monitoring/common/db"
	"github.com/c/server-monitoring/server/internal/v1/models"
	"github.com/c/server-monitoring/server/internal/v1/params"
)

// AddUpgradeServeRecord
func AddUpgradeServeRecord(param *params.UpgradeServeParams, packageName string) error {
	m := &models.MonUpgradeServeRecord{
		ServeId:        param.ServeId,
		ServeName:      param.ServeName,
		ServeAddress:   param.ServeAddress,
		ServeState:     param.ServeState,
		PackageName:    packageName,
		UpgradeVersion: param.UpgradeVersion,
	}
	i, err := db.Mysql.Insert(m)
	if err != nil {
		return err
	}
	if i != 1 {
		return errors.New("升级服务失败")
	}
	return nil
}

// CheckUpgradeVersion
func CheckUpgradeVersion(serveAddress string) (*models.MonUpgradeServeRecord, error) {
	m := &models.MonUpgradeServeRecord{}
	_, err := db.Mysql.Where("serve_address = ?", serveAddress).Desc("id").Get(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// UpgradeRecord
func UpgradeRecord(param *params.UpgradeRecord) (int64, []*models.MonUpgradeServeRecord, error) {
	data := make([]*models.MonUpgradeServeRecord, 0)
	var query *xorm.Session
	var query2 *xorm.Session
	query = db.Mysql.Limit(param.PageSize, (param.Page-1)*param.PageSize)
	query2 = db.Mysql.NewSession()

	count, _ := query2.Where("serve_address = ?", param.ServeAddress).Count(&models.MonUpgradeServeRecord{})
	query.Desc("id")
	err := query.Where("serve_address = ?", param.ServeAddress).Find(&data)
	return count, data, err
}
