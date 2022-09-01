package driver

import (
	"bz.service.cloud.monitoring/common/utils"
	"github.com/go-xorm/xorm"
)

var My *MysqlConfig

type MysqlConfig struct {
	DbDsn   string
	ShowSQL bool
}

// CreateMysql
func (m *MysqlConfig) CreateMysql(config MysqlConfig) (*xorm.Engine, error) {
	mysql, err := xorm.NewEngine("mysql", config.DbDsn)
	if err != nil {
		return nil, err
	}
	mysql.ShowSQL(config.ShowSQL)
	err = mysql.Ping()
	utils.CheckErr(err)
	return mysql, nil
}
