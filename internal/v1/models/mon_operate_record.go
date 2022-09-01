package models

import (
	"time"
)

type MonOperateRecord struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	AdminId       int64     `json:"admin_id" xorm:"not null comment('管理id') BIGINT"`
	AdminUsername string    `json:"admin_username" xorm:"not null comment('管理员名字') VARCHAR(255)"`
	AdminRealName string    `json:"admin_real_name" xorm:"not null comment('管理员真实名字') VARCHAR(255)"`
	Url           string    `json:"url" xorm:"not null comment('api路径') VARCHAR(255)"`
	Method        string    `json:"method" xorm:"not null comment('方法') VARCHAR(255)"`
	Content       string    `json:"content" xorm:"not null comment('操作内容') TEXT"`
	CreatedAt     time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' comment('操作时间') TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version       int64     `json:"version" xorm:"not null comment('版本号') BIGINT"`
}

func (mor *MonOperateRecord) TableName() string {
	return "mon_operate_record"
}
