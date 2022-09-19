package models

import "time"

type MonUpgradeServeRecord struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	ServeId        int64     `json:"serve_id" xorm:"not null comment('服务id') BIGINT"`
	ServeName      string    `json:"serve_name" xorm:"not null comment('服务名称') VARCHAR(255)"`
	ServeAddress   string    `json:"serve_address" xorm:"comment('服务地址') VARCHAR(255)"`
	ServeCreatedAt time.Time `json:"last_check_time" xorm:"comment('服务创建时间') TIMESTAMP"`
	ServeState     int       `json:"serve_state" xorm:"comment('服务状态') INT"`
	PackageName    string    `json:"package_name" xorm:"comment('包名') VARCHAR(255)"`
	UpgradeVersion string    `json:"upgrade_version" xorm:"comment('版本号') VARCHAR(255)"`
	CreatedAt      time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version        int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (this *MonUpgradeServeRecord) TableName() string {
	return "mon_upgrade_serve_record"
}
