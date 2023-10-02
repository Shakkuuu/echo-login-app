package controller

import (
	"echo-login-app/api/service"
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GachaController struct{}

// GET ガチャ
func (gc GachaController) DrawGacha(c echo.Context) error {
	var gs service.GachaService
	var is service.ItemService

	str_times := c.Param("times")
	times, err := strconv.Atoi(str_times)
	if err != nil {
		message := fmt.Sprintf("strconv.Atoi error: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// アイテム全取得処理
	item, err := is.GetAll()
	if err != nil {
		message := fmt.Sprintf("ItemService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// ガチャ
	result := gs.DrawGacha(times, item)

	return c.JSON(200, result)
}
