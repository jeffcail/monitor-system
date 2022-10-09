package driver

import (
	"bz.service.cloud.monitoring/common/utils"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v5"
)

// CreateRedis
func CreateRedis(addr string, password string, db int) (*redis.Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := rc.Ping().Result()
	utils.CheckErr(err)
	return rc, nil
}
