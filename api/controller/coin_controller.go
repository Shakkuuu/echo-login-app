package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type CoinController struct{}

// GET コイン全取得
func (cc CoinController) GetAll(c echo.Context) error {
	var cs service.CoinService

	// コイン全取得処理
	coin, err := cs.GetAll()
	if err != nil {
		message := fmt.Sprintf("CoinService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, coin)
}

// POST コイン作成
func (cc CoinController) Create(c echo.Context) error {
	var cs service.CoinService

	var coin entity.Coin
	// JSONをGoのデータに変換
	err := c.Bind(&coin)
	if err != nil {
		message := fmt.Sprintf("Coin Create Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// コイン作成処理
	rescoin, err := cs.Create(&coin)
	if err != nil {
		message := fmt.Sprintf("CoinService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, rescoin)
}

// GET ユーザーIDからコイン取得
func (cc CoinController) GetByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var cs service.CoinService
	// ユーザーIDからコイン取得処理
	coin, err := cs.GetByUserID(user_id)
	if err != nil {
		message := fmt.Sprintf("CoinService.GetByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, coin)
}

// PUT User_IDからコイン更新
func (cc CoinController) PutByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var cs service.CoinService

	// JSONをGoのデータに変換
	var coin entity.Coin
	err := c.Bind(&coin)
	if err != nil {
		message := fmt.Sprintf("Coin Update Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	print(&coin)

	// ユーザーIDからコイン更新処理
	co, err := cs.PutByUserID(&coin, user_id)
	if err != nil {
		message := fmt.Sprintf("CoinService.PutByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, co)
}
