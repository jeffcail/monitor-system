package bootstarp

import (
	"flag"

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
	machine.GenerateUniqueMachineCode()
}

// parseRemoteConfig
func parseRemoteConfig(ip string, port int, cfg string) {
	config.ParseConfig(ip, port, cfg)
}
