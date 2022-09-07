package models

import "time"

type MonServeCheckRecord struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	ServeId       int64     `json:"serve_id" xorm:"not null comment('服务id') BIGINT"`
	ServeName     string    `json:"serve_name" xorm:"not null comment('服务名称') VARCHAR(255)"`
	ServeAddress  string    `json:"serve_address" xorm:"comment('服务地址') VARCHAR(255)"`
	LastCheckTime time.Time `json:"last_check_time" xorm:"comment('上次检测时间') TIMESTAMP"`
	CreatedAt     time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version       int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (m *MonServeCheckRecord) TableName() string {
	return "mon_serve_check_record"
}
