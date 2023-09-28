package server

import (
	"echo-login-app/api/controller"
	"net/http"
	"os"

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

	// ユーザー
	var uc controller.UserController
	usr := e.Group("/user")
	usr.GET("", uc.GetAll)
	usr.POST("", uc.Create)
	usr.GET("/id/:id", uc.GetByID)
	usr.GET("/username/:username", uc.GetByName)
	usr.PUT("/:id", uc.PutByID)
	usr.DELETE("/:id", uc.Delete)
	usr.POST("/login", uc.Login)

	// メモ
	var mc controller.MemoController
	memo := e.Group("/memo")
	memo.Use(middleware.JWT([]byte(os.Getenv("TOKEN_KEY"))))
	memo.GET("", mc.GetAll)
	memo.POST("", mc.Create)
	memo.GET("/id/:id", mc.GetByID)
	memo.GET("/user_id/:user_id", mc.GetByUserID)
	memo.PUT("/:id", mc.PutByID)
	memo.DELETE("/:id", mc.Delete)

	// コイン
	var cc controller.CoinController
	coin := e.Group("/coin")
	coin.Use(middleware.JWT([]byte(os.Getenv("TOKEN_KEY"))))
	coin.GET("", cc.GetAll)
	coin.POST("", cc.Create)
	coin.GET("/id/:id", cc.GetByID)
	coin.GET("/user_id/:user_id", cc.GetByUserID)
	coin.PUT("/:user_id", cc.PutByUserID)
	coin.DELETE("/:user_id", cc.Delete)

	// アイテム
	var ic controller.ItemController
	item := e.Group("/item")
	item.Use(middleware.JWT([]byte(os.Getenv("TOKEN_KEY"))))
	item.GET("", ic.GetAll)
	item.GET("/:id", ic.GetByID)
	item.DELETE("/:id", ic.Delete)

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
