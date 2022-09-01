package models

import (
	"time"
)

type MonAdmin struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT"`
	Username      string    `json:"username" xorm:"not null comment('用户昵称') VARCHAR(255)"`
	RealName      string    `json:"real_name" xorm:"not null comment('真实姓名') VARCHAR(255)"`
	Password      string    `json:"password" xorm:"not null comment('密码') VARCHAR(32)"`
	Phone         string    `json:"phone" xorm:"comment('手机号') VARCHAR(11)"`
	RoleId        string    `json:"role_id" xorm:"comment('角色id') VARCHAR(255)"`
	Department    string    `json:"department" xorm:"comment('部门') VARCHAR(255)"`
	OfficePost    string    `json:"office_post" xorm:"comment('职位') VARCHAR(255)"`
	State         int       `json:"state" xorm:"not null default 2 comment('状态 1: 正常 2: 禁用') INT"`
	LastLoginTime time.Time `json:"last_login_time" xorm:"comment('上次登陆时间') TIMESTAMP"`
	CreatedAt     time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version       int64     `json:"version" xorm:"not null comment('版本号') BIGINT"`
}

func (m *MonAdmin) TableName() string {
	return "mon_admin"
}
