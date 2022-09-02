package router

import (
	"net/http"

	"bz.service.cloud.monitoring/client/config"
	"github.com/labstack/echo"
)

// RunClientServer
func RunClientServer() {
	e := echo.New()

	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "pong")
	})

	e.Logger.Fatal(e.Start(config.Config().ClientHttpBind))
}
