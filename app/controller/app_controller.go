package controller

import (
	"echo-login-app/backend/service"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SessionCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// セッション
		sess, err := session.Get("session", c)
		if err != nil {
			log.Printf("session.Get error: %v\n", err)
			m := map[string]interface{}{
				"message": "セッションの取得に失敗しました。もう一度お試しください。",
			}
			return c.Render(http.StatusBadRequest, "login.html", m)
		}
		if sess.Values["auth"] != true {
			m := map[string]interface{}{
				"message": "ログインをしてください。",
			}
			return c.Render(http.StatusOK, "login.html", m)
		}
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

type AppController struct{}

func (ac AppController) Top(c echo.Context) error {
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
	u, err := us.GetByID(id)
	m := map[string]interface{}{
		"message": u.Name + "さんこんにちは!!!",
	}
	fmt.Println(m["message"])
	return c.Render(http.StatusOK, "top.html", m)
}
