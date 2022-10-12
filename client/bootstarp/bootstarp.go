package bootstarp

import (
	"github.com/c/server-monitoring/common/ubzer"

	"github.com/c/server-monitoring/client/config"
	"github.com/c/server-monitoring/client/machine"
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
