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

func (uc UserController) GetAll(c echo.Context) error {
	var us service.UserService

	u, err := us.GetAll()
	if err != nil {
		message := fmt.Sprintf("UserService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

func (uc UserController) Create(c echo.Context) error {
	var us service.UserService

	var u entity.User
	err := c.Bind(&u)
	if err != nil {
		message := fmt.Sprintf("User Create Bind: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	user, err := us.Create(&u)
	if err != nil {
		message := fmt.Sprintf("UserService.Create: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, user)
}

func (uc UserController) GetByID(c echo.Context) error {
	id := c.Param("id")

	var us service.UserService

	u, err := us.GetByID(id)
	if err != nil {
		message := fmt.Sprintf("UserService.GetByID: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

func (uc UserController) GetByName(c echo.Context) error {
	id := c.Param("username")

	var us service.UserService

	u, err := us.GetByName(id)
	if err != nil {
		message := fmt.Sprintf("UserService.GetByName: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

func (uc UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	var us service.UserService

	err := us.Delete(id)
	if err != nil {
		message := fmt.Sprintf("UserService.GetAll: %v", err)
		log.Println(message)
		e := ResMess{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	m := ResMess{Status: 200, Message: "User Deleted: " + id}

	return c.JSON(200, m)
}
