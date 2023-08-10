package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type CoinService struct{}

// コイン全取得処理
func (cs CoinService) GetAll() ([]entity.Coin, error) {
	db := db.GetDB()
	var coin []entity.Coin

	err := db.Find(&coin).Error
	if err != nil {
		return coin, err
	}

	return coin, nil
}

// コイン作成処理
func (cs CoinService) Create(coin *entity.Coin) (*entity.Coin, error) {
	db := db.GetDB()

	err := db.Create(&coin).Error
	if err != nil {
		return coin, err
	}

	return coin, nil
}

// IDからのコイン取得処理
func (cs CoinService) GetByID(id string) (entity.Coin, error) {
	db := db.GetDB()
	var coin entity.Coin

	err := db.Where("id = ?", id).First(&coin).Error
	if err != nil {
		return coin, err
	}

	return coin, nil
}

// ユーザーIDからのコインの取得処理
func (cs CoinService) GetByUserID(user_id string) (entity.Coin, error) {
	db := db.GetDB()
	var coin entity.Coin

	err := db.Where("user_id = ?", user_id).Find(&coin).Error
	if err != nil {
		return coin, err
	}

	return coin, nil
}

// UserIDからのコイン更新処理
func (cs CoinService) AddByUserID(coin *entity.Coin, user_id string) (*entity.Coin, error) {
	db := db.GetDB()

	err := db.Where("user_id = ?", user_id).Model(&coin).Updates(&coin).Error
	if err != nil {
		return coin, err
	}

	return coin, nil
}

// User_IDからのコイン削除処理
func (cs CoinService) Delete(user_id string) error {
	db := db.GetDB()

	var coin entity.Coin

	err := db.Where("user_id = ?", user_id).Delete(&coin).Error
	if err != nil {
		return err
	}

	return nil
}
