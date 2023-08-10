package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

// type ResMess entity.ResponseMessage

type MemoController struct{}

// GET メモ全取得
func (mc MemoController) GetAll(c echo.Context) error {
	var ms service.MemoService

	// メモ全取得処理
	m, err := ms.GetAll()
	if err != nil {
		message := fmt.Sprintf("MemoService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, m)
}

// POST メモ作成
func (mc MemoController) Create(c echo.Context) error {
	var ms service.MemoService

	var m entity.Memo
	// JSONをGoのデータに変換
	err := c.Bind(&m)
	if err != nil {
		message := fmt.Sprintf("Memo Create Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	// メモ作成処理
	memo, err := ms.Create(&m)
	if err != nil {
		message := fmt.Sprintf("MemoService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, memo)
}

// GET IDからメモ取得
func (mc MemoController) GetByID(c echo.Context) error {
	id := c.Param("id")

	var ms service.MemoService
	// IDからのメモ取得処理
	m, err := ms.GetByID(id)
	if err != nil {
		message := fmt.Sprintf("MemoService.GetByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, m)
}

// GET ユーザーIDからメモ取得
func (mc MemoController) GetByUserID(c echo.Context) error {
	user_id := c.Param("user_id")

	var ms service.MemoService
	// ユーザーIDからメモ取得処理
	m, err := ms.GetByUserID(user_id)
	if err != nil {
		message := fmt.Sprintf("MemoService.GetByUserID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, m)
}

// PUT IDからメモ更新
func (mc MemoController) PutByID(c echo.Context) error {
	id := c.Param("id")

	var ms service.MemoService

	// JSONをGoのデータに変換
	var m entity.Memo
	err := c.Bind(&m)
	if err != nil {
		message := fmt.Sprintf("Memo Update Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}
	print(&m)

	// IDからメモ更新処理
	memo, err := ms.PutByID(&m, id)
	if err != nil {
		message := fmt.Sprintf("MemoService.PutByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, memo)
}

// DELETE メモの削除
func (mc MemoController) Delete(c echo.Context) error {
	id := c.Param("id")

	var ms service.MemoService

	// メモ削除処理
	err := ms.Delete(id)
	if err != nil {
		message := fmt.Sprintf("MemoService.Delete: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	m := ResMess{Status: 200, Message: "Memo Deleted: " + id}

	return c.JSON(200, m)
}
