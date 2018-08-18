package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lpisces/web/config"
	"github.com/lpisces/web/model"
	"gopkg.in/urfave/cli.v1"
	"html/template"
	"os"
	"path/filepath"
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
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "debug mode",
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
		&config.DB{},
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

	// template
	templates := template.New("")
	templatePath := "template"

	if _, err := os.Stat(templatePath); err == nil {
		err = filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				_, err := templates.ParseGlob(path)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			e.Logger.Info(err)
			return err
		}
	}

	e.Renderer = &Template{
		templates: templates,
	}

	// public
	e.Static("/public", "public")

	// route
	if err := route(e, Conf); err != nil {
		return err
	}

	// set log level
	if Conf.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	e.Logger.Debugf("http server started on %s:%s, debug: %v", Conf.Srv.Host, Conf.Srv.Port, Conf.Debug)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", Conf.Srv.Host, Conf.Srv.Port)))

	defer model.DB.Close()
	return
}
