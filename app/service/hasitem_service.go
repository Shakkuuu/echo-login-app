package service

import (
	"bytes"
	"echo-login-app/app/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type HasItemService struct{}

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
		log.Printf("error http.Get: %v\n", err)
		return hasitem, err
	}
	// Headerセット
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v\n", err)
		return hasitem, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v\n", err)
		return hasitem, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &hasitem); err != nil {
		log.Printf("error json.Unmarshal: %v\n", err)
		return hasitem, err
	}

	return hasitem, nil
}

// 所有済みアイテムリストの追加処理
func (hs HasItemService) Add(token string, user_id int, result []entity.Item) error {
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/hasitem/" + sid

	hasitem, err := hs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("error hs.GetByUserID: %v\n", err)
		return err
	}

	var items []entity.Item
	var flag int = 0

	for _, item := range result {
		if item.ID <= 200 {
			for _, v := range hasitem.Items {
				if v.Name == item.Name {
					log.Println("既に持っているアタッチメントです。")
					flag = 1
					break
				}
			}
			for _, v2 := range items {
				if v2.Name == item.Name {
					log.Println("既に出現したアタッチメントです。")
					flag = 1
					break
				}
			}
			if flag == 0 {
				items = append(items, item)
				continue
			} else {
				flag = 0
				continue
			}
		}
		items = append(items, item)
	}

	posthasitem := entity.HasItem{Items: items, User_ID: user_id}

	// GoのデータをJSONに変換
	j, _ := json.Marshal(posthasitem)

	// apiへのユーザー情報送信
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(j),
	)
	if err != nil {
		log.Printf("error http.POST: %v\n", err)
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

// アイテムリストから削除する処理
func (hs HasItemService) Delete(id, token string, times int) error {
	// url := "http://echo-login-app-api:8081/hasitem/" + id
	url := "http://echo-login-app-api:8081/hasitem?item_id=" + id + "&times=" + strconv.Itoa(times)

	// apiへのユーザー情報送信
	req, err := http.NewRequest(
		"DELETE",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.DELETE: %v\n", err)
		return err
	}
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
