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

func (us UserService) GetByID(id string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (us UserService) GetByName(username string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	err := db.Where("name = ?", username).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (us UserService) Delete(id string) error {
	db := db.GetDB()

	var u entity.User

	err := db.Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}
