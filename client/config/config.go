package config

import _conf "bz.service.cloud.monitoring/client/conf"

var (
	config *_conf.MonitorConfig
)

func Config() *_conf.MonitorConfig {
	return config
}

// ParseConfig
func ParseConfig() {
	config = _conf.LoadMonitorConfig()
}
