package main

import (
	"go-asteroids/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	sprite *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	return &Player{sprite: sprite}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}

	screen.DrawImage(p.sprite, opt)
}

func (p *Player) Update() {

}
