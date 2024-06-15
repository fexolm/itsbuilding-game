package game

import (
	"github.com/fexolm/itsbuilding-game/src/assets"
	"github.com/hajimehoshi/ebiten/v2"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Player struct {
	pos      gmath.Vec
	waypoint gmath.Vec

	sprite    *graphics.Sprite
	animation *Animation
}

func NewPlayer() *Player {
	p := &Player{
		pos: gmath.Vec{},
	}

	p.sprite = assets.OpenSprite("characters/worker_player.png")
	p.sprite.SetFrameWidth(p.sprite.ImageWidth() / 4)
	p.sprite.SetFrameHeight(p.sprite.ImageHeight() / 4)
	p.sprite.Pos.Base = &p.pos

	p.animation = &Animation{}
	p.animation.SetSprite(p.sprite, 4)
	p.animation.SetFPS(10)
	p.animation.repeated = true
	p.animation.numAnimations = 4

	return p
}

func (p *Player) Init(s *gscene.Scene) {
	s.AddGraphics(p.sprite, 0)
}

func (p *Player) Update(delta float64) {
	if !p.pos.EqualApprox(p.waypoint) {
		p.animation.Tick(delta)
	}

	const SPEED = 50

	p.pos = p.pos.MoveTowards(p.waypoint, SPEED*delta)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		p.waypoint = gmath.Vec{X: float64(x), Y: float64(y)}
	}

	if p.pos.X < p.waypoint.X {
		p.animation.SetOffsetY(1)
	} else if p.pos.X > p.waypoint.X {
		p.animation.SetOffsetY(2)
	} else if p.pos.Y < p.waypoint.Y {
		p.animation.SetOffsetY(0)
	} else if p.pos.Y > p.waypoint.Y {
		p.animation.SetOffsetY(3)
	}
}

func (p *Player) IsDisposed() bool {
	return p.sprite.IsDisposed()
}

func (p *Player) Dispose() {
	p.sprite.Dispose()
}
