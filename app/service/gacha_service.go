package service

import (
	"echo-login-app/app/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type GachaService struct{}

// ガチャを引く
func (ms GachaService) Draw(times int, token string) ([]entity.Item, error) {
	var item []entity.Item
	stimes := strconv.Itoa(times)
	url := "http://echo-login-app-api:8081/gacha/" + stimes

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v\n", err)
		return item, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return item, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v\n", err)
		return item, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &item); err != nil {
		log.Printf("error json.Unmarshal: %v\n", err)
		return item, err
	}

	return item, nil
}
