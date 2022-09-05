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
	query := db.Mysql.NewSession().Limit(params.PageSize, (params.Page-1)*params.PageSize) //分页
	query2 := db.Mysql.NewSession()                                                        //统计，使用newSession时，当key,v为空时，Mysql中where关键字不生效，查询语句可正常运行
	for k, v := range filter {
		query.Where(k+" = ? ", v)
		query2.Where(k+" = ? ", v)
	}
	//if params.Username != "" {//模糊查询写法
	//	query.Where("username like concat('%',concat( ? ,'%'))", params.Username)
	//}
	count, _ := query2.Count()
	query.Desc("id")
	err := query.Find(&list)

	return count, list, err
}
