package router

import (
	"time"

	"github.com/c/monitor-system/server/internal/v1/handler"

	"github.com/c/monitor-system/common/middlewares"
	"github.com/c/monitor-system/common/ubzer"
	"github.com/c/monitor-system/server/bootstarp"
	"github.com/c/monitor-system/server/config"
	middle "github.com/c/monitor-system/server/internal/v1/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunServer
func RunServer() {
	e := echo.New()
	bootstarp.InitBoot()
	e.Static("/", "dist")
	e.Static("/monitor/index", "dist")
	e.Static("/monitor/board", "dist")
	e.Static("/monitor/admin/list", "dist")
	e.Static("/monitor/serve/list", "dist")
	e.Static("/monitor/machine/list", "dist")
	e.Static("/monitor/machine/dial", "dist")
	e.Static("/monitor/login", "dist")
	e.Static("/monitor/log/list", "dist")

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
	e.GET("/api/dl", handler.Download)

	// 初始化客户端
	e.GET("/init/client", handler.InitClient)
	// 客户端升级
	e.GET("/client/up", handler.ClientUpgrade)
	// 接收客户端cpu使用率
	e.GET("/client/cpu", handler.ClientCpu)
	// 接收客户端mem使用率
	e.GET("/client/mem", handler.ClientMemory)
	// 接收客户端disk使用率
	e.GET("/client/disk", handler.ClientDisk)

	//管理员增删改查
	admin := e.Group("/api/admin")
	admin.Use(middle.AuthCheck())
	{
		//新管理员创建
		admin.POST("/register", handler.AdminRegister)
		//查看管理员信息列表
		admin.POST("/select", handler.SelectAdmin)
		//根据id变更管理员信息
		admin.POST("/update", handler.UpdateAdminById)
		//根据ID删除管理员信息记录
		admin.POST("/delete", handler.DeleteAdmin)

		admin.POST("/enable/disable", handler.EnableDisableAdmin)

	}
	//查看日志
	l := e.Group("/api/log")
	l.Use(middle.AuthCheck())
	{
		//查看日志信息列表
		l.POST("/query", handler.LogQuery)
	}

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
		serve.POST("/upgrade/record", handler.UpgradeRecord)
	}

	// 机器路由组
	machine := e.Group("/api/machine")
	machine.Use(middle.AuthCheck())
	{
		machine.POST("/list", handler.MachineList)
		machine.GET("/all", handler.AllMachine)
		machine.POST("/send/com", handler.SendCom)
		machine.POST("/update/remark", handler.UpdateMachineRemark)
		machine.POST("/upgrade", handler.UpgradeClientMachine)
		machine.POST("/upgrade/record", handler.UpgradeClientMachineRecord)
		//machine.POST("/delete", handler.DeleteMachine)
	}

	// 报警信息
	wn := e.Group("/api/warning")
	wn.Use(middle.AuthCheck())
	{
		// 检测报警
		wn.GET("/serve/check/list", handler.ServeCheckRecordList)
		wn.POST("/ignore/serve/check/record", handler.IgnoreServeCheckRecord)

		// 风险报警
		wn.GET("/machine/check/list", handler.MachineCheckList)
		wn.POST("/ignore/machine/check/record", handler.IgnoreMachineCheckRecord)
	}

	e.GET("/ssh", handler.ShellWeb)

	e.Logger.Fatal(e.Start(config.Config().HTTPBind))
}
