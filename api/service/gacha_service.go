package service

import (
	"echo-login-app/api/entity"
	"fmt"
	"math/rand"
	"time"
)

type GachaService struct{}

// ガチャを引く
func (gs GachaService) DrawGacha(times int, allitem []entity.Item) []entity.Item {
	var results []entity.Item
	var totalRatio int

	// 重みの合計
	for _, value := range allitem {
		totalRatio += value.Ratio
	}
	// 乱数生成
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	// 乱数生成された数字を超えたものをスライスに突っ込む
	for i := 0; i < times; i++ {
		val := r.Intn(totalRatio)
		fmt.Println(val)
		var nowRatio int
		for _, value := range allitem {
			nowRatio += value.Ratio
			if nowRatio > val {
				results = append(results, value)
				break
			}
		}
	}
	return results
}
