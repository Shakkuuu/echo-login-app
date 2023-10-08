package controller

import (
	"echo-login-app/app/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CoinController struct{}

// GET Topページ表示
func (cc CoinController) Top(c echo.Context) error {
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
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}
	m := map[string]interface{}{
		"message": "",
		"coin":    coin,
	}
	return c.Render(http.StatusOK, "cointop.html", m)
}

// POST コイン追加
func (cc CoinController) QtyAdd(c echo.Context) error {
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
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	// htmlからformの取得
	s_addqty := c.FormValue("addqty")
	addqty, err := strconv.Atoi(s_addqty)
	if err != nil {
		log.Println("strconv.Atoi error")
		m := map[string]interface{}{
			"message": "コイン追加枚数を正しく指定してください。",
			"coin":    coin,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	addcoin := coin.Qty + addqty

	// コイン数変更
	err = cs.ChangeQty(token, user_id, addcoin)
	if err != nil {
		log.Printf("cs.ChangeQty error: %v\n", err)
		m := map[string]interface{}{
			"message": "コインの増加に失敗しました。",
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	return c.Redirect(http.StatusFound, "/app/coin")
}

// GET コイン減少
func (cc CoinController) QtySub(c echo.Context) error {
	var auc AuthController
	var cs service.CoinService
	var subcoin int

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
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	// htmlからformの取得
	s_subqty := c.FormValue("subqty")
	subqty, err := strconv.Atoi(s_subqty)
	if err != nil {
		log.Println("strconv.Atoi error")
		m := map[string]interface{}{
			"message": "コイン減少枚数を正しく指定してください。",
			"coin":    coin,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	if coin.Qty <= 0 {
		subcoin = 0
	} else {
		subcoin = coin.Qty - subqty
	}

	// コイン数変更
	err = cs.ChangeQty(token, user_id, subcoin)
	if err != nil {
		log.Printf("cs.ChangeQty error: %v\n", err)
		m := map[string]interface{}{
			"message": "コインの減少に失敗しました。",
			"coin":    nil,
		}
		return c.Render(http.StatusBadRequest, "cointop.html", m)
	}

	return c.Redirect(http.StatusFound, "/app/coin")
}
