package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type UserService struct{}

func (us UserService) GetAll() ([]entity.User, error) {
	db := db.GetDB()
	var u []entity.User

	err := db.Find(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (us UserService) Create(u *entity.User) (*entity.User, error) {
	db := db.GetDB()

	err := db.Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
