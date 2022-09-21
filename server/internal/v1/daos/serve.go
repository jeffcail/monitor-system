package daos

import (
	"errors"
	"time"

	"github.com/go-xorm/xorm"

	"bz.service.cloud.monitoring/common/db"

	"bz.service.cloud.monitoring/server/internal/v1/models"

	params2 "bz.service.cloud.monitoring/server/internal/v1/params"
)

// ServeList
func ServeList(params *params2.ServeListParams) (int64, []*models.MonServe, error) {
	data := make([]*models.MonServe, 0)
	var query *xorm.Session
	var query2 *xorm.Session
	query = db.Mysql.Limit(params.PageSize, (params.Page-1)*params.PageSize)
	query2 = db.Mysql.NewSession()

	count, _ := query2.Count(&models.MonServe{})
	query.Desc("id")
	err := query.Find(&data)
	return count, data, err
}

// CreateServe
func CreateServe(params *params2.CreateServeParams, serveState int, checkTime time.Time) error {
	serve := &models.MonServe{
		ServeName:     params.ServeName,
		ServeAddress:  params.ServeAddress,
		ServeState:    serveState,
		LastCheckTime: checkTime,
	}
	affected, err := db.Mysql.Insert(serve)
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("创建服务检测失败")
	}
	return nil
}

// DeleteServe
func DeleteServe(id int64) (string, error) {
	var name string
	serve := &models.MonServe{}
	has, err := db.Mysql.ID(id).Get(serve)
	if err != nil {
		return "", err
	}
	if !has {
		return "", errors.New("服务不存在")
	}
	name = serve.ServeName
	i, err := db.Mysql.ID(id).Delete(serve)
	if err != nil {
		return "", err
	}
	if i != 1 {
		return "", errors.New("删除服务失败")
	}
	return name, nil
}

// GetAllServe
func GetAllServe() ([]*models.MonServe, error) {
	data := make([]*models.MonServe, 0)
	err := db.Mysql.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateServeById
func UpdateServeById(serve *models.MonServe) error {
	_, err := db.Mysql.ID(serve.Id).Update(serve)
	if err != nil {
		return err
	}
	//if affected != 1 {
	//	return errors.New("修改serve 失败")
	//}
	return nil
}
