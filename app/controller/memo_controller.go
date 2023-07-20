package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	// メモ全取得
	u, err := ms.GetByUserID(user_id)
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

// GET メモの中身表示
func (mc MemoController) ContentView(c echo.Context) error {
	form_id := c.Param("id")
	id, err := strconv.Atoi(form_id)
	if err != nil {
		log.Println("strconv.Atoi error")
		m := map[string]interface{}{
			"message": "メモ取得時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	var ms service.MemoService
	// IDからメモ取得
	u, err := ms.GetByID(id)
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

	form_id := c.Param("id")
	id, err := strconv.Atoi(form_id)
	if err != nil {
		log.Println("strconv.Atoi error")
		m := map[string]interface{}{
			"message": "メモID取得時にエラーが発生しました。",
			"memo":    nil,
		}
		return c.Render(http.StatusBadRequest, "memotop.html", m)
	}

	// ユーザー削除処理
	err = ms.Delete(id)
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
