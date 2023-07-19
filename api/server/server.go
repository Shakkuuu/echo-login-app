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

	e.GET("/", Pong)

	var uc controller.UserController
	usr := e.Group("/user")
	usr.GET("", uc.GetAll)
	usr.POST("", uc.Create)
	usr.GET("/id/:id", uc.GetByID)
	usr.GET("/username/:username", uc.GetByName)
	usr.PUT("/:id", uc.PutByID)
	usr.DELETE("/:id", uc.Delete)

	var mc controller.MemoController
	memo := e.Group("/memo")
	memo.GET("", mc.GetAll)
	memo.POST("", mc.Create)
	memo.GET("/id/:id", mc.GetByID)
	memo.GET("/user_id/:user_id", mc.GetByUserID)
	memo.PUT("/:id", mc.PutByID)
	memo.DELETE("/:id", mc.Delete)

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
