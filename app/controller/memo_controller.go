package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type MemoController struct{}

// GET Topページ表示
func (mc MemoController) Top(c echo.Context) error {
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
	user_id := sess.Values["ID"].(int)

	var ms service.MemoService
	// ユーザー全取得
	u, err := ms.GetByUserID(user_id)
	if err != nil {
		log.Println("us.GetAll error")
	}
	return c.Render(http.StatusOK, "memotop.html", u)
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
	user_id := sess.Values["ID"].(int)

	var ms service.MemoService

	// メモ作成
	err = ms.Create(title, content, user_id)
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
