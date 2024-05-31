package main

import (
	"Game2s/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {

	// endereco de memoria da struct
	g := &game.Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
