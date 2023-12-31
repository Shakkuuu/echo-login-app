package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppController struct{}

// GET ログイン後のTopページ表示
func (apc AppController) Top(c echo.Context) error {
	// セッション
	var auc AuthController
	user_id, err := auc.IDGetBySession(c)
	if err != nil {
		log.Printf("auc.IDGetSession error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	var us service.UserService
	// セッションのIDからユーザーデータを取得
	u, err := us.GetByID(user_id)
	m := map[string]interface{}{
		"message": u.Name + "さんこんにちは!!!",
	}
	fmt.Println(m["message"])
	return c.Render(http.StatusOK, "top.html", m)
}
