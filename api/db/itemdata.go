package db

import (
	"echo-login-app/api/entity"
	"log"
)

// アイテムデータ
var (
	Right     = entity.Item{ID: 101, Name: "右アタッチメント", Rarity: entity.RarityUR, Ratio: 5}
	Left      = entity.Item{ID: 102, Name: "左アタッチメント", Rarity: entity.RarityUR, Ratio: 5}
	Down      = entity.Item{ID: 103, Name: "下アタッチメント", Rarity: entity.RaritySSR, Ratio: 17}
	UpRight   = entity.Item{ID: 104, Name: "右上アタッチメント", Rarity: entity.RarityLR, Ratio: 5}
	UpLeft    = entity.Item{ID: 105, Name: "左上アタッチメント", Rarity: entity.RarityLR, Ratio: 5}
	DownRight = entity.Item{ID: 106, Name: "右下アタッチメント", Rarity: entity.RaritySSR, Ratio: 17}
	DownLeft  = entity.Item{ID: 107, Name: "左下アタッチメント", Rarity: entity.RaritySSR, Ratio: 17}

	Damage    = entity.Item{ID: 201, Name: "ダメージアップの素材", Rarity: entity.RarityN, Ratio: 1000}
	Hp        = entity.Item{ID: 202, Name: "HPアップの素材", Rarity: entity.RarityN, Ratio: 1000}
	ShotSpeed = entity.Item{ID: 203, Name: "連射速度アップの素材", Rarity: entity.RarityR, Ratio: 300}
	EnmCool   = entity.Item{ID: 204, Name: "敵増加間隔ダウンの素材", Rarity: entity.RarityR, Ratio: 300}
	Score     = entity.Item{ID: 205, Name: "スコアアップの素材", Rarity: entity.RaritySR, Ratio: 110}

	itemlist = []entity.Item{Right, Left, Down, UpRight, UpLeft, DownRight, DownLeft, Damage, Hp, ShotSpeed, EnmCool, Score}
)

// dbに初期アイテムデータ追加
func itemdatainsert() {
	db := GetDB()
	for _, v := range itemlist {
		err := db.Create(&v).Error
		if err != nil {
			log.Printf("itemID: %v, db Init error itemdatalist insert: %v\n", v.ID, err)
		}
	}
	log.Println("初期データが登録されました。")
}
