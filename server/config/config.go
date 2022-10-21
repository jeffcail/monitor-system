package config

import _conf "github.com/c/monitor-system/server/conf"

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
