package main

import (
	"net/http"

	"github.com/Richard-R/arrangeit/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Theoretical home page")
	})

	e.GET("/apple", func(c echo.Context) error {
		return c.String(http.StatusCreated, "App won't go to different web pages right? Why use this?")
	})

	e.GET("/users/:id", controllers.GetUser)
	e.GET("/users/:id/:name", controllers.TestTwo)

	e.Logger.Fatal(e.Start(":8081"))

}
