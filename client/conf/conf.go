package _conf

import (
	"log"

	"github.com/go-ini/ini"
)

var Cfg *ini.File

type MonitorConfig struct {
	RunMode          bool
	ClientHttpBind   string
	ClientLoggerPath string
	GoFileServe      string
}

// LoadMonitorConfig
func LoadMonitorConfig() *MonitorConfig {
	var err error

	Cfg, err = ini.Load("/root/client/conf/conf.ini")
	//Cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse conf", err)
	}

	// client 配置节点读取
	client, err := Cfg.GetSection("client")
	if err != nil {
		log.Fatal(2, "Fail to get section 'client': %v", err)
	}

	Config := &MonitorConfig{
		RunMode:          Cfg.Section("").Key("RUN_MODE").MustBool(),
		ClientHttpBind:   client.Key("ClientHttpBind").MustString(""),
		ClientLoggerPath: client.Key("ClientLoggerPath").MustString(""),
		GoFileServe:      client.Key("GoFileServe").MustString(""),
	}
	return Config
}
