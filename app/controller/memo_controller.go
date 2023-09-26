package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MemoController struct{}

// GET Topページ表示
func (mc MemoController) Top(c echo.Context) error {
	var auc AuthController
	var ms service.MemoService

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

	// メモ全取得
	u, err := ms.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ms.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "メモの取得に失敗しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}
	m := map[string]interface{}{
		"message": "",
		"memo":    u,
	}
	return c.Render(http.StatusOK, "memotop.html", m)
}

// GET メモ作成ページ
func (mc MemoController) CreatePage(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "memocreate.html", m)
}

// POST メモ作成
func (mc MemoController) Create(c echo.Context) error {
	var auc AuthController
	var ms service.MemoService

	// htmlからformの取得
	title := c.FormValue("title")
	content := c.FormValue("content")

	// 入力漏れチェック
	if title == "" || content == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "memocreate.html", m)
	}

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

	// メモ作成
	err = ms.Create(title, content, user_id, token)
	if err != nil {
		log.Println("ms.Create error")
		m := map[string]interface{}{
			"message": "メモ作成時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "memocreate.html", m)
	}

	fmt.Println("メモ作成成功したよ")
	return c.Redirect(http.StatusFound, "/app/memo")
}

// GET メモの中身表示
func (mc MemoController) ContentView(c echo.Context) error {
	var auc AuthController
	var ms service.MemoService

	param_id := c.Param("id")
	id, err := strconv.Atoi(param_id)
	if err != nil {
		log.Println("strconv.Atoi error")
		m := map[string]interface{}{
			"message": "メモ取得時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Printf("TokenGet error: %v\n", err)
		m := map[string]interface{}{
			"message": "Tokenの取得に失敗しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	// IDからメモ取得
	u, err := ms.GetByID(id, token)
	if err != nil {
		log.Printf("ms.GetByID error: %v\n", err)
		m := map[string]interface{}{
			"message": "メモの取得に失敗しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	return c.Render(http.StatusOK, "memoview.html", u)
}

// GET メモ削除処理
func (mc MemoController) Delete(c echo.Context) error {
	var ms service.MemoService
	var auc AuthController

	id := c.Param("id")

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Println("TokenGet error")
		m := map[string]interface{}{
			"message": "Token取得時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	// メモ削除処理
	err = ms.Delete(id, token)
	if err != nil {
		log.Println("ms.Delete error")
		m := map[string]interface{}{
			"message": "メモ削除時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	fmt.Println("メモを削除しました")
	m := map[string]interface{}{
		"message": "メモを削除しました。",
		"memo":    nil,
	}

	return c.Render(http.StatusFound, "memotop.html", m)
}

// POST メモ変更処理
func (mc MemoController) Change(c echo.Context) error {
	var ms service.MemoService
	var auc AuthController

	param_id := c.Param("id")

	// htmlのformから値の取得
	title := c.FormValue("title")
	content := c.FormValue("content")

	// 入力漏れのチェック
	if title == "" || content == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Println("TokenGet error")
		m := map[string]interface{}{
			"message": "Token取得時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	// メモ変更処理
	err = ms.Change(param_id, title, content, token)
	if err != nil {
		log.Println("ms.Change error")
		m := map[string]interface{}{
			"message": "メモ変更時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	fmt.Println("ユーザー名変更成功したよ")
	return c.Redirect(http.StatusFound, "/app/memo/view/"+param_id)
}
