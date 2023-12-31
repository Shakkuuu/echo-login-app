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
	var cs service.CoinService
	var ss service.StatusService
	var hs service.HasItemService

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

	// ユーザー作成時にコインとユーザーのステータス一覧の作成と、アイテム1つ付与
	coin := entity.Coin{User_ID: user.ID}

	_, err = cs.Create(&coin)
	if err != nil {
		message := fmt.Sprintf("CoinService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	status := entity.Status{Damage: 1, Hp: 10, ShotSpeed: 20, EnmCool: 100, Score: 1, User_ID: user.ID}

	_, err = ss.Create(&status)
	if err != nil {
		message := fmt.Sprintf("StatusService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	hasitem := entity.HasItemList{UserID: user.ID, ItemID: 201}

	_, err = hs.Add(&hasitem)
	if err != nil {
		message := fmt.Sprintf("HasItemService.Add: %v", err)
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
