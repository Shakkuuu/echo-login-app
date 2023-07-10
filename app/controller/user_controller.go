package controller

import (
	"echo-login-app/backend/service"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type UserController struct{}

func (uc UserController) Index(c echo.Context) error {
	var us service.UserService
	u, err := us.GetAll()
	if err != nil {
		log.Println("us.GetAll error")
	}
	return c.Render(http.StatusOK, "index.html", u)
}
