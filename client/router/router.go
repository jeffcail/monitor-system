package router

import (
	"net/http"

	"bz.service.cloud.monitoring/client/machine"

	"bz.service.cloud.monitoring/client/config"
	"github.com/labstack/echo"
)

// RunClientServer
func RunClientServer() {
	e := echo.New()

	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "pong")
	})

	// 客户端系统信息
	sys := e.Group("/c/sys")
	{
		sys.GET("/cpu", machine.GetCpuSample)
		sys.GET("/mem", machine.GetMemSample)
		sys.GET("/disk", machine.GetDiskSample)
	}

	e.Logger.Fatal(e.Start(config.Config().ClientHttpBind))
}
