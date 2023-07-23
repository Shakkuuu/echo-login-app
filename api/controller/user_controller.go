package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type ResMess entity.ResponseMessage

type UserController struct{}

// GET ユーザー全取得
func (uc UserController) GetAll(c echo.Context) error {
	var us service.UserService

	// ユーザー全取得処理
	u, err := us.GetAll()
	if err != nil {
		message := fmt.Sprintf("UserService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

// POST ユーザー作成
func (uc UserController) Create(c echo.Context) error {
	var us service.UserService

	var u entity.User
	// JSONをGoのデータに変換
	err := c.Bind(&u)
	if err != nil {
		message := fmt.Sprintf("User Create Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// ユーザー作成処理
	user, err := us.Create(&u)
	if err != nil {
		message := fmt.Sprintf("UserService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, user)
}

// GET IDからユーザーデータ取得
func (uc UserController) GetByID(c echo.Context) error {
	id := c.Param("id")

	var us service.UserService
	// IDからのユーザー取得処理
	u, err := us.GetByID(id)
	if err != nil {
		message := fmt.Sprintf("UserService.GetByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

// GET 名前からユーザーデータ取得
func (uc UserController) GetByName(c echo.Context) error {
	username := c.Param("username")

	var us service.UserService
	// 名前からユーザーデータ取得処理
	u, err := us.GetByName(username)
	if err != nil {
		message := fmt.Sprintf("UserService.GetByName: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

// PUT IDからユーザデータ更新
func (uc UserController) PutByID(c echo.Context) error {
	id := c.Param("id")

	var us service.UserService

	// JSONをGoのデータに変換
	var u entity.User
	err := c.Bind(&u)
	if err != nil {
		message := fmt.Sprintf("User Update Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	print(&u)

	// IDからユーザーデータ更新処理
	user, err := us.PutByID(&u, id)
	if err != nil {
		message := fmt.Sprintf("UserService.PutByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, user)
}

// DELETE ユーザーの削除
func (uc UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	var us service.UserService

	// ユーザー削除処理
	err := us.Delete(id)
	if err != nil {
		message := fmt.Sprintf("UserService.Delete: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	m := ResMess{Status: 200, Message: "User Deleted: " + id}

	return c.JSON(200, m)
}

// POST ユーザーログイン
func (uc UserController) Login(c echo.Context) error {
	var us service.UserService

	var u entity.User
	// JSONをGoのデータに変換
	err := c.Bind(&u)
	if err != nil {
		message := fmt.Sprintf("User Login Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// ログイン処理
	err = us.Login(&u)
	if err != nil {
		message := fmt.Sprintf("UserService.Login: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// m := ResMess{Status: 200, Message: "ログイン成功しました。"}

	// Tokens作成処理
	t, err := us.TokenCreate(u.ID)
	if err != nil || t == "" {
		message := fmt.Sprintf("us.TokenCreate: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	jtoken := entity.Token{Token: t}

	return c.JSON(200, jtoken)
}

// // Token確認処理
// func TokenCheck(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }

// // 署名の検証
// token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
// 	b := []byte(secretKey)
// 	return b, nil
// })
// if err != nil {
// 	return "", err
// }

// claims := token.Claims.(jwt.MapClaims)
// uname := fmt.Sprint(claims["user"])
// msg := fmt.Sprintf("こんにちは、「 %s 」さん", uname)
// fmt.Println(msg)

// return uname, nil
