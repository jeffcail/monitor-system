package daos

import (
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	"errors"
)

// GetAdminInfoByUsername
func GetAdminInfoByUsername(username string) (*models.MonAdmin, error) {
	admin := &models.MonAdmin{}
	has, err := db.Mysql.Where("username = ?", username).Get(admin)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("查无此用户")
	}
	return admin, nil
}

// UpdateAdminById
func UpdateAdminById(admin *models.MonAdmin) error {
	affected, err := db.Mysql.Id(admin.Id).Update(admin)
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("登陆出现异常")
	}
	return nil
}

// 管理员插入
func AddAdmin(admin *models.MonAdmin) error {
	count, err := db.Mysql.Insert(admin)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("添加不是1条")
	}
	return nil
}

// 查询
//func SelAdmin(admin *models.MonAdmin,params *params2.ServeListParams) (int64,[]*models.MonAdmin,error) {
//	list:=make([]*models.MonAdmin, 0)
//	var query *xorm.Session
//	query = db.Mysql.Limit(params.PageSize, (params.Page-1)*params.PageSize)
//	count,_ :=query.Count(&admin)
//	query.Desc("id")
//
//
//	return
//}
