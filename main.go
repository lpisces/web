package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lpisces/web/config"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bootstrap"
	app.Usage = "bootstrap for website server development"

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "start web server",
			Action:  serve,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "env, e",
					Usage: "set run env",
					Value: "development",
				},
				cli.StringFlag{
					Name:  "port, p",
					Usage: "listen port",
				},
				cli.StringFlag{
					Name:  "bind, b",
					Usage: "bind host",
				},
				cli.StringFlag{
					Name:  "config, c",
					Usage: "load config file",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func serve(c *cli.Context) (err error) {
	Conf := &config.Config{
		false,
		"config.ini",
		&config.Srv{
			"0.0.0.0",
			"1323",
		},
	}

	if err = Conf.Load(c); err != nil {
		return
	}

	// new echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.HideBanner = true
	route(e, Conf)

	// set log level
	if Conf.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	e.Logger.Infof("http server started on %s:%s, debug: %v", Conf.Srv.Host, Conf.Srv.Port, Conf.Debug)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", Conf.Srv.Host, Conf.Srv.Port)))
	return
}
