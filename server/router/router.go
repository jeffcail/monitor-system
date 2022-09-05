package router

import (
	"time"

	"bz.service.cloud.monitoring/server/internal/v1/handler"

	"bz.service.cloud.monitoring/common/middlewares"
	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/bootstarp"
	"bz.service.cloud.monitoring/server/config"
	middle "bz.service.cloud.monitoring/server/internal/v1/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunServer
func RunServer() {
	e := echo.New()
	bootstarp.InitBoot()

	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: ubzer.EchoLog,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middlewares.RequestLog())
	e.Use(middleware.BodyDumpWithConfig(middlewares.DefaultBodyDumpConfig))

	// 登陆
	e.POST("/api/login", handler.Login)

	// 菜单路由组
	menu := e.Group("/api/menus")
	menu.Use(middle.AuthCheck())
	{
		menu.GET("/list", handler.MenusList)
	}

	// 首页看板
	cs := e.Group("/api/client")
	cs.Use(middle.AuthCheck())
	{
		cs.POST("/sys/cpu", handler.ClientCpuPercent)
		cs.POST("/sys/men", handler.ClientMemPercent)
		cs.POST("/sys/disk", handler.ClientDiskPercent)
	}

	// 服务检测路由租
	serve := e.Group("/api/serve")
	serve.Use(middle.AuthCheck())
	{
		serve.POST("/create", handler.CreateServe)
		serve.POST("/delete", handler.DeleteServe)
		serve.POST("/list", handler.ServeList)
		serve.POST("/upgrade", handler.UpgradeServe)
	}

	// 机器路由组
	machine := e.Group("/api/machine")
	machine.Use(middle.AuthCheck())
	{
		machine.POST("/list", handler.MachineList)
		machine.GET("/all", handler.AllMachine)
		//machine.GET("/ssh", handler.RunWebSSH)
		//machine.POST("/delete", handler.DeleteMachine)
	}
	e.GET("/ssh", handler.RunWebSSH)

	e.Logger.Fatal(e.Start(config.Config().HTTPBind))
}
