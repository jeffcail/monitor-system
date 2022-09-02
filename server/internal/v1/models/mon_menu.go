package models

import "time"

type MonMenus struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	ParentId  string    `json:"parent_id" xorm:"not null comment('父级菜单') VARCHAR(10)"`
	MenuName  string    `json:"menu_name" xorm:"not null comment('权限名字') VARCHAR(255)"`
	Icons     string    `json:"icons" xorm:"comment('icons图标') VARCHAR(255)"`
	Url       string    `json:"url" xorm:"comment('路径') VARCHAR(255)"`
	FrontUrl  string    `json:"front_url" xorm:"comment('前端地址url') VARCHAR(255)"`
	Method    string    `json:"method" xorm:"default '' comment('方法') VARCHAR(20)"`
	CreatedAt time.Time `json:"created_at" xorm:"created not null comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated not null comment('更新时间') TIMESTAMP"`
	Version   int64     `json:"version" xorm:"not null comment('版本号') BIGINT version"`
}

func (mm *MonMenus) TableName() string {
	return "mon_menus"
}
