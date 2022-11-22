package config

import _conf "github.com/c/monitor-system/client/conf"

var (
	config *_conf.MonitorConfig
)

func Config() *_conf.MonitorConfig {
	return config
}

// ParseConfig
func ParseConfig(d string) {
	config = _conf.LoadMonitorConfig(d)
}
