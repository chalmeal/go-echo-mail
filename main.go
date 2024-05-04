package main

import (
	"go-echo-mail/api"
	"go-echo-mail/config"
	"log"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port, err := config.Cfg.Section("server").GetKey("SERVER_PORT")
	if err != nil {
		log.Fatal(err)
	}

	e := api.RegisterRouter()
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":" + port.String()))
}
