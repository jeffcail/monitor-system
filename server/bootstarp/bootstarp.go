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

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
}

func InitBoot() {
	parseRemoteConfig()
	initLogger()
	initMysql()
	initRedis()
}

// parseRemoteConfig
func parseRemoteConfig() {
	config.ParseConfig()
}

// initLogger
func initLogger() {
	ubzer.InitLogger(config.Config().LoggerPath)
}

// initMysql
func initMysql() {
	engine, err := driver.CreateMysql(config.Config().DbDsn, config.Config().ShowSql)
	utils.CheckErr(err)
	engine.SetLogger(xorm.NewSimpleLogger(ubzer.XLogger))
	db.DB.SetMysql(engine)
}

// initRedis
func initRedis() {
	rc, err := driver.CreateRedis(config.Config().RedisAddr, config.Config().Password, config.Config().RedisDb)
	utils.CheckErr(err)
	db.DB.SetRedis(rc)
}
