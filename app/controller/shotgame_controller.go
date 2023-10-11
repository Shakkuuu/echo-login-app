package controller

import (
	"echo-login-app/app/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	RIGHTID     int = 101
	LEFTID      int = 102
	DOWNID      int = 103
	UPRIGHTID   int = 104
	UPLEFTID    int = 105
	DOWNRIGHTID int = 106
	DOWNLEFT    int = 107

	DAMAGEID    int = 201
	HPID        int = 202
	SHOTSPEEDID int = 203
	ENMCOOLID   int = 204
	SCOREID     int = 205

	DAMAGERATE    int = 30
	HPRATE        int = 20
	SHOTSPEEDRATE int = 25
	ENMCOOLRATE   int = 25
	SCORERATE     int = 10

	DAMAGELIMIT    int = 20
	HPLIMIT        int = 999
	SHOTSPEEDLIMIT int = 1
	ENMCOOLLIMIT   int = 600
	SCORELIMIT     int = 100
)

type ShotGameController struct{}

// GET Topページ表示
func (gc ShotGameController) Top(c echo.Context) error {
	var auc AuthController
	var ss service.StatusService
	var hs service.HasItemService

	// セッション
	user_id, err := auc.IDGetBySession(c)
	if err != nil {
		log.Printf("auc.IDGetSession error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Printf("TokenGet error: %v\n", err)
		m := map[string]interface{}{
			"message": "Tokenの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// ユーザーのステータス一覧取得
	status, err := ss.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ss.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーのステータス一覧の取得に失敗しました。",
			"status":  nil,
			"att":     nil,
		}
		return c.Render(http.StatusBadRequest, "shotgame.html", m)
	}

	// ユーザーの所持アイテム一覧取得
	hasitem, err := hs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("hs.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーの所持アイテムの取得に失敗しました。",
			"status":  status,
			"att":     nil,
		}
		return c.Render(http.StatusBadRequest, "shotgame.html", m)
	}

	atttf := map[string]int{"右アタッチメント": 0, "左アタッチメント": 0, "下アタッチメント": 0, "右上アタッチメント": 0, "左上アタッチメント": 0, "右下アタッチメント": 0, "左下アタッチメント": 0}

	for _, v := range hasitem.Items {
		if v.ID == RIGHTID || v.ID == LEFTID || v.ID == DOWNID || v.ID == UPRIGHTID || v.ID == UPLEFTID || v.ID == DOWNRIGHTID || v.ID == DOWNLEFT {
			atttf[v.Name] = 1
		}
	}

	m := map[string]interface{}{
		"message": "",
		"status":  status,
		"att":     atttf,
	}

	return c.Render(http.StatusOK, "shotgame.html", m)
}

// GET ステータス強化ページ表示
func (gc ShotGameController) StatusPage(c echo.Context) error {
	var auc AuthController
	var ss service.StatusService

	// セッション
	user_id, err := auc.IDGetBySession(c)
	if err != nil {
		log.Printf("auc.IDGetSession error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Printf("TokenGet error: %v\n", err)
		m := map[string]interface{}{
			"message": "Tokenの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// ユーザーのステータス一覧取得
	status, err := ss.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ss.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーのステータス一覧の取得に失敗しました。",
			"status":  nil,
		}
		return c.Render(http.StatusBadRequest, "status.html", m)
	}

	m := map[string]interface{}{
		"message": "",
		"status":  status,
	}

	return c.Render(http.StatusOK, "status.html", m)
}

// POST ステータス強化処理
func (sc ShotGameController) StatusUp(c echo.Context) error {
	var ss service.StatusService
	var auc AuthController
	var hs service.HasItemService

	// セッション
	user_id, err := auc.IDGetBySession(c)
	if err != nil {
		log.Printf("auc.IDGetSession error: %v\n", err)
		m := map[string]interface{}{
			"message": "セッションの取得に失敗しました。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	token, err := auc.TokenGet(c)
	if err != nil {
		log.Printf("TokenGet error: %v\n", err)
		m := map[string]interface{}{
			"message": "Tokenの取得に失敗しました。もう一度お試しください。",
		}
		return c.Render(http.StatusBadRequest, "login.html", m)
	}

	// ユーザーのステータス一覧取得
	status, err := ss.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ss.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "ユーザーのステータス一覧の取得に失敗しました。",
			"status":  nil,
		}
		return c.Render(http.StatusBadRequest, "status.html", m)
	}

	// htmlのformから値の取得
	s_damage := c.FormValue("damage")
	s_hp := c.FormValue("hp")
	s_shotspeed := c.FormValue("shotspeed")
	s_enmcool := c.FormValue("enmcool")
	s_score := c.FormValue("score")

	// 所持済みアイテム取得
	hasitem, err := hs.GetByUserID(user_id, token)
	if err != nil {
		log.Printf("ms.GetByUserID error: %v\n", err)
		m := map[string]interface{}{
			"message": "所持済みアイテムの取得に失敗しました。",
			"status":  status,
		}
		return c.Render(http.StatusBadRequest, "status.html", m)
	}

	// ステータス変更処理
	s_user_id := strconv.Itoa(user_id)
	// どのステータスを変更するか
	switch {
	case s_damage != "":
		{
			damage, err := strconv.Atoi(s_damage)
			if err != nil {
				log.Println("strconv.Atoi error")
				m := map[string]interface{}{
					"message": "強化度数を正しく入力してください。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			if status.Damage >= DAMAGELIMIT {
				log.Println("Damageをこれ以上強化できません。")
				m := map[string]interface{}{
					"message": "Damageをこれ以上強化できません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 素材が必要個数足りるか確認
			var count int = 0
			for _, item := range hasitem.Items {
				if item.ID == DAMAGEID {
					count++
				}
			}
			if count < DAMAGERATE*damage {
				log.Println("素材が足りません")
				m := map[string]interface{}{
					"message": "ダメージアップの素材が足りません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// ステータス変更処理
			err = ss.DamageChange(s_user_id, token, status.Damage, damage)
			if err != nil {
				log.Println("ss.DamageChange error")
				m := map[string]interface{}{
					"message": "ステータス変更時にエラーが発生しました。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 消費した素材の削除
			i := 0
			for i < DAMAGERATE*damage {
				err = hs.Delete(strconv.Itoa(DAMAGEID), token)
				if err != nil {
					log.Println("hs.Delete error")
					m := map[string]interface{}{
						"message": "素材の消費に失敗しました。",
						"status":  status,
					}
					return c.Render(http.StatusBadRequest, "status.html", m)
				}
				i++
			}
		}
	case s_hp != "":
		{
			hp, err := strconv.Atoi(s_hp)
			if err != nil {
				log.Println("strconv.Atoi error")
				m := map[string]interface{}{
					"message": "強化度数を正しく入力してください。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			if status.Hp == HPLIMIT {
				log.Println("Hpをこれ以上強化できません。")
				m := map[string]interface{}{
					"message": "Hpをこれ以上強化できません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 素材が必要個数足りるか確認
			var count int = 0
			for _, item := range hasitem.Items {
				if item.ID == HPID {
					count++
				}
			}
			if count < HPRATE*hp {
				log.Println("素材が足りません")
				m := map[string]interface{}{
					"message": "HPの素材が足りません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// ステータス変更処理
			err = ss.HpChange(s_user_id, token, status.Hp, hp)
			if err != nil {
				log.Println("ss.HpChange error")
				m := map[string]interface{}{
					"message": "ステータス変更時にエラーが発生しました。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 消費した素材の削除
			i := 0
			for i < HPRATE*hp {
				err = hs.Delete(strconv.Itoa(HPID), token)
				if err != nil {
					log.Println("hs.Delete error")
					m := map[string]interface{}{
						"message": "素材の消費に失敗しました。",
						"status":  status,
					}
					return c.Render(http.StatusBadRequest, "status.html", m)
				}
				i++
			}
		}
	case s_shotspeed != "":
		{
			shotspeed, err := strconv.Atoi(s_shotspeed)
			if err != nil {
				log.Println("strconv.Atoi error")
				m := map[string]interface{}{
					"message": "強化度数を正しく入力してください。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			if status.ShotSpeed == SHOTSPEEDLIMIT {
				log.Println("ShotSpeedをこれ以上強化できません。")
				m := map[string]interface{}{
					"message": "ShotSpeedをこれ以上強化できません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 素材が必要個数足りるか確認
			var count int = 0
			for _, item := range hasitem.Items {
				if item.ID == SHOTSPEEDID {
					count++
				}
			}
			if count < SHOTSPEEDRATE*shotspeed {
				log.Println("素材が足りません")
				m := map[string]interface{}{
					"message": "ShotSpeedの素材が足りません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// ステータス変更処理
			err = ss.ShotSpeedChange(s_user_id, token, status.ShotSpeed, shotspeed)
			if err != nil {
				log.Println("ss.ShotSpeedChange error")
				m := map[string]interface{}{
					"message": "ステータス変更時にエラーが発生しました。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 消費した素材の削除
			i := 0
			for i < SHOTSPEEDRATE*shotspeed {
				err = hs.Delete(strconv.Itoa(SHOTSPEEDID), token)
				if err != nil {
					log.Println("hs.Delete error")
					m := map[string]interface{}{
						"message": "素材の消費に失敗しました。",
						"status":  status,
					}
					return c.Render(http.StatusBadRequest, "status.html", m)
				}
				i++
			}
		}
	case s_enmcool != "":
		{
			enmcool, err := strconv.Atoi(s_enmcool)
			if err != nil {
				log.Println("strconv.Atoi error")
				m := map[string]interface{}{
					"message": "強化度数を正しく入力してください。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			if status.EnmCool == ENMCOOLLIMIT {
				log.Println("EnmCoolをこれ以上強化できません。")
				m := map[string]interface{}{
					"message": "EnmCoolをこれ以上強化できません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 素材が必要個数足りるか確認
			var count int = 0
			for _, item := range hasitem.Items {
				if item.ID == ENMCOOLID {
					count++
				}
			}
			if count < ENMCOOLRATE*enmcool {
				log.Println("素材が足りません")
				m := map[string]interface{}{
					"message": "EnmCoolの素材が足りません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// ステータス変更処理
			err = ss.EnmCoolChange(s_user_id, token, status.EnmCool, enmcool)
			if err != nil {
				log.Println("ss.EnmCoolChange error")
				m := map[string]interface{}{
					"message": "ステータス変更時にエラーが発生しました。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 消費した素材の削除
			i := 0
			for i < ENMCOOLRATE*enmcool {
				err = hs.Delete(strconv.Itoa(ENMCOOLID), token)
				if err != nil {
					log.Println("hs.Delete error")
					m := map[string]interface{}{
						"message": "素材の消費に失敗しました。",
						"status":  status,
					}
					return c.Render(http.StatusBadRequest, "status.html", m)
				}
				i++
			}
		}
	case s_score != "":
		{
			score, err := strconv.Atoi(s_score)
			if err != nil {
				log.Println("strconv.Atoi error")
				m := map[string]interface{}{
					"message": "強化度数を正しく入力してください。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			if status.Score == SCORELIMIT {
				log.Println("Scoreをこれ以上強化できません。")
				m := map[string]interface{}{
					"message": "Scoreをこれ以上強化できません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 素材が必要個数足りるか確認
			var count int = 0
			for _, item := range hasitem.Items {
				if item.ID == SCOREID {
					count++
				}
			}
			if count < SCORERATE*score {
				log.Println("素材が足りません")
				m := map[string]interface{}{
					"message": "Scoreの素材が足りません",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// ステータス変更処理
			err = ss.ScoreChange(s_user_id, token, status.Score, score)
			if err != nil {
				log.Println("ss.ScoreChange error")
				m := map[string]interface{}{
					"message": "ステータス変更時にエラーが発生しました。",
					"status":  status,
				}
				return c.Render(http.StatusBadRequest, "status.html", m)
			}

			// 消費した素材の削除
			i := 0
			for i < SCORERATE*score {
				err = hs.Delete(strconv.Itoa(SCOREID), token)
				if err != nil {
					log.Println("hs.Delete error")
					m := map[string]interface{}{
						"message": "素材の消費に失敗しました。",
						"status":  status,
					}
					return c.Render(http.StatusBadRequest, "status.html", m)
				}
				i++
			}
		}
	default:
		log.Println("どのステータスも入力されていません。")
		m := map[string]interface{}{
			"message": "どのステータスも入力されていません。",
			"status":  status,
		}
		return c.Render(http.StatusBadRequest, "status.html", m)
	}

	fmt.Println("ステータス変更成功したよ")
	return c.Redirect(http.StatusFound, "/app/game/shot/status")
}
