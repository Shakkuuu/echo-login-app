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
func (cs CoinService) PutByUserID(coin *entity.Coin, user_id string) (*entity.Coin, error) {
	db := db.GetDB()

	err := db.Where("user_id = ?", user_id).Model(&coin).Updates(&coin).Error
	if err != nil {
		return coin, err
	}

	return coin, nil
}
