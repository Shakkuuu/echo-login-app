package server

import (
	"echo-login-app/api/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\n" + `time: ${time_rfc3339_nano}` + "\n" +
			`method: ${method}` + "\n" +
			`remote_ip: ${remote_ip}` + "\n" +
			`host: ${host}` + "\n" +
			`uri: ${uri}` + "\n" +
			`status: ${status}` + "\n" +
			`error: ${error}` + "\n" +
			`latency: ${latency}(${latency_human})` + "\n",
	}))

	var uc controller.UserController
	e.GET("/", Pong)
	e.GET("/user", uc.GetAll)
	e.POST("/user", uc.Create)
	e.GET("/user/id/:id", uc.GetByID)
	e.GET("/user/username/:username", uc.GetByName)
	e.DELETE("/user/:id", uc.Delete)

	e.Logger.Fatal(e.Start(":8081"))
}

func Pong(c echo.Context) error {
	type PingCheck struct {
		Status  int
		Message string
	}
	p := PingCheck{Status: 200, Message: "Pong"}
	return c.JSON(http.StatusOK, p)
}
