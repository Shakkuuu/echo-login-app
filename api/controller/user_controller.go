package controller

import (
	"echo-login-app/api/entity"
	"echo-login-app/api/service"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

type ErrRes entity.ErrorResponse

type UserController struct{}

func (uc UserController) GetAll(c echo.Context) error {
	var us service.UserService

	u, err := us.GetAll()
	if err != nil {
		message := fmt.Sprintf("UserService.GetAll: %d", err)
		log.Printf(message)
		e := ErrRes{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(200, u)
}

func (uc UserController) Create(c echo.Context) error {
	var us service.UserService

	var u entity.User
	err := c.Bind(&u)
	if err != nil {
		message := fmt.Sprintf("User Create Bind: %d", err)
		log.Printf(message)
		e := ErrRes{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	user, err := us.Create(&u)
	if err != nil {
		message := fmt.Sprintf("UserService.Create: %d", err)
		log.Printf(message)
		e := ErrRes{Status: 500, Message: message}
		return c.JSON(e.Status, e)
	}

	return c.JSON(201, user)
}
