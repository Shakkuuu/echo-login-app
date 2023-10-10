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

const (
	ENMCOOLCHANGE int = 5
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

// ユーザーのステータス一覧の変更処理

// damage
func (ss StatusService) DamageChange(user_id, token string, now_status, up_status_count int) error {
	var status entity.Status
	url := "http://echo-login-app-api:8081/status/" + user_id

	status.Damage = now_status + up_status_count

	// GoのデータをJSONに変換
	j, _ := json.Marshal(status)

	// apiへのメモ情報送信
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

// hp
func (ss StatusService) HpChange(user_id, token string, now_status, up_status_count int) error {
	var status entity.Status
	url := "http://echo-login-app-api:8081/status/" + user_id

	status.Hp = now_status + up_status_count

	// GoのデータをJSONに変換
	j, _ := json.Marshal(status)

	// apiへのメモ情報送信
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

// shotspeed
func (ss StatusService) ShotSpeedChange(user_id, token string, now_status, up_status_count int) error {
	var status entity.Status
	url := "http://echo-login-app-api:8081/status/" + user_id

	status.ShotSpeed = now_status - up_status_count

	// GoのデータをJSONに変換
	j, _ := json.Marshal(status)

	// apiへのメモ情報送信
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

// enm_cool
func (ss StatusService) EnmCoolChange(user_id, token string, now_status, up_status_count int) error {
	var status entity.Status
	url := "http://echo-login-app-api:8081/status/" + user_id

	status.EnmCool = now_status + up_status_count*ENMCOOLCHANGE

	// GoのデータをJSONに変換
	j, _ := json.Marshal(status)

	// apiへのメモ情報送信
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

// score
func (ss StatusService) ScoreChange(user_id, token string, now_status, up_status_count int) error {
	var status entity.Status
	url := "http://echo-login-app-api:8081/status/" + user_id

	status.Score = now_status + up_status_count

	// GoのデータをJSONに変換
	j, _ := json.Marshal(status)

	// apiへのメモ情報送信
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
