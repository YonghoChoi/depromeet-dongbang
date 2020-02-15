package main

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/service"
	"github.com/YonghoChoi/depromeet-dongbang/model/attendance"
	"github.com/YonghoChoi/depromeet-dongbang/model/notice"
	"github.com/YonghoChoi/depromeet-dongbang/model/packet"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/YonghoChoi/depromeet-dongbang/model/vote"
	"github.com/labstack/echo"
	"net/http"
)

func GetUsers(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	users, err := service.GetUsers()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())

	}
	resp.Data = users
	return nil
}

func Join(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var u user.User
	if err := c.Bind(&u); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	u, err := service.Join(u)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = u
	return nil
}

func Login(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var u user.User
	if err := c.Bind(&u); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	u, err := service.Login(u)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = u
	return nil
}

func GetNotices(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	notices, err := service.GetNotices()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())

	}
	resp.Data = notices
	return nil
}

func CreateNotice(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var n notice.Notice
	if err := c.Bind(&n); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	data, err := service.CreateNotice(n)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func EditNotice(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	id := c.QueryParam("id")
	var n notice.Notice
	if err := c.Bind(&n); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	data, err := service.EditNotice(id, n)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func DelNotice(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	id := c.QueryParam("id")
	data, err := service.DelNotice(id)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func GetVotes(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	votes, err := service.GetVotes()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())

	}
	resp.Data = votes
	return nil
}

func CreateVote(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var v vote.VoteRequest
	if err := c.Bind(&v); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	data, err := service.CreateVote(v)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func EditVote(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	id := c.QueryParam("id")
	var v vote.Vote
	if err := c.Bind(&v); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	data, err := service.EditVote(id, v)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func DelVote(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	id := c.QueryParam("id")
	data, err := service.DelVote(id)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func GetAttendances(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	attendances, err := service.GetAttendances()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())

	}
	resp.Data = attendances
	return nil
}

func CreateAttendance(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var o attendance.Attendance
	if err := c.Bind(&o); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	data, err := service.CreateAttendance(o)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func EditAttendance(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	id := c.QueryParam("id")
	var o attendance.Attendance
	if err := c.Bind(&o); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	data, err := service.EditAttendance(id, o)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}

func DelAttendance(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	id := c.QueryParam("id")
	data, err := service.DelAttendance(id)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = data
	return nil
}
