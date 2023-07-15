package controller

import (
	"echo-login-app/backend/service"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
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

	err = us.Login(u.ID, username, password)
	if err != nil {
		log.Println("us.Login error")
		m := map[string]interface{}{
			"message": err,
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	fmt.Println("ログイン成功したよ")
	return c.Redirect(http.StatusFound, "/")
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

	err = us.Create(id, username, password)
	if err != nil {
		log.Println("us.Create error")
		m := map[string]interface{}{
			"message": "作成時にエラーが発生しました。",
		}
		return c.Render(http.StatusBadRequest, "signup.html", m)
	}

	fmt.Println("ユーザー登録成功したよ")
	return c.Redirect(http.StatusFound, "/")
}
