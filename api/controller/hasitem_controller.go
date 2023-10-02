package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type HasItemController struct{}

// GET 全ユーザーの取得済みアイテムリスト全取得
func (hc HasItemController) GetAll(c echo.Context) error {
	var hs service.HasItemService

	// 全ユーザーの取得済みアイテムリスト全取得処理
	hasitem, err := hs.GetAll()
	if err != nil {
		message := fmt.Sprintf("HasItemService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, hasitem)
}

// POST 取得済みアイテムリスト作成
func (hc HasItemController) Create(c echo.Context) error {
	var hs service.HasItemService

	var hasitem entity.HasItem
	// JSONをGoのデータに変換
	err := c.Bind(&hasitem)
	if err != nil {
		message := fmt.Sprintf("HasItem Create Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// 取得済みアイテムリスト作成処理
	reshasitem, err := hs.Create(&hasitem)
	if err != nil {
		message := fmt.Sprintf("HasItemService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, reshasitem)
}

// GET IDから取得済みアイテムリスト取得
func (hc HasItemController) GetByID(c echo.Context) error {
	id := c.Param("id")

	var hs service.HasItemService
	// IDからの取得済みアイテムリスト取得処理
	hasitem, err := hs.GetByID(id)
	if err != nil {
		message := fmt.Sprintf("HasItemService.GetByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, hasitem)
}

// GET ユーザーIDから取得済みアイテムリスト取得
func (hc HasItemController) GetByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var hs service.HasItemService
	// ユーザーIDから取得済みアイテムリスト取得処理
	hasietm, err := hs.GetByUserID(user_id)
	if err != nil {
		message := fmt.Sprintf("HasItemService.GetByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, hasietm)
}

// PUT User_IDから取得済みアイテムリスト更新
func (hc HasItemController) PutByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var hs service.HasItemService

	// JSONをGoのデータに変換
	var hasitem entity.HasItem
	err := c.Bind(&hasitem)
	if err != nil {
		message := fmt.Sprintf("HasItem Update Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	print(&hasitem)

	// ユーザーIDから取得済みアイテムリスト更新処理
	hi, err := hs.PutByUserID(&hasitem, user_id)
	if err != nil {
		message := fmt.Sprintf("HasItemService.PutByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, hi)
}

// DELETE 取得済みアイテムリストの削除
func (hc HasItemController) Delete(c echo.Context) error {
	user_id := c.Param("user_id")

	var hs service.HasItemService

	// 取得済みアイテムリスト削除処理
	err := hs.Delete(user_id)
	if err != nil {
		message := fmt.Sprintf("HasItemService.Delete: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	m := ResMess{Status: 200, Message: " HasItem Deleted: " + user_id}

	return c.JSON(200, m)
}
