package bootstarp

import (
	"flag"

	"github.com/go-xorm/xorm"

	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/common/driver"
	"bz.service.cloud.monitoring/common/utils"

	"bz.service.cloud.monitoring/common/ubzer"

	"bz.service.cloud.monitoring/client/config"
	"bz.service.cloud.monitoring/client/machine"
)

var (
	ip   = flag.String("ip", "192.168.0.125", "The config of ip address")
	port = flag.Int("p", 7848, "The config of port")
	cfg  = flag.String("c", "service-cloud-monitor.yml", "The path of configuration file")
)

func init() {
	flag.Parse()
}

// InitBoot
func InitBoot() {
	parseRemoteConfig(*ip, *port, *cfg)
	ubzer.InitLogger(config.Config().ClientLoggerPath)
	initRedis()
	initMysql()
	machine.GenerateUniqueMachineCode()
}

// parseRemoteConfig
func parseRemoteConfig(ip string, port int, cfg string) {
	config.ParseConfig(ip, port, cfg)
}

// initRedis
func initRedis() {
	rc, err := driver.Rc.CreateRedis(config.Config().Redis)
	utils.CheckErr(err)
	db.DB.SetRedis(rc)
}

// initMysql
func initMysql() {
	engine, err := driver.My.CreateMysql(config.Config().Mysql)
	utils.CheckErr(err)
	engine.SetLogger(xorm.NewSimpleLogger(ubzer.XLogger))
	engine.ShowSQL(config.Config().Mysql.ShowSQL)
	db.DB.SetMysql(engine)
}
