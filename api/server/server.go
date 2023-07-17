package server

import (
	"echo-login-app/api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Recover())
	// ログを見やすいように調整
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

	usr := e.Group("/user")
	usr.GET("", uc.GetAll)
	usr.POST("", uc.Create)
	usr.GET("/id/:id", uc.GetByID)
	usr.GET("/username/:username", uc.GetByName)
	usr.PUT("/:id", uc.PutByID)
	usr.DELETE("/:id", uc.Delete)

	e.Logger.Fatal(e.Start(":8081"))
}

// apiの起動確認用(app起動時に使用)
func Pong(c echo.Context) error {
	type PingCheck struct {
		Status  int
		Message string
	}
	p := PingCheck{Status: 200, Message: "Pong"}
	return c.JSON(http.StatusOK, p)
}
