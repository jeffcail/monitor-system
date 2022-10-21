package bootstarp

import (
	"flag"
	"runtime"

	"github.com/c/monitor-system/common/db"

	"github.com/c/monitor-system/common/utils"
	"github.com/c/monitor-system/server/config"
	"github.com/c/monitor-system/server/driver"
	"github.com/go-xorm/xorm"

	"github.com/c/monitor-system/common/ubzer"
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
