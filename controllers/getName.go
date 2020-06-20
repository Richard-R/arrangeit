package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Controller struct{}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Hello my name is "+id)
}

func TestTwo(c echo.Context) error {
	id := c.Param("id")
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello my id is "+id+" and my name is "+name)
}
