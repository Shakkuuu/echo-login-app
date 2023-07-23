package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AppController struct{}

// GET ログイン後のTopページ表示
func (apc AppController) Top(c echo.Context) error {
	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	if id, ok := sess.Values["ID"].(int); ok != true {
		log.Printf("不明なIDが保存されています: %v\n", id)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	id := sess.Values["ID"].(int)
	var us service.UserService
	// セッションのIDからユーザーデータを取得
	u, err := us.GetByID(id)
	m := map[string]interface{}{
		"message": u.Name + "さんこんにちは!!!",
	}
	fmt.Println(m["message"])
	return c.Render(http.StatusOK, "top.html", m)
}
