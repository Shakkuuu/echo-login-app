package controller

import (
	"echo-login-app/app/entity"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthController struct{}

// ログインしてセッションがあるか確認するミドルウェア
func (auc AuthController) SessionCheck(next echo.HandlerFunc) echo.HandlerFunc {
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
		// セッションが有効化されているか
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

func (auc AuthController) TokenGet(c echo.Context) (string, error) {
	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		return "", err
	}
	if id, ok := sess.Values["token"].(string); ok != true {
		log.Printf("不明なIDが保存されています: %v\n", id)
		return "", err
	}
	token := sess.Values["token"].(string)
	return token, nil
}

func (auc AuthController) SessionCreate(c echo.Context, u entity.User, token *entity.Token) error {
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		return err
	}
	// セッション作成
	sess.Options = &sessions.Options{
		MaxAge:   600,
		HttpOnly: true,
	}
	// セッションに値入れ
	sess.Values["ID"] = u.ID
	sess.Values["auth"] = true
	sess.Values["token"] = token.Token
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil
}

func (auc AuthController) IDGetBySession(c echo.Context) (int, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		return 0, nil
	}
	if id, ok := sess.Values["ID"].(int); ok != true {
		log.Printf("不明なIDが保存されています: %v\n", id)
		return 0, nil
	}
	id := sess.Values["ID"].(int)

	return id, nil
}
