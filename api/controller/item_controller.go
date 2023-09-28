package controller

import (
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type ItemController struct{}

// GET アイテム全取得
func (ic ItemController) GetAll(c echo.Context) error {
	var is service.ItemService

	// アイテム全取得処理
	i, err := is.GetAll()
	if err != nil {
		message := fmt.Sprintf("ItemService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, i)
}

// GET IDからアイテム取得
func (ic ItemController) GetByID(c echo.Context) error {
	id := c.Param("id")

	var is service.ItemService
	// IDからのアイテム取得処理
	i, err := is.GetByID(id)
	if err != nil {
		message := fmt.Sprintf("ItemService.GetByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, i)
}

// DELETE アイテムの削除
func (ic ItemController) Delete(c echo.Context) error {
	id := c.Param("id")

	var is service.ItemService

	// アイテム削除処理
	err := is.Delete(id)
	if err != nil {
		message := fmt.Sprintf("ItemService.Delete: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	m := ResMess{Status: 200, Message: "Item Deleted: " + id}

	return c.JSON(200, m)
}
