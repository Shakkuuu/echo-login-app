package main

import (
	"encoding/json"
	"io"
	"log"
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

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", Hello)
	e.GET("/map", Mappp)
	e.GET("struct", Structtt)
	e.GET("taroget", GetTaro)
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

type Taro struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetTaro(c echo.Context) error {
	url := "http://localhost:8081/taro"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
	}

	// openしたjsonを構造体にデコード
	var d Taro
	if err := json.Unmarshal(body, &d); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
	}

	return c.Render(http.StatusOK, "gettaro", d)
}
