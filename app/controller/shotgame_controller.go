package controller

import (
	"echo-login-app/app/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShotGameController struct{}

// GET Topページ表示
func (gc ShotGameController) Top(c echo.Context) error {
	var auc AuthController
	var ss service.StatusService

	// セッション
	user_id, err := auc.IDGetBySession(c)
	if err != nil {
		log.Printf("auc.IDGetSession error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Printf("TokenGet error: %v\n", err)
		m := map[string]interface{}{
			"message": "Tokenの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// ユーザーのステータス一覧取得
	status, err := ss.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ss.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーのステータス一覧の取得に失敗しました。",
			"status":  nil,
		}
		return c.Render(http.StatusBadRequest, "shotgame.html", m)
	}

	m := map[string]interface{}{
		"message": "",
		"status":  status,
	}

	return c.Render(http.StatusOK, "shotgame.html", m)
}
