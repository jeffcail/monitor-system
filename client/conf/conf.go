package _conf

import (
	"fmt"
	"log"
	"os"

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
func LoadMonitorConfig(d string) *MonitorConfig {
	var err error
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("=========== 获取当前工作目录失败: %v", err)
	}
	log.Printf("========== 当前工作目录: %v", dir)
	log.Printf("========== 目录命令行参数: %v", d)

	wd := fmt.Sprintf("%s%s", d, "/conf.ini")
	log.Printf("========== 配置文件目录: %v", wd)
	//  ========== 配置文件目录: /root/client/conf/conf.ini

	Cfg, err = ini.Load(wd)
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
