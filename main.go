package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is dongbang")
	})

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.0.0")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
