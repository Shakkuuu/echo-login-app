package service

import (
	"bytes"
	"echo-login-app/backend/entity"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (us UserService) GetAll() ([]entity.User, error) {
	var u []entity.User
	url := "http://echo-login-app-api:8081/user"

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

	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return u, err
	}

	return u, nil
}

func (us UserService) GetByName(username string) (entity.User, error) {
	var u entity.User
	url := "http://echo-login-app-api:8081/user/username/" + username

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

	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		return u, err
	}

	return u, nil
}

func (us UserService) Login(id int, username, password string) error {
	var u entity.User
	sid := strconv.Itoa(id)
	url := "http://echo-login-app-api:8081/user/id/" + sid

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error http.Get: %v", err)
		err := fmt.Errorf("ログイン処理時にエラーが派生しました。")
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error io.ReadAll: %v", err)
		err := fmt.Errorf("ログイン処理時にエラーが派生しました。")
		return err
	}

	if err := json.Unmarshal(body, &u); err != nil {
		log.Printf("error json.Unmarshal: %v", err)
		err := fmt.Errorf("ログイン処理時にエラーが派生しました。")
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Printf("error bcrypt.CompareHashAndPassword: %v", err)
		err := fmt.Errorf("パスワードが一致していません。")
		log.Printf("パスワードチェック: %v", err)
		return err
	}
	// if password != u.Password {
	// 	err := fmt.Errorf("パスワードが一致していません。")
	// 	log.Printf("パスワードチェック: %v", err)
	// 	return err
	// }

	return nil
}

func (us UserService) Create(id int, username, password string) error {
	var u entity.User
	url := "http://echo-login-app-api:8081/user"

	u.ID = id
	u.Name = username
	u.Password = password

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
