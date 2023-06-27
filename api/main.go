package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	e.GET("/taro", Taro)
	e.Logger.Fatal(e.Start(":8081"))
}

func Taro(c echo.Context) error {
	d := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Taro",
		Age:  1200,
	}
	return c.JSON(http.StatusOK, d)
}
