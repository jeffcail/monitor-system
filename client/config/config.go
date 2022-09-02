package config

import (
	"bz.service.cloud.monitoring/common/driver"
	"bz.service.cloud.monitoring/common/nacos"
)

type ClientGlobalConfig struct {
	Debug            bool
	ClientHttpBind   string
	Redis            driver.RedisConfig
	Mysql            driver.MysqlConfig
	ClientLoggerPath string
}

var config ClientGlobalConfig

func Config() ClientGlobalConfig {
	return config
}

// ParseConfig
func ParseConfig(ip string, port int, cfg string) {
	nacos.Nac.ParseConfig(ip, port, cfg, &config)
}
