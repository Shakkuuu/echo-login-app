package service

import (
	"echo-login-app/app/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type StatusService struct{}

// メモ全取得
func (ss StatusService) GetAll(token string) ([]entity.Status, error) {
	var status []entity.Status
	url := "http://echo-login-app-api:8081/status"

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v\n", err)
		return status, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return status, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v\n", err)
		return status, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &status); err != nil {
		log.Printf("error json.Unmarshal: %v\n", err)
		return status, err
	}

	return status, nil
}

// ユーザーIDからユーザーのステータス一覧の取得
func (ss StatusService) GetByUserID(user_id int, token string) (entity.Status, error) {
	var status entity.Status
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/status/user_id/" + sid

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v\n", err)
		return status, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return status, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v\n", err)
		return status, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &status); err != nil {
		log.Printf("error json.Unmarshal: %v\n", err)
		return status, err
	}

	return status, nil
}

// // ユーザーのステータス一覧の変更処理
// func (ss StatusService) Change(user_id, title, content, token string) error {
// 	var status entity.Status
// 	url := "http://echo-login-app-api:8081/status/" + user_id

// 	m.Title = title
// 	m.Content = content

// 	// GoのデータをJSONに変換
// 	j, _ := json.Marshal(status)

// 	// apiへのメモ情報送信
// 	req, err := http.NewRequest(
// 		"PUT",
// 		url,
// 		bytes.NewBuffer(j),
// 	)
// 	if err != nil {
// 		log.Printf("error http.PUT: %v\n", err)
// 		return err
// 	}

// 	// Headerセット
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	client := &http.Client{}
// 	re, err := client.Do(req)
// 	if err != nil {
// 		log.Printf("error http.client.Do: %v\n", err)
// 		return err
// 	}
// 	defer re.Body.Close()

// 	return nil
// }
