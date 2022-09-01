package config

import (
	"bz.service.cloud.monitoring/driver"
	"bz.service.cloud.monitoring/pkg/nacos"
)

type GlobalConfig struct {
	Debug      bool
	HTTPBind   string
	Mysql      driver.MysqlConfig
	Redis      driver.RedisConfig
	LoggerPath string
	Slat       string
}

var config GlobalConfig

func Config() GlobalConfig {
	return config
}

// ParseConfig
func ParseConfig(ip string, port int, cfg string) {
	nacos.Nac.ParseConfig(ip, port, cfg, &config)
}
