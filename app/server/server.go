package server

import (
	"echo-login-app/app/controller"
	"io"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init(sk string) {
	e := echo.New()
	e.Use(middleware.Recover())
	// ログの整理
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
	// セッション用ミドルウェア
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sk))))

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e.Renderer = renderer

	var auc controller.AuthController

	// ログイン系
	var uc controller.UserController
	e.GET("/", uc.Index)
	e.GET("/signup", uc.SignupView)
	e.GET("/login", uc.LoginView)
	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)

	// 設定系
	setting := e.Group("/setting")
	setting.Use(auc.SessionCheck)
	setting.GET("/logout", uc.Logout)
	setting.GET("/changename", uc.ChangeNameView)
	setting.POST("/changename", uc.ChangeName)
	setting.GET("/changepassword", uc.ChangePasswordView)
	setting.POST("/changepassword", uc.ChangePassword)
	setting.GET("/delete", uc.Delete)

	// ログイン後のアプリ系
	app := e.Group("/app")
	var apc controller.AppController
	app.Use(auc.SessionCheck)
	app.GET("", apc.Top)
	app.GET("/userpage", uc.UserPage)

	// メモ機能
	memo := app.Group("/memo")
	var mc controller.MemoController
	memo.GET("", mc.Top)
	memo.GET("/create", mc.CreatePage)
	memo.POST("/create", mc.Create)
	memo.GET("/view/:id", mc.ContentView)
	memo.GET("/delete/:id", mc.Delete)
	memo.POST("/change/:id", mc.Change)

	// コイン機能
	coin := app.Group("/coin")
	var cc controller.CoinController
	coin.GET("", cc.Top)
	coin.POST("/add", cc.QtyAdd)
	coin.POST("/sub", cc.QtySub)

	// ガチャ機能
	gacha := app.Group("/gacha")
	var gc controller.GachaController
	gacha.GET("", gc.Top)
	gacha.POST("/draw", gc.Draw)

	// アイテム機能
	item := app.Group("/item")
	var ic controller.ItemController
	item.GET("", ic.Top)

	// ゲーム機能
	game := app.Group("/game")

	// ShotGame
	shot := game.Group("/shot")
	var sc controller.ShotGameController
	shot.GET("", sc.Top)

	e.Logger.Fatal(e.Start(":8082"))
}
