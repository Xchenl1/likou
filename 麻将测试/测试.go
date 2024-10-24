package main

import (
	"fmt"
	"llikou/majiang"
)

func main() {
	game := majiang.Game{}
	game.Init()
	fmt.Println(game, len(game.Pai))
}
