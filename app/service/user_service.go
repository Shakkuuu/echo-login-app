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

type UserService struct{}

// ユーザー全取得
func (us UserService) GetAll() ([]entity.User, error) {
	var u []entity.User
	url := "http://echo-login-app-api:8081/user"

	// APIから取得
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return u, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return u, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return u, err
	}

	return u, nil
}

// 名前からユーザーデータの取得
func (us UserService) GetByName(username string) (entity.User, error) {
	var u entity.User
	url := "http://echo-login-app-api:8081/user/username/" + username

	// APIから取得
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return u, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return u, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return u, err
	}

	return u, nil
}

// IDからユーザーデータの取得
func (us UserService) GetByID(id int) (entity.User, error) {
	var u entity.User
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/user/id/" + sid

	// APIから取得
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		return u, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return u, err
	}

	// JSONをGoのデータに変換
	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return u, err
	}

	return u, nil
}

// Login処理
func (us UserService) Login(id int, password string, pleasetoken bool) (*entity.Token, error) {
	var u entity.User
	url := "http://echo-login-app-api:8081/user/login"

	u.ID = id
	u.Password = password

	// GoのデータをJSONに変換
	j, _ := json.Marshal(u)

	// apiへのユーザー情報送信
	req, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(j),
	)
	if err != nil {
		log.Printf("error http.POST: %v", err)
		return nil, err
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		return nil, err
	}
	if req.StatusCode != 200 {
		log.Printf("error http.POST: %v", string(body))
		err = fmt.Errorf("パスワードが一致していません。")
		log.Printf("パスワードチェック: %v", err)
		return nil, err
	}
	defer req.Body.Close()

	if pleasetoken == true {
		var token entity.Token
		// JSONをGoのデータに変換
		if err := json.Unmarshal(body, &token); err != nil {
			log.Printf("error json.Unmarshal: %v", err)
			return nil, err
		}

		return &token, nil
	}

	return nil, nil
}

// ユーザー作成処理
func (us UserService) Create(id int, username, password string) error {
	var u entity.User
	url := "http://echo-login-app-api:8081/user"

	u.ID = id
	u.Name = username
	u.Password = password

	// GoのデータをJSONに変換
	j, _ := json.Marshal(u)

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

// ユーザー名の変更処理
func (us UserService) ChangeName(id int, username string) error {
	var u entity.User
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/user/" + sid

	u.Name = username

	// GoのデータをJSONに変換
	j, _ := json.Marshal(u)

	// apiへのユーザー情報送信
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
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return err
	}
	defer re.Body.Close()

	return nil
}

// パスワード変更処理
func (us UserService) ChangePassword(id int, password string) error {
	var u entity.User
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/user/" + sid

	u.Password = password

	// GoのデータをJSONに変換
	j, _ := json.Marshal(u)

	// apiへのユーザー情報送信
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewBuffer(j),
	)
	if err != nil {
		log.Printf("error http.PUT: %v", err)
		return err
	}

	// Headerのセット
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		log.Printf("error http.client.Do: %v", err)
		return err
	}
	defer re.Body.Close()

	return nil
}

// ユーザー削除処理
func (us UserService) Delete(id int) error {
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/user/" + sid

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
