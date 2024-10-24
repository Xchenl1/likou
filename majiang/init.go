package majiang

import (
	"math/rand"
	"strconv"
	"time"
)

// 条0 万1 饼2 东3 南4 西5 北6 白7 中8 发9

type Game struct {
	Pai  []string //牌库
	Hand []string //手牌
}

// Init 初始化牌库
func (game *Game) Init() {
	for i := 0; i < 3; i++ { //首先生成条万饼
		for j := 1; j <= 9; j++ {
			for k := 1; k <= 4; k++ {
				k1 := strconv.Itoa(i)
				k2 := strconv.Itoa(j)
				k3 := strconv.Itoa(k)
				game.Pai = append(game.Pai, k1+k2+k3)
			}
		}
	}
	//继续生成东南西北
	for i := 3; i <= 6; i++ {
		for j := 1; j <= 4; j++ {
			k1 := strconv.Itoa(i)
			k2 := strconv.Itoa(j)
			game.Pai = append(game.Pai, k1+k2)
		}
	}
	//生成百中发
	for i := 7; i <= 9; i++ { //首先生成条万饼
		for j := 1; j <= 4; j++ {
			k1 := strconv.Itoa(i)
			k2 := strconv.Itoa(j)
			game.Pai = append(game.Pai, k1+k2)
		}
	}
	//初始化手牌
	for i := 0; i < 13; i++ {
		game.InitHand()
	}
}

// InitHand  初始化手牌
func (game *Game) InitHand() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(game.Pai))
	game.Hand = append(game.Hand, game.Pai[index])
	game.Pai = append(game.Pai[:index], game.Pai[index+1:]...) //移除这个元素
}
