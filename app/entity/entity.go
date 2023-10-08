package entity

import "time"

// ユーザー
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// メモ
type Memo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdat"`
	User_ID   int       `json:"user_id"`
}

// コイン
type Coin struct {
	ID      int `json:"id"`
	Qty     int `json:"qty"`
	User_ID int `json:"user_id"`
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

// アイテム
type Item struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Rarity Rarity `json:"rarity"`
	Ratio  int    `json:"ratio"`
}

// 所有済みアイテム
type HasItem struct {
	// ID      int    `json:"id"`
	Items   []Item `json:"items"`
	User_ID int    `json:"user_id"`
}

// アイテム表示用
type ShowItems struct {
	Item Item
	Qty  int
}

// シューティングゲーム用ユーザーのステータス
type Status struct {
	ID        int `json:"id"`
	Damage    int `json:"damage"`
	Hp        int `json:"hp"`
	ShotSpeed int `json:"shot_speed"`
	EnmCool   int `json:"enm_cool"`
	Score     int `json:"score"`
	User_ID   int `json:"user_id"`
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
