package main

import (
	"go-asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
	speed             = 2
)

var curAcceleration float64

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	posiion        Vector
	playerVelocity float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	return &Player{
		game:   game,
		sprite: sprite,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(-halfW, -halfH)
	opt.GeoM.Rotate(p.rotation)
	opt.GeoM.Translate(halfW, halfH)

	opt.GeoM.Translate(p.posiion.X, p.posiion.Y)

	screen.DrawImage(p.sprite, opt)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}

	p.accelerate()
}

func (p *Player) accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if curAcceleration < maxAcceleration {
			curAcceleration = p.playerVelocity + speed
		} else if curAcceleration >= maxAcceleration {
			curAcceleration = maxAcceleration
		}

		p.playerVelocity = curAcceleration

		// Move in the direction pointing
		dx := math.Sin(p.rotation) * curAcceleration
		dy := math.Cos(p.rotation) * -curAcceleration

		// Move the player
		p.posiion.X += dx
		p.posiion.Y += dy
	}
}
