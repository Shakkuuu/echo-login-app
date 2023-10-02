package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type HasItemService struct{}

// 全ユーザーの取得済みアイテム全取得処理
func (hs HasItemService) GetAll() ([]entity.HasItem, error) {
	db := db.GetDB()
	var hasitem []entity.HasItem

	err := db.Find(&hasitem).Error
	if err != nil {
		return hasitem, err
	}

	return hasitem, nil
}

// 取得済みアイテムリスト作成処理
func (hs HasItemService) Create(hasitem *entity.HasItem) (*entity.HasItem, error) {
	db := db.GetDB()

	err := db.Create(&hasitem).Error
	if err != nil {
		return hasitem, err
	}

	return hasitem, nil
}

// IDからの取得済みアイテム取得処理
func (hs HasItemService) GetByID(id string) (entity.HasItem, error) {
	db := db.GetDB()
	var hasitem entity.HasItem

	err := db.Where("id = ?", id).First(&hasitem).Error
	if err != nil {
		return hasitem, err
	}

	return hasitem, nil
}

// ユーザーIDからの取得済みアイテムの取得処理
func (hs HasItemService) GetByUserID(user_id string) (entity.HasItem, error) {
	db := db.GetDB()
	var hasitem entity.HasItem

	err := db.Where("user_id = ?", user_id).Find(&hasitem).Error
	if err != nil {
		return hasitem, err
	}

	return hasitem, nil
}

// UserIDからの取得済みアイテム更新処理
func (hs HasItemService) PutByUserID(hasitem *entity.HasItem, user_id string) (*entity.HasItem, error) {
	db := db.GetDB()

	err := db.Where("user_id = ?", user_id).Model(&hasitem).Updates(&hasitem).Error
	if err != nil {
		return hasitem, err
	}

	return hasitem, nil
}

// User_IDからの取得済みアイテム削除処理
func (hs HasItemService) Delete(user_id string) error {
	db := db.GetDB()

	var hasitem entity.HasItem

	err := db.Where("user_id = ?", user_id).Delete(&hasitem).Error
	if err != nil {
		return err
	}

	return nil
}
