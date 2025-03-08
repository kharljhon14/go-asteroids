package main

import (
	"go-asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const rotationPerSecond = math.Pi

type Player struct {
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	return &Player{sprite: sprite}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(-halfW, -halfH)
	opt.GeoM.Rotate(p.rotation)
	opt.GeoM.Translate(halfW, halfH)

	screen.DrawImage(p.sprite, opt)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}
}
