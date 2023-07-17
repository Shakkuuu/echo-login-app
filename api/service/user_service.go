package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type UserService struct{}

// ユーザー全取得処理
func (us UserService) GetAll() ([]entity.User, error) {
	db := db.GetDB()
	var u []entity.User

	err := db.Find(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// ユーザー作成処理
func (us UserService) Create(u *entity.User) (*entity.User, error) {
	db := db.GetDB()

	err := db.Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// IDからのユーザー取得処理
func (us UserService) GetByID(id string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// 名前からのユーザー取得処理
func (us UserService) GetByName(username string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	err := db.Where("name = ?", username).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// IDからのユーザーデータ更新処理
func (us UserService) PutByID(u *entity.User, id string) (*entity.User, error) {
	db := db.GetDB()

	err := db.Where("id = ?", id).Model(&u).Updates(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// IDからのユーザー削除処理
func (us UserService) Delete(id string) error {
	db := db.GetDB()

	var u entity.User

	err := db.Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}
