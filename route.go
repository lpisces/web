package main

import (
	"github.com/labstack/echo"
	"github.com/lpisces/web/config"
	"github.com/lpisces/web/handler"
)

func route(e *echo.Echo, conf *config.Config) {

	e.GET("/", handler.HandleWelcome())
	e.GET("/config", handler.HandleConfig(conf))

}
