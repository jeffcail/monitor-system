package _conf

import (
	"log"

	"github.com/go-ini/ini"
)

var Cfg *ini.File

type MonitorConfig struct {
	RunMode         bool
	HTTPBind        string
	UpPkgPath       string
	UpClientPkgPath string
	LoggerPath      string
	Slat            string
	ClientHttpBind  string
	DbDsn           string
	ShowSql         bool
	RedisAddr       string
	Password        string
	RedisDb         int
}

// LoadMonitorConfig
func LoadMonitorConfig() *MonitorConfig {
	var err error

	Cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse conf", err)
	}

	// server 配置节点读取
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}

	// client 配置节点读取
	client, err := Cfg.GetSection("client")
	if err != nil {
		log.Fatal(2, "Fail to get section 'client': %v", err)
	}

	// database 配置节点读取
	database, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//redis 配置节点读取
	redis, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}

	Config := &MonitorConfig{
		RunMode:         Cfg.Section("").Key("RUN_MODE").MustBool(),
		HTTPBind:        server.Key("HTTPBind").MustString(""),
		UpPkgPath:       server.Key("UpPkgPath").MustString(""),
		UpClientPkgPath: server.Key("UpClientPkgPath").MustString(""),
		LoggerPath:      server.Key("LoggerPath").MustString("LoggerPath"),
		Slat:            server.Key("Slat").MustString(""),
		ClientHttpBind:  client.Key("ClientHttpBind").MustString(""),
		DbDsn:           database.Key("DbDsn").MustString(""),
		ShowSql:         database.Key("ShowSql").MustBool(),
		RedisAddr:       redis.Key("RedisAddr").MustString(""),
		Password:        redis.Key("Password").MustString(""),
		RedisDb:         redis.Key("RedisDb").MustInt(),
	}
	return Config
}
