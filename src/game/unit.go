package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Unit struct {
	pos  gmath.Vec
	path []gmath.Vec

	sprite    *graphics.Sprite
	animation *Animation
	m         *Map
}

const UnitSpeed = 50

func NewUnit(m *Map, sprite *graphics.Sprite) *Unit {
	u := &Unit{
		pos: gmath.Vec{},
	}

	u.sprite = sprite
	u.sprite.SetFrameWidth(u.sprite.ImageWidth() / 4)
	u.sprite.SetFrameHeight(u.sprite.ImageHeight() / 4)
	u.sprite.Pos.Base = &u.pos

	u.animation = &Animation{}
	u.animation.SetSprite(u.sprite, 4)
	u.animation.SetFPS(10)
	u.animation.repeated = true
	u.animation.numAnimations = 4

	u.m = m

	return u
}

func (u *Unit) SetWaypoint(pos gmath.Vec) {
	u.path = u.m.FindPath(u.pos, pos)
}

func (u *Unit) Init(s *gscene.Scene) {
	s.AddGraphics(u.sprite, 0)
}

func (u *Unit) Update(delta float64) {
	if len(u.path) != 0 {
		u.move(delta)
	}
}

func (u *Unit) move(delta float64) {
	if u.path[0].EqualApprox(u.pos) {
		u.path = u.path[1:]
	}

	if len(u.path) == 0 {
		return
	}

	u.animation.Tick(delta)

	wp := u.path[0]

	u.pos = u.pos.MoveTowards(wp, UnitSpeed*delta)

	if u.pos.X < wp.X {
		u.animation.SetOffsetY(1)
	} else if u.pos.X > wp.X {
		u.animation.SetOffsetY(2)
	} else if u.pos.Y < wp.Y {
		u.animation.SetOffsetY(0)
	} else if u.pos.Y > wp.Y {
		u.animation.SetOffsetY(3)
	}
}

func (u *Unit) IsDisposed() bool {
	return false
}
