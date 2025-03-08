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
	ScreenWidth       = 1280
	screenHeight      = 720
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

	// Center player on screen
	bounds := sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	return &Player{
		game:    game,
		sprite:  sprite,
		posiion: pos,
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
		p.keepsOnScreen()

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

func (p *Player) keepsOnScreen() {
	if p.posiion.X >= float64(ScreenWidth) {
		p.posiion.X = 0
	} else if p.posiion.X < 0 {
		p.posiion.X = ScreenWidth
	}

	if p.posiion.Y >= float64(screenHeight) {
		p.posiion.Y = 0
	} else if p.posiion.Y < 0 {
		p.posiion.Y = screenHeight
	}
}
