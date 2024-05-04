package api

import (
	"go-echo-mail/api/handlers"
	"go-echo-mail/api/validator"
	"go-echo-mail/domain"

	"github.com/labstack/echo/v4"
)

func RegisterRouter() *echo.Echo {
	e := echo.New()

	// setup
	{
		domain.SetUp()
		e.Validator = validator.NewValidator()
	}

	router := e.Group("/api")

	{
		c := router.Group("/mail")
		// register mail
		c.POST("", handlers.RegisterMail)
		// get all history
		c.GET("/history", handlers.GetAllHistorys)
		// notify mail
		c.POST("/notify", handlers.NotifyMail)
	}

	return e

}
