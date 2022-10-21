package bootstarp

import (
	"github.com/c/monitor-system/common/ubzer"

	"github.com/c/monitor-system/client/config"
	"github.com/c/monitor-system/client/machine"
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
