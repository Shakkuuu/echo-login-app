package entity

import "time"

// ユーザー
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

// メモ
type Memo struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdat"`
	User_ID   int       `json:"user_id"`
}

// レスポンスメッセージ用構造体
type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
