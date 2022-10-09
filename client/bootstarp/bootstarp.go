package bootstarp

import (
	"bz.service.cloud.monitoring/common/ubzer"

	"bz.service.cloud.monitoring/client/config"
	"bz.service.cloud.monitoring/client/machine"
)

// InitBoot
func InitBoot() {
	parseRemoteConfig()
	ubzer.InitLogger(config.Config().ClientLoggerPath)
	machine.GenerateUniqueMachineCode()
}

// parseRemoteConfig
func parseRemoteConfig() {
	config.ParseConfig()
}
