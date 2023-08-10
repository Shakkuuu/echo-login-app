package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type CoinService struct{}

// コイン全取得処理
func (cs CoinService) GetAll() ([]entity.Coin, error) {
	db := db.GetDB()
	var c []entity.Coin

	err := db.Find(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}

// コイン作成処理
func (cs CoinService) Create(c *entity.Coin) (*entity.Coin, error) {
	db := db.GetDB()

	err := db.Create(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}

// IDからのコイン取得処理
func (cs CoinService) GetByID(id string) (entity.Coin, error) {
	db := db.GetDB()
	var c entity.Coin

	err := db.Where("id = ?", id).First(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}

// ユーザーIDからのコインの取得処理
func (cs CoinService) GetByUserID(user_id string) (entity.Coin, error) {
	db := db.GetDB()
	var c entity.Coin

	err := db.Where("user_id = ?", user_id).Find(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}

// UserIDからのコイン更新処理
func (cs CoinService) AddByUserID(c *entity.Coin, user_id string) (*entity.Coin, error) {
	db := db.GetDB()

	err := db.Where("user_id = ?", user_id).Model(&c).Updates(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}

// User_IDからのコイン削除処理
func (cs CoinService) Delete(user_id string) error {
	db := db.GetDB()

	var c entity.Coin

	err := db.Where("user_id = ?", user_id).Delete(&c).Error
	if err != nil {
		return err
	}

	return nil
}
