package daos

import (
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
)

// MenuList
func MenuList() ([]*models.MonMenus, []*models.MonMenus) {
	data := make([]*models.MonMenus, 0)
	children := make([]*models.MonMenus, 0)

	err := db.Mysql.Asc("id").Where("parent_id = ?", "0").Find(&data)
	if err != nil {
		return nil, nil
	}
	for _, v := range data {
		err = db.Mysql.Asc("id").Where("parent_id = ?", v.Id).Find(&children)
		if err != nil {
			return data, nil
		}
	}

	return data, children
}
