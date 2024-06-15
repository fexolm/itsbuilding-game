package game

import (
	"github.com/fexolm/itsbuilding-game/src/assets"
	"github.com/hajimehoshi/ebiten/v2"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"image/png"
	"log"
)

type Player struct {
	pos    gmath.Vec
	sprite *graphics.Sprite
}

func NewPlayer() *Player {
	return &Player{
		pos: gmath.Vec{X: 5., Y: 6.},
	}
}

func (p *Player) Init(s *gscene.Scene) {
	file := assets.OpenAsset("commander.png")
	imgFile, err := png.Decode(file)
	defer file.Close()

	if err != nil {
		log.Fatalf("Failed to decode png: %v", err)
	}

	img := ebiten.NewImageFromImage(imgFile)

	p.sprite = graphics.NewSprite()
	p.sprite.SetImage(img)
	p.sprite.Pos.Base = &p.pos

	s.AddGraphics(p.sprite, 0)
}

func (p *Player) Update(delta float64) {
	const SPEED = 50

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.pos.X += SPEED * delta
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.pos.X -= SPEED * delta
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.pos.Y -= SPEED * delta
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.pos.Y += SPEED * delta
	}
}

func (p *Player) IsDisposed() bool {
	return p.sprite.IsDisposed()
}

func (p *Player) Dispose() {
	p.sprite.Dispose()
}
