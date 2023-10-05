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

// コイン
type Coin struct {
	ID      int  `json:"id" gorm:"primaryKey"`
	Qty     *int `json:"qty"`
	User_ID int  `json:"user_id"`
}

type Rarity string

const (
	RarityN   Rarity = "N"
	RarityR   Rarity = "R"
	RaritySR  Rarity = "SR"
	RaritySSR Rarity = "SSR"
	RarityUR  Rarity = "UR"
	RarityLR  Rarity = "LR"
)

// // アイテム
// type Item struct {
// 	ID       int       `json:"id" gorm:"primaryKey"`
// 	Name     string    `json:"name"`
// 	Rarity   Rarity    `json:"rarity"`
// 	Ratio    int       `json:"ratio"`
// 	HasItems []HasItem `gorm:"many2many:hasitem_items;"`
// }

// type HasItem struct {
// 	ID      int    `json:"id" gorm:"primaryKey"`
// 	Items   []Item `json:"items" gorm:"many2many:hasitem_items;"`
// 	User_ID int    `json:"user_id"`
// }

// アイテム
type Item struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Rarity Rarity `json:"rarity"`
	Ratio  int    `json:"ratio"`
}

type HasItemList struct {
	ID     int `gorm:"primaryKey"`
	UserID int
	ItemID int
}

type HasItem struct {
	// ID      int    `json:"id"`
	Items   []Item `json:"items"`
	User_ID int    `json:"user_id"`
}

// レスポンスメッセージ用構造体
type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// トークン
type Token struct {
	Token string `json:"token"`
}
