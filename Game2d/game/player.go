package game

import (
	"Game2s/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image    *ebiten.Image
	position Vector
}

// returna um ponteiro de player
func NewPlayer() *Player {

	//cannot use image (variable of type []*ebiten.Image) as *ebiten.Image value in struct literalcompiler
	//O erro que acontecia ->> image := assets.PlanetsSprites é devido está tentando passar um slice de imagem em vez de apenas uma!
	// o que fiz foi grantir passar apenas uma imagem
	// PS. estava passando o desenho errado XD kkkkk

	// bounds := image.Bounds()

	position := Vector{
		X: (screenWidth / 2),
		Y: 500,
	}

	image := assets.PlayerSprite
	return &Player{
		image:    image,
		position: position,
	}
}

func (p *Player) Update() {}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(p.position.X, p.position.Y)
	// desenha imagem na tela
	screen.DrawImage(p.image, op)
}
