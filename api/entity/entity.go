package entity

import "time"

// ユーザー
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

// レスポンスメッセージ用構造体
type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
