package entity

// ユーザー
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// レスポンスメッセージ用構造体
type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
