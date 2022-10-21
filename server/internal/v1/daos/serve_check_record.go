package daos

import (
	"errors"
	"time"

	"github.com/c/monitor-system/common/db"
	"github.com/c/monitor-system/server/internal/v1/models"
)

// CreateCheckServeRecord
func CreateCheckServeRecord(scr *models.MonServeCheckRecord) error {
	s := &models.MonServeCheckRecord{}
	has, err := db.Mysql.Where("serve_id = ?", scr.ServeId).Get(s)
	if err != nil {
		return err
	}
	if !has {
		insert, err := db.Mysql.Insert(scr)
		if err != nil {
			return err
		}
		if insert != 1 {
			return errors.New("写入检测异常的服务失败")
		}
		return nil
	}
	s.LastCheckTime = time.Now()
	update, err := db.Mysql.Where("serve_id = ?", scr.ServeId).Update(s)
	if err != nil {
		return err
	}
	if update != 1 {
		return errors.New("更新检测异常的服务失败")
	}
	return nil
}

// ServeCheckRecordList
func ServeCheckRecordList() (int64, []*models.MonServeCheckRecord, error) {
	data := make([]*models.MonServeCheckRecord, 0)
	count, _ := db.Mysql.Count(&models.MonServeCheckRecord{})
	err := db.Mysql.Find(&data)
	if err != nil {
		return 0, nil, err
	}
	return count, data, nil
}

// IgnoreServeCheckRecord
func IgnoreServeCheckRecord(serveId int64) error {
	s := &models.MonServeCheckRecord{}
	has, err := db.Mysql.Where("serve_id = ?", serveId).Get(s)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("记录不存在, 请尝试刷新页面")
	}
	d, err := db.Mysql.ID(s.Id).Delete(s)
	if err != nil {
		return err
	}
	if d != 1 {
		return errors.New("记录忽略失败")
	}
	return nil
}
