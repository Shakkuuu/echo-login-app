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

type HasItemService struct{}

// 全ユーザーの取得済みアイテムリスト全取得
func (hs HasItemService) GetAll(token string) ([]entity.HasItem, error) {
	var hasitem []entity.HasItem
	url := "http://echo-login-app-api:8081/hasitem"

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return hasitem, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return hasitem, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return hasitem, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &hasitem); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return hasitem, err
	}

	return hasitem, nil
}

// ユーザーIDから取得済みアイテムリストの取得
func (hs HasItemService) GetByUserID(user_id int, token string) (entity.HasItem, error) {
	var hasitem entity.HasItem
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/hasitem/user_id/" + sid

	// APIから取得
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return hasitem, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return hasitem, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return hasitem, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &hasitem); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return hasitem, err
	}

	return hasitem, nil
}

// 所有済みアイテムリストの変更処理
func (hs HasItemService) Change(token string, user_id int, item entity.Item) error {
	// var hasitem entity.HasItem
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/hasitem/" + sid

	hasitem, _ := hs.GetByUserID(user_id, token)
	fmt.Printf("追加前:%v", hasitem)

	hasitem.Items = append(hasitem.Items, item)
	fmt.Printf("追加後:%v", hasitem)

	// GoのデータをJSONに変換
	j, _ := json.Marshal(hasitem)

	fmt.Printf("所有済みアイテムリスト変更:%v", hasitem)

	// apiへの情報送信
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewBuffer(j),
	)
	if err != nil {
		log.Printf("error http.PUT: %v", err)
		return err
	}

	// Headerセット
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return err
	}
	defer re.Body.Close()

	return nil
}
