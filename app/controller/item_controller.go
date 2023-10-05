package controller

import (
	"echo-login-app/app/entity"
	"echo-login-app/app/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemController struct{}

// GET Topページ表示
func (ic ItemController) Top(c echo.Context) error {
	var auc AuthController
	var hs service.HasItemService

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

	// 所持済みアイテム取得
	hasitem, err := hs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ms.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "所持済みアイテムの取得に失敗しました。",
			"items":   nil,
		}
		return c.Render(http.StatusBadRequest, "itemtop.html", m)
	}

	// 重複カウント用map
	counts := make(map[entity.Item]int)
	for _, v := range hasitem.Items {
		counts[v]++
	}

	// 改めて格納
	var items []entity.ShowItems
	for key, value := range counts {
		item := entity.ShowItems{
			Item: key,
			Qty:  value,
		}
		items = append(items, item)
	}

	m := map[string]interface{}{
		"message": "",
		"items":   items,
	}
	return c.Render(http.StatusOK, "itemtop.html", m)
}
