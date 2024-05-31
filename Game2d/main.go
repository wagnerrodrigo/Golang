package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

// Func que cuida da parte de atualizar a logica do jogo
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// func que cuida da parte de desenhar objetos na tela
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

}

// func respnsavel por retornar o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	// endereco de memoria da struct
	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
