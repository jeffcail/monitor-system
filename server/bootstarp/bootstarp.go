package bootstarp

import (
	"flag"
	"runtime"

	"bz.service.cloud.monitoring/common/db"

	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/config"
	"bz.service.cloud.monitoring/server/driver"
	"github.com/go-xorm/xorm"

	"bz.service.cloud.monitoring/common/ubzer"
)

var (
	ip   = flag.String("ip", "192.168.0.125", "The config of ip address")
	port = flag.Int("p", 7848, "The config of port")
	cfg  = flag.String("c", "service-cloud-monitor.yml", "The path of configuration file")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
}

func InitBoot() {
	parseRemoteConfig(*ip, *port, *cfg)
	initLogger()
	initMysql()
	initRedis()
}

// parseRemoteConfig
func parseRemoteConfig(ip string, port int, cfg string) {
	config.ParseConfig(ip, port, cfg)
}

// initLogger
func initLogger() {
	ubzer.InitLogger(config.Config().LoggerPath)
}

// initMysql
func initMysql() {
	engine, err := driver.My.CreateMysql(config.Config().Mysql)
	utils.CheckErr(err)
	engine.SetLogger(xorm.NewSimpleLogger(ubzer.XLogger))
	engine.ShowSQL(config.Config().Mysql.ShowSQL)
	db.DB.SetMysql(engine)
}

// initRedis
func initRedis() {
	rc, err := driver.Rc.CreateRedis(config.Config().Redis)
	utils.CheckErr(err)
	db.DB.SetRedis(rc)
}
