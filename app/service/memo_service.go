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

type MemoService struct{}

// メモ全取得
func (ms MemoService) GetAll() ([]entity.Memo, error) {
	var m []entity.Memo
	url := "http://echo-login-app-api:8081/memo"

	// APIから取得
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return m, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return m, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return m, err
	}

	return m, nil
}

// ユーザーIDからメモの取得
func (ms MemoService) GetByUserID(user_id int) ([]entity.Memo, error) {
	var m []entity.Memo
	sid := strconv.Itoa(user_id)
	url := "http://echo-login-app-api:8081/memo/user_id/" + sid

	// APIから取得
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return m, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return m, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return m, err
	}

	return m, nil
}

// IDからメモの取得
func (ms MemoService) GetByID(id int) (entity.Memo, error) {
	var m entity.Memo
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/memo/id/" + sid

	// APIから取得
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return m, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return m, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return m, err
	}

	return m, nil
}

// メモ作成処理
func (ms MemoService) Create(title, content string, user_id int) error {
	var m entity.Memo
	url := "http://echo-login-app-api:8081/memo"

	m.Title = title
	m.Content = content
	m.User_ID = user_id

	// GoのデータをJSONに変換
	j, _ := json.Marshal(m)

	// apiへのユーザー情報送信
	req, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(j),
	)
	if err != nil {
		log.Printf("error http.POST: %v", err)
		return err
	}
	defer req.Body.Close()

	return nil
}

// // ユーザー名の変更処理
// func (us UserService) ChangeName(id int, username string) error {
// 	var u entity.User
// 	sid := strconv.Itoa(id)
// 	url := "http://echo-login-app-api:8081/user/" + sid

// 	u.Name = username

// 	// GoのデータをJSONに変換
// 	j, _ := json.Marshal(u)

// 	// apiへのユーザー情報送信
// 	req, err := http.NewRequest(
// 		"PUT",
// 		url,
// 		bytes.NewBuffer(j),
// 	)
// 	if err != nil {
// 		log.Printf("error http.PUT: %v", err)
// 		return err
// 	}

// 	// Headerセット
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	re, err := client.Do(req)
// 	if err != nil {
// 		log.Printf("error http.client.Do: %v", err)
// 		return err
// 	}
// 	defer re.Body.Close()

// 	return nil
// }

// // パスワード変更処理
// func (us UserService) ChangePassword(id int, password string) error {
// 	var u entity.User
// 	sid := strconv.Itoa(id)
// 	url := "http://echo-login-app-api:8081/user/" + sid

// 	u.Password = password

// 	// GoのデータをJSONに変換
// 	j, _ := json.Marshal(u)

// 	// apiへのユーザー情報送信
// 	req, err := http.NewRequest(
// 		"PUT",
// 		url,
// 		bytes.NewBuffer(j),
// 	)
// 	if err != nil {
// 		log.Printf("error http.PUT: %v", err)
// 		return err
// 	}

// 	// Headerのセット
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	re, err := client.Do(req)
// 	if err != nil {
// 		log.Printf("error http.client.Do: %v", err)
// 		return err
// 	}
// 	defer re.Body.Close()

// 	return nil
// }

// メモ削除処理
func (ms MemoService) Delete(id int) error {
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/memo/" + sid

	// apiへのユーザー情報送信
	req, err := http.NewRequest(
		"DELETE",
		url,
		nil,
	)
	if err != nil {
		log.Printf("error http.DELETE: %v", err)
		return err
	}
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return err
	}
	defer re.Body.Close()

	return nil
}