package router

import (
	"net/http"

	"bz.service.cloud.monitoring/client/config"
	"bz.service.cloud.monitoring/client/machine"

	"github.com/labstack/echo"
)

// RunClientServer
func RunClientServer() {
	e := echo.New()

	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "pong")
	})

	//// 客户端系统信息改为 websocket -- 废弃
	//sys := e.Group("/c/sys")
	//{
	//	sys.GET("/cpu", machine.WsGetCpuSample)
	//	sys.GET("/mem", machine.WsGetMemSample)
	//	sys.GET("/disk", machine.WsGetDiskSample)
	//}

	m := e.Group("/c/machine")
	{
		m.GET("/receive/com", machine.WsReceiveCom)
	}

	e.Logger.Fatal(e.Start(config.Config().ClientHttpBind))
}
