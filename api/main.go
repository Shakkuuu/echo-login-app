package main

import (
	"echo-login-app/api/db"
	"echo-login-app/api/server"
	"os"
)

func main() {
	// 環境変数読み込み
	un, up, dbn := loadEnv()
	db.Init(un, up, dbn)
	server.Init()

	db.Close()
}

// 環境変数読み込み
func loadEnv() (string, string, string) {
	// Docker-compose.ymlでDocker起動時に設定した環境変数の取得
	username := os.Getenv("USERNAME")
	userpass := os.Getenv("USERPASS")
	database := os.Getenv("DATABASE")

	return username, userpass, database
}
