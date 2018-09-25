package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main()  {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo!")
}
