package service

import (
	"echo-login-app/app/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ItemService struct{}

// アイテム全取得
func (is ItemService) GetAll(token string) ([]entity.Item, error) {
	var item []entity.Item
	url := "http://echo-login-app-api:8081/item"

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return item, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return item, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return item, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &item); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return item, err
	}

	return item, nil
}

// IDからアイテムの取得
func (is ItemService) GetByID(id int, token string) ([]entity.Item, error) {
	var item []entity.Item
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/item/" + sid

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return item, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return item, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return item, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &item); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return item, err
	}

	return item, nil
}
