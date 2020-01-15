package main

import (
	"example.com/cmd/api/conf"
	"flag"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	configPath := flag.String("config", "./config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is dongbang")
	})

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.0.0")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
