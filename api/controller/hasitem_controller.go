package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HasItemController struct{}

// GET 全ユーザーの取得済みアイテムリスト全取得
// func (hc HasItemController) GetAll(c echo.Context) error {
// 	var hs service.HasItemService

// 	// 全ユーザーの取得済みアイテムリスト全取得処理
// 	hasitemlist, err := hs.GetAll()
// 	if err != nil {
// 		message := fmt.Sprintf("HasItemService.GetAll: %v", err)
// 		log.Println(message)
// 		e := ResMess{Status: 500, Message: message}
// 		return c.JSON(e.Status, e)
// 	}

// 	return c.JSON(200, hasitem)
// }

// POST 取得済みアイテムリスト追加
func (hc HasItemController) Add(c echo.Context) error {
	user_id := c.Param("user_id")

	var hs service.HasItemService

	var item entity.Item
	// JSONをGoのデータに変換
	err := c.Bind(&item)
	if err != nil {
		message := fmt.Sprintf("HasItem Create Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	i_user_id, err := strconv.Atoi(user_id)
	if err != nil {
		message := fmt.Sprintf("strconv.Atoi error: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	hasitemlist := entity.HasItemList{UserID: i_user_id, ItemID: item.ID}

	// 取得済みアイテムリスト作成処理
	reshasitem, err := hs.Add(&hasitemlist)
	if err != nil {
		message := fmt.Sprintf("HasItemService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, reshasitem)
}

// GET ユーザーIDから取得済みアイテムリスト取得
func (hc HasItemController) GetByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var hs service.HasItemService
	var is service.ItemService

	var items []entity.Item

	// ユーザーIDから取得済みアイテムリスト取得処理
	hasitemlist, err := hs.GetByUserID(user_id)
	if err != nil {
		message := fmt.Sprintf("HasItemService.GetByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	for _, hil := range hasitemlist {
		// IDからアイテム取得処理
		item, err := is.GetByID(strconv.Itoa(hil.ItemID))
		if err != nil {
			message := fmt.Sprintf("ItemService.GetByID: %v", err)
			log.Println(message)
			e := ResMess{Status: 500, Message: message}
			return c.JSON(e.Status, e)
		}
		items = append(items, item)
	}

	i_user_id, err := strconv.Atoi(user_id)
	if err != nil {
		message := fmt.Sprintf("strconv.Atoi error: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	hasitem := entity.HasItem{Items: items, User_ID: i_user_id}

	return c.JSON(200, hasitem)
}

// // PUT User_IDから取得済みアイテムリスト更新
// func (hc HasItemController) PutByUserID(c echo.Context) error {
// 	user_id := c.Param("user_id")

// 	var hs service.HasItemService

// 	// JSONをGoのデータに変換
// 	var hasitem entity.HasItem
// 	err := c.Bind(&hasitem)
// 	if err != nil {
// 		message := fmt.Sprintf("HasItem Update Bind: %v", err)
// 		log.Println(message)
// 		e := ResMess{Status: 500, Message: message}
// 		return c.JSON(e.Status, e)
// 	}
// 	print(&hasitem)

// 	// ユーザーIDから取得済みアイテムリスト更新処理
// 	hi, err := hs.PutByUserID(&hasitem, user_id)
// 	if err != nil {
// 		message := fmt.Sprintf("HasItemService.PutByUserID: %v", err)
// 		log.Println(message)
// 		e := ResMess{Status: 500, Message: message}
// 		return c.JSON(e.Status, e)
// 	}

// 	return c.JSON(200, hi)
// }

// DELETE アイテムIDから取得済みアイテムリストの削除
func (hc HasItemController) Delete(c echo.Context) error {
	item_id := c.Param("item_id")

	var hs service.HasItemService

	// 取得済みアイテムリスト削除処理
	err := hs.Delete(item_id)
	if err != nil {
		message := fmt.Sprintf("HasItemService.Delete: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	m := ResMess{Status: 200, Message: " HasItem Deleted: " + item_id}

	return c.JSON(200, m)
}
