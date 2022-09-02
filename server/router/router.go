package router

import (
	"time"

	"bz.service.cloud.monitoring/server/internal/v1/handler"

	"bz.service.cloud.monitoring/common/middlewares"
	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/bootstarp"
	"bz.service.cloud.monitoring/server/config"
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

	admin := e.Group("/api/admin")
	{
		admin.POST("/register", handler.AdminRegister) //管理员注册
	}

	e.Logger.Fatal(e.Start(config.Config().HTTPBind))
}
