package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type StatusController struct{}

// GET 全ユーザーのステータス一覧取得
func (sc StatusController) GetAll(c echo.Context) error {
	var ss service.StatusService

	// 全ユーザーのステータス一覧取得処理
	coin, err := ss.GetAll()
	if err != nil {
		message := fmt.Sprintf("StatusService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, coin)
}

// GET ユーザーIDからユーザーのステータス一覧取得
func (sc StatusController) GetByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var ss service.StatusService
	// ユーザーIDからユーザーのステータス一覧取得処理
	status, err := ss.GetByUserID(user_id)
	if err != nil {
		message := fmt.Sprintf("StatusService.GetByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, status)
}

// PUT User_IDからユーザーのステータス一覧更新
func (sc StatusController) PutByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var ss service.StatusService

	// JSONをGoのデータに変換
	var status entity.Status
	err := c.Bind(&status)
	if err != nil {
		message := fmt.Sprintf("Status Update Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	print(&status)

	// ユーザーIDからユーザーのステータス一覧更新処理
	co, err := ss.PutByUserID(&status, user_id)
	if err != nil {
		message := fmt.Sprintf("StatusService.PutByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, co)
}
