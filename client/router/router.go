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

	// 客户端系统信息改为 websocket
	sys := e.Group("/c/sys")
	{
		sys.GET("/cpu", machine.WsGetCpuSample)
		sys.GET("/mem", machine.WsGetMemSample)
		sys.GET("/disk", machine.WsGetDiskSample)
	}

	m := e.Group("/c/machine")
	{
		m.GET("/receive/com", machine.WsReceiveCom)
		//m.GET("/upgrade/client", machine.WsUpgradeClient)
	}

	e.Logger.Fatal(e.Start(config.Config().ClientHttpBind))

	// 客户端系统信息
	//sys := e.Group("/c/sys")
	//{
	//	sys.GET("/cpu", machine.GetCpuSample)
	//	sys.GET("/mem", machine.GetMemSample)
	//	sys.GET("/disk", machine.GetDiskSample)
	//}

	// 服务升级指令
	//serve := e.Group("/c/serve")
	//{
	//	serve.GET("/upgrade", machine.ServeUpgrade)
	//}

	//m := e.Group("/c/machine")
	//{
	//	m.GET("/receive/com", machine.ReceiveCom)
	//}
}
