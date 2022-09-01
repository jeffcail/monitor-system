package driver

import (
	"bz.service.cloud.monitoring/utils"
	"gopkg.in/redis.v5"
)

var Rc *RedisConfig

type RedisConfig struct {
	RedisAddr string
	Password  string
	RedisDb   int
}

// CreateRedis
func (r *RedisConfig) CreateRedis(config RedisConfig) (*redis.Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.Password,
		DB:       config.RedisDb,
	})
	_, err := rc.Ping().Result()
	utils.CheckErr(err)
	return rc, nil
}
