package models

import "time"

type MonUpgradeMachineRecord struct {
	Id              int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	MachineCode     string    `json:"machine_code" xorm:"not null comment('机器码') BIGINT"`
	MachineIp       string    `json:"machine_ip" xorm:"not null comment('机器ip') VARCHAR(255)"`
	MachineHostname string    `json:"machine_hostname" xorm:"comment('主机名称') VARCHAR(255)"`
	MachineRemark   string    `json:"machine_remark" xorm:"comment('机器备注信息') TIMESTAMP"`
	PackageName     string    `json:"package_name" xorm:"comment('客户端包名') VARCHAR(255)"`
	UpgradeVersion  string    `json:"upgrade_version" xorm:"comment('版本号') VARCHAR(255)"`
	Md5Sum          string    `json:"md5_sum" xorm:"comment('文件md5值') VARCHAR(255)"`
	CreatedAt       time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version         int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (this *MonUpgradeMachineRecord) TableName() string {
	return "mon_upgrade_machine_record"
}
