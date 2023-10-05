package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type HasItemService struct{}

// 全取得済みアイテム取得処理
func (hs HasItemService) GetAll() ([]entity.HasItemList, error) {
	db := db.GetDB()
	var hasitemlist []entity.HasItemList

	err := db.Find(&hasitemlist).Error
	if err != nil {
		return hasitemlist, err
	}

	return hasitemlist, nil
}

// 取得済みアイテムリスト追加処理
func (hs HasItemService) Add(hasitemlist *entity.HasItemList) (*entity.HasItemList, error) {
	db := db.GetDB()

	err := db.Create(&hasitemlist).Error
	if err != nil {
		return hasitemlist, err
	}

	return hasitemlist, nil
}

// ユーザーIDからの取得済みアイテムの取得処理
func (hs HasItemService) GetByUserID(user_id string) ([]entity.HasItemList, error) {
	db := db.GetDB()
	var hasitemlist []entity.HasItemList

	err := db.Where("user_id = ?", user_id).Find(&hasitemlist).Error
	if err != nil {
		return hasitemlist, err
	}

	return hasitemlist, nil
}

// // UserIDからの取得済みアイテム更新処理
// func (hs HasItemService) PutByUserID(hasitem *entity.HasItem, user_id string) (*entity.HasItem, error) {
// 	db := db.GetDB()

// 	// err := db.Where("user_id = ?", user_id).Model(&hasitem).Preload("Items").Updates(&hasitem).Error
// 	err := db.Where("user_id = ?", user_id).Model(&hasitem).Updates(&hasitem).Error
// 	if err != nil {
// 		return hasitem, err
// 	}

// 	return hasitem, nil
// }

// Item_IDからの取得済みアイテムの最初の削除処理
func (hs HasItemService) Delete(id string) error {
	db := db.GetDB()

	var hasitemlist entity.HasItemList

	err := db.Where("id = ?", id).First(&hasitemlist).Delete(&hasitemlist).Error
	if err != nil {
		return err
	}

	return nil
}
