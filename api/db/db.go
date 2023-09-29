package db

import (
	"fmt"
	"log"
	"time"

	"echo-login-app/api/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// データベースと接続
func Init(un string, up string, dbn string) {
	DBMS := "mysql"            // データベースの種類
	USER := un                 // ユーザー名
	PASS := up                 // パスワード
	PROTOCOL := "tcp(db:3307)" // 3307ポート
	DBNAME := dbn              // データベース名

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	// 接続できるまで一定回数リトライ
	count := 0
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 { // countgaが180になるまでリトライ
				fmt.Println("")
				log.Printf("db Init error: %v\n", err)
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	autoMigration()

	itemdatainsert()
}

// serviceでデータベースとやりとりする用
func GetDB() *gorm.DB {
	return db
}

// サーバ終了時にデータベースとの接続終了
func Close() {
	if err := db.Close(); err != nil {
		log.Printf("db Close error: %v\n", err)
		panic(err)
	}
}

// entityを参照してテーブル作成　マイグレーション
func autoMigration() {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Memo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&entity.Coin{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&entity.Item{})
}
