package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type StatusService struct{}

// 全ステータス取得処理
func (ss StatusService) GetAll() ([]entity.Status, error) {
	db := db.GetDB()
	var status []entity.Status

	err := db.Find(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}

// ユーザーのステータス一覧作成処理
func (ss StatusService) Create(status *entity.Status) (*entity.Status, error) {
	db := db.GetDB()

	err := db.Create(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}

// ユーザーIDからのステータスの全取得処理
func (ss StatusService) GetByUserID(user_id string) (entity.Status, error) {
	db := db.GetDB()
	var status entity.Status

	err := db.Where("user_id = ?", user_id).Find(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}

// User_IDからのステータス更新処理
func (ss StatusService) PutByUserID(status *entity.Status, id string) (*entity.Status, error) {
	db := db.GetDB()

	err := db.Where("user_id = ?", id).Model(&status).Updates(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}
