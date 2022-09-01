package daos

import (
	"errors"
	"time"

	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
)

// RecordOperateLog
func RecordOperateLog(adminId int64, adminUserName, realName, url, method, content string) error {
	recordLog := &models.MonOperateRecord{
		AdminId:       adminId,
		AdminUsername: adminUserName,
		AdminRealName: realName,
		Url:           url,
		Method:        method,
		Content:       content,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	i, err := db.Mysql.Insert(recordLog)
	if err != nil {
		return err
	}
	if i != 1 {
		return errors.New("记录操作日志失败")
	}
	return nil
}
