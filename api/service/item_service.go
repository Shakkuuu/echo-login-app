package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type ItemService struct{}

// アイテム全取得
func (is ItemService) GetAll() ([]entity.Item, error) {
	db := db.GetDB()
	var i []entity.Item

	err := db.Find(&i).Error
	if err != nil {
		return i, err
	}

	return i, nil
}

// IDからのアイテム取得
func (is ItemService) GetByID(id string) (entity.Item, error) {
	db := db.GetDB()
	var i entity.Item

	err := db.Where("id = ?", id).First(&i).Error
	if err != nil {
		return i, err
	}

	return i, nil
}

// IDからのアイテム削除処理
func (is ItemService) Delete(id string) error {
	db := db.GetDB()

	var i entity.Item

	err := db.Where("id = ?", id).Delete(&i).Error
	if err != nil {
		return err
	}

	return nil
}
