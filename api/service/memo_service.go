package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
)

type MemoService struct{}

// メモ全取得処理
func (ms MemoService) GetAll() ([]entity.Memo, error) {
	db := db.GetDB()
	var m []entity.Memo

	err := db.Find(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// メモ作成処理
func (ms MemoService) Create(m *entity.Memo) (*entity.Memo, error) {
	db := db.GetDB()

	err := db.Create(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// IDからのメモ取得処理
func (ms MemoService) GetByID(id string) (entity.Memo, error) {
	db := db.GetDB()
	var m entity.Memo

	err := db.Where("id = ?", id).First(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// ユーザーIDからのメモの全取得処理
func (ms MemoService) GetByUserID(user_id string) ([]entity.Memo, error) {
	db := db.GetDB()
	var m []entity.Memo

	err := db.Where("user_id = ?", user_id).First(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// IDからのメモ更新処理
func (ms MemoService) PutByID(m *entity.Memo, id string) (*entity.Memo, error) {
	db := db.GetDB()

	err := db.Where("id = ?", id).Model(&m).Updates(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// IDからのメモ削除処理
func (ms MemoService) Delete(id string) error {
	db := db.GetDB()

	var m entity.Memo

	err := db.Where("id = ?", id).Delete(&m).Error
	if err != nil {
		return err
	}

	return nil
}
