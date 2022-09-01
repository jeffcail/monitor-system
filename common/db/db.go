package db

import (
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"gopkg.in/redis.v5"
)

var DB *DBR

type DBR struct{}

var (
	Mysql *xorm.Engine
	Rc    *redis.Client
)

// SetMysql
func (D *DBR) SetMysql(e *xorm.Engine) {
	Mysql = e
	go func() {
		for {
			Mysql.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()
}

// SetRedis
func (D *DBR) SetRedis(_rc *redis.Client) {
	Rc = _rc

}

// Transaction
func (D *DBR) Transaction(fs ...func(s *xorm.Session) error) error {
	session := Mysql.NewSession()
	session.Begin()
	for _, f := range fs {
		err := f(session)
		if err != nil {
			log.Println(err)
			session.Rollback()
			session.Clone()
			return err
		}
	}
	session.Commit()
	session.Clone()
	return nil
}

func (D *DBR) CloseMysql() {
	Mysql.Close()
}
