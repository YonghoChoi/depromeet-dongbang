package main

import (
	"flag"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/conf"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	configPath := flag.String("config", "./conf/config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "이 앱은 디프만 파이널 프로젝트 동방입니다.")
	})

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.0.0")
	})

	e.POST("/join", Join)
	e.POST("/login", Login)
	e.GET("/api/users", GetUsers)
	e.GET("/api/notices", GetNotices)
	e.POST("/api/notices", CreateNotice)
	e.PUT("/api/notices", EditNotice)
	e.DELETE("/api/notices", DelNotice)
	e.GET("/api/votes", GetVotes)
	e.POST("/api/votes", CreateVote)
	e.PUT("/api/votes", EditVote)
	e.DELETE("/api/votes", DelVote)
	e.GET("/api/attendances", GetAttendances)
	e.POST("/api/attendances", CreateAttendance)
	e.PUT("/api/attendances", EditAttendance)
	e.DELETE("/api/attendances", DelAttendance)

	e.Logger.Fatal(e.Start(":8000"))
}
