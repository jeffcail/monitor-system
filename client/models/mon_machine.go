package models

import "time"

type MonMachine struct {
	Id          int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	MachineCode string    `json:"machine_code" xorm:"not null comment('机器码') VARCHAR(255)"`
	Ip          string    `json:"ip" xorm:"not null comment('服务器ip') VARCHAR(255)"`
	Hostname    string    `json:"hostname" xorm:"not null comment('系统名') VARCHAR(255)"`
	Remark      string    `json:"remark" xorm:"not null comment('备注') VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version     int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (m *MonMachine) TableName() string {
	return "mon_machine"
}
