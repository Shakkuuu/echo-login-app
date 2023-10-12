package service

import (
	"echo-login-app/api/db"
	"echo-login-app/api/entity"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// ユーザー全取得処理
func (us UserService) GetAll() ([]entity.User, error) {
	db := db.GetDB()
	var u []entity.User

	err := db.Find(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// ユーザー作成処理
func (us UserService) Create(u *entity.User) (*entity.User, error) {
	db := db.GetDB()

	err := db.Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// IDからのユーザー取得処理
func (us UserService) GetByID(id string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// 名前からのユーザー取得処理
func (us UserService) GetByName(username string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	err := db.Where("name = ?", username).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// IDからのユーザーデータ更新処理
func (us UserService) PutByID(u *entity.User, id string) (*entity.User, error) {
	db := db.GetDB()

	err := db.Where("id = ?", id).Model(&u).Updates(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

// IDからのユーザー削除処理
func (us UserService) Delete(id string) error {
	db := db.GetDB()

	var u entity.User

	err := db.Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}

// ユーザーログイン処理
func (us UserService) Login(u *entity.User) error {
	sid := strconv.Itoa(u.ID)
	// IDからのユーザー取得処理
	getu, err := us.GetByID(sid)
	if err != nil {
		return err
	}
	// ハッシュ化されたパスワードの解読と一致確認
	err = bcrypt.CompareHashAndPassword([]byte(getu.Password), []byte(u.Password))
	if err != nil {
		log.Printf("error bcrypt.CompareHashAndPassword: %v\n", err)
		err := fmt.Errorf("パスワードが一致していません。")
		log.Printf("パスワードチェック: %v\n", err)
		return err
	}
	return nil
}

// Token作成処理
func (us UserService) TokenCreate(id int) (string, error) {
	// Token発行用のアルゴリズムの指定
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 2).Unix(), // 2時間でToken失効
	}

	t, err := token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
	if err != nil {
		return "", err
	}
	// fmt.Printf("token: %v", t)
	return t, nil
}
