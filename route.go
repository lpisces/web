package main

import (
	"github.com/labstack/echo"
	"github.com/lpisces/web/config"
	"github.com/lpisces/web/handler"
	"github.com/lpisces/web/model"
)

func route(e *echo.Echo, conf *config.Config) (err error) {

	if _, err = model.InitDB(conf.DB); err != nil {
		return
	}

	e.GET("/", handler.HandleWelcome())
	e.GET("/config", handler.HandleConfig(conf))

	return
}
