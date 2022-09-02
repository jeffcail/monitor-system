package main

import (
	"bz.service.cloud.monitoring/client/bootstarp"
	"bz.service.cloud.monitoring/client/router"
)

func main() {
	bootstarp.InitBoot()
	router.RunClientServer()
}
