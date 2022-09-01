package daos

import (
	"errors"

	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
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
