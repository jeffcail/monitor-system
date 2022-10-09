package driver

import (
	"bz.service.cloud.monitoring/common/utils"
	"github.com/go-xorm/xorm"
)

// CreateMysql
func CreateMysql(dbDsn string, showSql bool) (*xorm.Engine, error) {
	mysql, err := xorm.NewEngine("mysql", dbDsn)
	if err != nil {
		return nil, err
	}
	mysql.ShowSQL(showSql)
	err = mysql.Ping()
	utils.CheckErr(err)
	return mysql, nil
}
