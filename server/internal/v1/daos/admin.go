package daos

import (
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	"bz.service.cloud.monitoring/server/internal/v1/params"
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
func AddAdmin(params *models.MonAdmin) error {
	count, err := db.Mysql.Insert(params)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("添加不是1条")
	}
	return nil
}

// 根据输入条件查看管理员信息列表
func SelAdmin(params *params.SelAdminParam, filter map[string]interface{}) (int64, []*models.MonAdmin, error) {
	list := make([]*models.MonAdmin, 0)
	query := db.Mysql.NewSession().Limit(params.PageSize, (params.Page-1)*params.PageSize)
	query2 := db.Mysql.NewSession()
	for k, v := range filter {
		query.Where(k+" = ? ", v)
		query2.Where(k+" = ? ", v)
	}

	count, err := query.FindAndCount(&list)
	//count, _ := query2.Count()
	query.Desc("id")
	//err := query.Find(&list)
	return count, list, err
}

// 根据id 变更管理员信息
func UpdAdmin(bean *models.MonAdmin) error {
	count, err := db.Mysql.ID(bean.Id).Update(bean)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("变更不是1条")
	}
	return nil

}

// GetAdminInfoById
func GetAdminInfoById(id int64) (*models.MonAdmin, error) {
	admin := &models.MonAdmin{}
	has, err := db.Mysql.ID(id).Get(admin)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("查无此用户")
	}
	return admin, nil
}

// DeleteAdminById
func DeleteAdminById(id int64) error {

	admin, err := GetAdminInfoById(id)
	if err != nil {
		return err
	}
	count, err := db.Mysql.ID(id).Delete(admin)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("删除失败")
	}
	return nil
}
