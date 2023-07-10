package main

import (
	"echo-login-app/backend/server"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	resp *http.Response
	err  error
)

func main() {
	ConnectCheck()
	server.Init()
}

func ConnectCheck() {
	// type PingCheck struct {
	// 	Status  int
	// 	Message string
	// }
	// var p PingCheck

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
	if resp.Status != "200 OK" {
		for {
			if resp.Status == "200 OK" {
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
