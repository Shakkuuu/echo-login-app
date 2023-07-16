package server

import (
	"echo-login-app/backend/controller"
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
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e.Renderer = renderer

	var uc controller.UserController
	e.GET("/", uc.Index)
	e.GET("/signup", uc.SignupView)
	e.GET("/login", uc.LoginView)
	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)

	app := e.Group("/app")
	var ac controller.AppController
	app.Use(controller.SessionCheck)
	app.GET("", ac.Top)

	e.Logger.Fatal(e.Start(":8082"))
}
