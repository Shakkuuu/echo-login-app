package service

import (
	"echo-login-app/backend/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type UserService struct{}

func (us UserService) GetAll() ([]entity.User, error) {
	var u []entity.User
	url := "http://echo-login-app-api:8081/user"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return u, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return u, err
	}

	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return u, err
	}

	return u, nil
}
