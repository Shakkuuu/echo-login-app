package controller

import (
	"echo-login-app/app/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	GACHARATE int = 300 // ガチャ一回に必要なコイン枚数
)

type GachaController struct{}

// GET Topページ表示
func (gc GachaController) Top(c echo.Context) error {
	var auc AuthController
	var cs service.CoinService

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

	// コイン取得
	coin, err := cs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("cs.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "コインの取得に失敗しました。",
			"result":  nil,
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "gachatop.html", m)
	}
	m := map[string]interface{}{
		"message": "",
		"result":  nil,
		"coin":    coin,
	}
	return c.Render(http.StatusOK, "gachatop.html", m)
}

// POST ガチャdraw
func (gc GachaController) Draw(c echo.Context) error {
	var auc AuthController
	var gs service.GachaService
	var cs service.CoinService
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

	// コイン取得
	coin, err := cs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("cs.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "コインの取得に失敗しました。",
			"result":  nil,
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "gachatop.html", m)
	}

	// htmlからformの取得
	stimes := c.FormValue("times")
	times, err := strconv.Atoi(stimes)
	if err != nil {
		log.Println("strconv.Atoi error")
		m := map[string]interface{}{
			"message": "ガチャを引く回数を正しく指定してください。",
			"result":  nil,
			"coin":    coin,
		}
		return c.Render(http.StatusBadRequest, "gachatop.html", m)
	}

	subcoin := coin.Qty - times*GACHARATE
	// subcoin := coin.Qty - times
	if subcoin < 0 {
		log.Printf("gs.Draw error: %v\n", err)
		m := map[string]interface{}{
			"message": "コインが足りません",
			"result":  nil,
			"coin":    coin,
		}
		return c.Render(http.StatusBadRequest, "gachatop.html", m)
	}

	// ガチャを引く
	result, err := gs.Draw(times, token)
	if err != nil {
		log.Printf("gs.Draw error: %v\n", err)
		m := map[string]interface{}{
			"message": "ガチャのDrawに失敗しました。",
			"result":  nil,
			"coin":    coin,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	// 所持アイテムリストに追加
	for _, item := range result {
		err = hs.Add(token, user_id, item)
		if err != nil {
			log.Printf("hs.Change error: %v\n", err)
			m := map[string]interface{}{
				"message": "所持アイテムリスト追加に失敗しました。",
				"result":  result,
				"coin":    coin,
			}
			return c.Render(http.StatusBadRequest, "gachatop.html", m)
		}
	}

	// コイン消費
	err = cs.ChangeQty(token, user_id, subcoin)
	if err != nil {
		log.Printf("cs.ChangeQty error: %v\n", err)
		m := map[string]interface{}{
			"message": "コインの消費に失敗しました。",
			"result":  result,
			"coin":    coin,
		}
		return c.Render(http.StatusBadRequest, "gachatop.html", m)
	}

	// 減少後コイン取得
	coin, err = cs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("cs.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "消費後のコインの取得に失敗しました。",
			"result":  result,
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "gachatop.html", m)
	}

	m := map[string]interface{}{
		"message": "ガチャを引きました",
		"result":  result,
		"coin":    coin,
	}
	return c.Render(http.StatusOK, "gachatop.html", m)
}
