package router

import (
	"bz.service.cloud.monitoring/bootstarp"
	"bz.service.cloud.monitoring/config"
	"github.com/labstack/echo"
)

// RunServer
func RunServer() {
	e := echo.New()
	bootstarp.InitBoot()

	e.Logger.Fatal(e.Start(config.Config().HTTPBind))
}
