package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

// iniciar o game
func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// Func que cuida da parte de atualizar a logica do jogo
func (g *Game) Update() error {
	return nil
}

// func que cuida da parte de desenhar objetos na tela
func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!") --- código primeira versão
	g.player.Draw(screen)

}

// func respnsavel por retornar o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
