package models

import "time"

type MonMachineCheckRecord struct {
	Id          int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	MachineIp   string    `json:"machine_ip" xorm:"not null comment('机器ip') VARCHAR(255)"`
	MachineName string    `json:"machine_name" xorm:"not null comment('机器名字') VARCHAR(255)"`
	Category    string    `json:"category" xorm:"comment('类型') VARCHAR(255)"`
	Percent     string    `json:"percent" xorm:"comment('百分比') VARCHAR(255)"`
	Level       string    `json:"level" xorm:"comment('等级 低危: 60% 中危: 70% 高危: 80%') VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version     int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (m *MonMachineCheckRecord) TableName() string {
	return "mon_machine_check_record"
}
