package daos

import (
	"errors"

	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	"github.com/spf13/cast"
)

// CreateMachineCheckRecord
func CreateMachineCheckRecord(ip string, category string, percent int) error {
	m := &models.MonMachineCheckRecord{}
	has, err := db.Mysql.Where("machine_ip = ?", ip).Where("category = ?", category).Get(m)
	if err != nil {
		return err
	}
	if !has {
		m.MachineIp = ip
		m.Category = category
		m.Percent = cast.ToString(percent)

		if percent >= 60 && percent < 70 {
			m.Level = "低危"
		} else if percent >= 70 && percent < 80 {
			m.Level = "中危"
		} else {
			m.Level = "高危"
		}

		insert, err := db.Mysql.Insert(m)
		if err != nil {
			return err
		}
		if insert != 1 {
			return errors.New("写入服务器报警信息失败")
		}
	}
	if percent >= 60 && percent < 70 {
		m.Level = "低危"
	} else if percent >= 70 && percent < 80 {
		m.Level = "中危"
	} else {
		m.Level = "高危"
	}

	update, err := db.Mysql.Where("machine_ip = ?", ip).Where("category = ?", category).Update(m)
	if err != nil {
		return err
	}
	if update != 1 {
		return errors.New("修改服务器报警信息失败")
	}

	return nil
}

// MachineCheckList
func MachineCheckList() (int64, []*models.MonMachineCheckRecord, error) {
	data := make([]*models.MonMachineCheckRecord, 0)
	count, _ := db.Mysql.Count(&models.MonMachineCheckRecord{})
	err := db.Mysql.Find(&data)
	if err != nil {
		return 0, nil, err
	}
	return count, data, nil
}

// IgnoreMachineCheckRecord
func IgnoreMachineCheckRecord(id int64) error {
	s := &models.MonMachineCheckRecord{}
	has, err := db.Mysql.ID(id).Get(s)
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
