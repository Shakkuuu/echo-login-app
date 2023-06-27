package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", Hello)
	e.GET("/map", Mappp)
	e.GET("struct", Structtt)
	e.Logger.Fatal(e.Start(":8082"))
}

func Mappp(c echo.Context) error {
	m := map[string]interface{}{
		"name": "Yamada",
		"age":  "31",
	}
	return c.Render(http.StatusOK, "yamada", m)
}

func Structtt(c echo.Context) error {
	v := struct {
		Name string
		Age  int
	}{
		Name: "Kojima",
		Age:  96,
	}
	return c.Render(http.StatusOK, "kojima", v)
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
