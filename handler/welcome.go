package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/lpisces/web/config"
	"net/http"
)

func HandleWelcome() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome")
	}
}

func HandleHost(conf *config.Config) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%v", conf.Srv.Host))
	}
}
