package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

// GET indexページ表示
func (uc UserController) Index(c echo.Context) error {
	var us service.UserService
	// ユーザー全取得
	u, err := us.GetAll()
	if err != nil {
		log.Println("us.GetAll error")
	}
	return c.Render(http.StatusOK, "index.html", u)
}

// GET Loginページ表示
func (uc UserController) LoginView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "login.html", m)
}

// GET Signupページ表示
func (uc UserController) SignupView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "signup.html", m)
}

// POST Login処理
func (uc UserController) Login(c echo.Context) error {
	var us service.UserService

	// htmlのformから値取得
	username := c.FormValue("username")
	password := c.FormValue("password")

	// 入力漏れチェック
	if username == "" || password == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// ユーザー名が存在するかチェック
	ulist, err := us.GetAll()
	if err != nil {
		log.Println("us.GetAll error")
	}
	var count int = 0
	for _, v := range ulist {
		if v.Name == username {
			count++
		}
	}
	if count == 0 {
		log.Println("そのユーザー名は存在しません")
		m := map[string]interface{}{
			"message": "そのユーザー名は存在しません",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// ユーザー名からID取得
	u, err := us.GetByName(username)
	if err != nil {
		log.Println("ID取得時にエラーが発生しました。")
		m := map[string]interface{}{
			"message": "ID取得時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// Login処理 パスワードチェック
	token, err := us.Login(u.ID, password, true)
	if err != nil {
		log.Println("us.Login error")
		m := map[string]interface{}{
			"message": err,
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	fmt.Printf("ログイン処理後のToken: %v", token)

	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの確立に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
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
		log.Printf("session.Save error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの確立に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	fmt.Println("ログイン成功したよ")
	return c.Redirect(http.StatusFound, "/app")
}

// POST Signup処理
func (uc UserController) Signup(c echo.Context) error {
	var us service.UserService

	// htmlからformの取得
	username := c.FormValue("username")
	password := c.FormValue("password")
	checkpass := c.FormValue("checkpassword")

	// 入力漏れチェック
	if username == "" || password == "" || checkpass == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

	// 確認用再入力パスワードがあっているか
	if password != checkpass {
		log.Println("パスワードが一致していないよ。")
		m := map[string]interface{}{
			"message": "パスワードが一致していないよ。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

	// 既にユーザー名が使用されていないかチェック
	u, err := us.GetAll()
	if err != nil {
		log.Println("us.GetAll error")
	}
	for _, v := range u {
		if v.Name == username {
			log.Println("そのユーザー名は既に使われているよ")
			m := map[string]interface{}{
				"message": "そのユーザー名は既に使われているよ。",
			}
			return c.Render(http.StatusBadRequest, "signup.html", m)
		}
	}

	// ID生成
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100000000)

	// パスワードのハッシュ化
	hashp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword error")
		m := map[string]interface{}{
			"message": "ユーザー作成時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

	// ユーザー作成
	err = us.Create(id, username, string(hashp))
	if err != nil {
		log.Println("us.Create error")
		m := map[string]interface{}{
			"message": "ユーザー作成時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

	fmt.Println("ユーザー登録成功したよ")
	return c.Redirect(http.StatusFound, "/login")
}

// GET ログイン後のユーザーページ
func (uc UserController) UserPage(c echo.Context) error {
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
	// セッションに保存されているIDからユーザーデータの取得
	u, err := us.GetByID(id)
	if err != nil {
		log.Printf("service.GetByID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーデータの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	return c.Render(http.StatusOK, "userpage.html", u)
}

// GET ログアウト処理
func (uc UserController) Logout(c echo.Context) error {
	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// セッションの無効化
	sess.Values["auth"] = false
	sess.Options.MaxAge = -1
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		log.Printf("session.Save error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの削除に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	m := map[string]interface{}{
		"message": "ログアウトしたよ。",
	}
	return c.Render(http.StatusOK, "login.html", m)
}

// GET ユーザー名変更ページ
func (uc UserController) ChangeNameView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "userchangename.html", m)
}

// POST ユーザー名変更処理
func (uc UserController) ChangeName(c echo.Context) error {
	var us service.UserService

	// htmlのformから値の取得
	username := c.FormValue("username")

	// 入力漏れのチェック
	if username == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "userchangename.html", m)
	}

	// ユーザー名が既に使用されていないかチェック
	u, err := us.GetAll()
	if err != nil {
		log.Println("us.GetAll error")
	}
	for _, v := range u {
		if v.Name == username {
			log.Println("そのユーザー名は既に使われているよ")
			m := map[string]interface{}{
				"message": "そのユーザー名は既に使われているよ。",
			}
			return c.Render(http.StatusBadRequest, "userchangename.html", m)
		}
	}

	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "userchangename.html", m)
	}
	if id, ok := sess.Values["ID"].(int); ok != true {
		log.Printf("不明なIDが保存されています: %v\n", id)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "userchangename.html", m)
	}
	id := sess.Values["ID"].(int)

	// ユーザー名変更処理
	err = us.ChangeName(id, username)
	if err != nil {
		log.Println("us.ChangeName error")
		m := map[string]interface{}{
			"message": "ユーザー名変更時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "userchangename.html", m)
	}

	fmt.Println("ユーザー名変更成功したよ")
	return c.Redirect(http.StatusFound, "/app")
}

// GET ユーザーパスワード変更ページ
func (uc UserController) ChangePasswordView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "userchangepassword.html", m)
}

// POST ユーザーパスワード変更処理
func (uc UserController) ChangePassword(c echo.Context) error {
	var us service.UserService

	// htmlのformから値の取得
	oldpassword := c.FormValue("oldpassword")
	newpassword := c.FormValue("newpassword")
	newcheckpassword := c.FormValue("newcheckpassword")

	// 入力漏れのチェック
	if oldpassword == "" || newpassword == "" || newcheckpassword == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

	// 新しいパスワードと確認用再入力パスワードが一致しているかチェック
	if newpassword != newcheckpassword {
		log.Println("新しいパスワードが一致していないよ。")
		m := map[string]interface{}{
			"message": "新しいパスワードが一致していないよ。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}
	if id, ok := sess.Values["ID"].(int); ok != true {
		log.Printf("不明なIDが保存されています: %v\n", id)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}
	id := sess.Values["ID"].(int)

	// パスワードチェック
	_, err = us.Login(id, oldpassword, false)
	if err != nil {
		log.Println("us.Login error")
		m := map[string]interface{}{
			"message": err,
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

	// パスワードのハッシュ化
	hashp, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword error")
		m := map[string]interface{}{
			"message": "パスワード変更時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, ".html", m)
	}

	// パスワードの変更処理
	err = us.ChangePassword(id, string(hashp))
	if err != nil {
		log.Println("us.ChangePassword error")
		m := map[string]interface{}{
			"message": "パスワード変更時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

	fmt.Println("パスワードの変更成功したよ")
	m := map[string]interface{}{
		"message": "パスワードの変更成功したよ",
	}
	return c.Render(http.StatusFound, "userchangepassword.html", m)
}

// GET ユーザー削除処理
func (uc UserController) Delete(c echo.Context) error {
	var us service.UserService

	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "top.html", m)
	}
	if id, ok := sess.Values["ID"].(int); ok != true {
		log.Printf("不明なIDが保存されています: %v\n", id)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "top.html", m)
	}
	id := sess.Values["ID"].(int)

	// ユーザー削除処理
	err = us.Delete(id)
	if err != nil {
		log.Println("us.Delete error")
		m := map[string]interface{}{
			"message": "ユーザー削除時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "top.html", m)
	}

	// セッションの無効化
	sess.Values["auth"] = false
	sess.Options.MaxAge = -1
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		log.Printf("session.Save error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの削除に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	fmt.Println("ユーザーを削除しました")
	m := map[string]interface{}{
		"message": "ユーザーを削除しました。",
	}

	return c.Render(http.StatusFound, "login.html", m)
}
