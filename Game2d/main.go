package main

import (
	"Game2s/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// endereco de memoria da struct
	// g := &game.Game{}
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
