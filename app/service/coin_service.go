package service

import (
	"bytes"
	"echo-login-app/app/entity"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CoinService struct{}

// コイン全取得
func (cs CoinService) GetAll(token string) ([]entity.Coin, error) {
	var coin []entity.Coin
	url := "http://echo-login-app-api:8081/coin"

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v\n", err)
		return coin, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return coin, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v\n", err)
		return coin, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &coin); err != nil {
		log.Printf("error json.Unmarshal: %v\n", err)
		return coin, err
	}

	return coin, nil
}

// ユーザーIDからコインの取得
func (cs CoinService) GetByUserID(user_id int, token string) (entity.Coin, error) {
	var coin entity.Coin
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/coin/user_id/" + sid

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v\n", err)
		return coin, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return coin, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v\n", err)
		return coin, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &coin); err != nil {
		log.Printf("error json.Unmarshal: %v\n", err)
		return coin, err
	}

	return coin, nil
}

// コインの変更処理
func (cs CoinService) ChangeQty(token string, user_id, qty int) error {
	var coin entity.Coin
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/coin/" + sid

	coin.Qty = qty

	// GoのデータをJSONに変換
	j, _ := json.Marshal(coin)

	fmt.Printf("コイン変更:%v\n", coin)

	// apiへのコイン情報送信
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewBuffer(j),
	)
	if err != nil {
		log.Printf("error http.PUT: %v\n", err)
		return err
	}

	// Headerセット
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return err
	}
	defer re.Body.Close()

	return nil
}
