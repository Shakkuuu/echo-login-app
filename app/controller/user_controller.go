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

func (uc UserController) Index(c echo.Context) error {
	var us service.UserService
	u, err := us.GetAll()
	if err != nil {
		log.Println("us.GetAll error")
	}
	return c.Render(http.StatusOK, "index.html", u)
}

func (uc UserController) LoginView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "login.html", m)
}

func (uc UserController) SignupView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "signup.html", m)
}

func (uc UserController) Login(c echo.Context) error {
	var us service.UserService

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

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

	u, err := us.GetByName(username)
	if err != nil {
		log.Println("ID取得時にエラーが発生しました。")
		m := map[string]interface{}{
			"message": "ID取得時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	err = us.Login(u.ID, password)
	if err != nil {
		log.Println("us.Login error")
		m := map[string]interface{}{
			"message": err,
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// セッション
	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("session.Get error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの確立に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}
	sess.Options = &sessions.Options{
		MaxAge:   600,
		HttpOnly: true,
	}
	sess.Values["ID"] = u.ID
	sess.Values["auth"] = true
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

func (uc UserController) Signup(c echo.Context) error {
	var us service.UserService

	// id, _ := strconv.Atoi(c.FormValue("id"))
	username := c.FormValue("username")
	password := c.FormValue("password")
	checkpass := c.FormValue("checkpassword")

	if username == "" || password == "" || checkpass == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

	if password != checkpass {
		log.Println("パスワードが一致していないよ。")
		m := map[string]interface{}{
			"message": "パスワードが一致していないよ。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

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

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100000000)

	hashp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword error")
		m := map[string]interface{}{
			"message": "ユーザー作成時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

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

func (uc UserController) ChangeNameView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "userchangename.html", m)
}

func (uc UserController) ChangeName(c echo.Context) error {
	var us service.UserService

	username := c.FormValue("username")

	if username == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "userchangename.html", m)
	}

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
	user, err := us.GetByID(id)
	if err != nil {
		log.Printf("service.GetByID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーデータの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "userchangename.html", m)
	}

	err = us.ChangeName(user.ID, username)
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

func (uc UserController) ChangePasswordView(c echo.Context) error {
	m := map[string]interface{}{
		"message": "",
	}
	return c.Render(http.StatusOK, "userchangepassword.html", m)
}

func (uc UserController) ChangePassword(c echo.Context) error {
	var us service.UserService

	oldpassword := c.FormValue("oldpassword")
	newpassword := c.FormValue("newpassword")
	newcheckpassword := c.FormValue("newcheckpassword")

	if oldpassword == "" || newpassword == "" || newcheckpassword == "" {
		log.Println("入力されていない項目があるよ。")
		m := map[string]interface{}{
			"message": "入力されていない項目があるよ。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

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
	user, err := us.GetByID(id)
	if err != nil {
		log.Printf("service.GetByID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーデータの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

	err = us.Login(user.ID, oldpassword)
	if err != nil {
		log.Println("us.Login error")
		m := map[string]interface{}{
			"message": err,
		}
		return c.Render(http.StatusBadRequest, "userchangepassword.html", m)
	}

	hashp, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword error")
		m := map[string]interface{}{
			"message": "パスワード変更時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, ".html", m)
	}

	err = us.ChangePassword(user.ID, string(hashp))
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
	// return c.Redirect(http.StatusFound, "/app")
}

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
	user, err := us.GetByID(id)
	if err != nil {
		log.Printf("service.GetByID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーデータの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "top.html", m)
	}

	err = us.Delete(user.ID)
	if err != nil {
		log.Println("us.Delete error")
		m := map[string]interface{}{
			"message": "ユーザー削除時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "top.html", m)
	}

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
