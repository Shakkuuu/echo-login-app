package main

import (
	"echo-login-app/app/server"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	resp *http.Response
	err  error
)

func main() {
	// 環境変数読み込み
	sk := loadEnv()
	ConnectCheck()
	server.Init(sk)
}

// 環境変数読み込み
func loadEnv() string {
	// Docker-compose.ymlでDocker起動時に設定した環境変数の取得
	session_key := os.Getenv("SESSION_KEY")

	return session_key
}

func ConnectCheck() {
	url := "http://echo-login-app-api:8081/"

	// 接続できるまで一定回数リトライ
	count := 0
	resp, err = http.Get(url)
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
				log.Printf("api connect error: %v\n", err)
				panic(err)
			}
			resp, err = http.Get(url)
		}
	}
	if resp.StatusCode != 200 {
		for {
			if resp.StatusCode == 200 {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 { // countgaが180になるまでリトライ
				fmt.Println("")
				log.Printf("api connect error: %v status: %v\n", err, resp.Status)
				panic(err)
			}
			resp, err = http.Get(url)
		}
	}
}
