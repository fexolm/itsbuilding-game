package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Unit struct {
	pos      gmath.Vec
	waypoint gmath.Vec

	sprite    *graphics.Sprite
	animation *Animation
}

const UnitSpeed = 50

func NewUnit(sprite *graphics.Sprite) *Unit {
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

	return u
}

func (u *Unit) Init(s *gscene.Scene) {
	s.AddGraphics(u.sprite, 0)
}

func (u *Unit) Update(delta float64) {
	if !u.pos.EqualApprox(u.waypoint) {
		u.animation.Tick(delta)
	}

	u.pos = u.pos.MoveTowards(u.waypoint, UnitSpeed*delta)

	if u.pos.X < u.waypoint.X {
		u.animation.SetOffsetY(1)
	} else if u.pos.X > u.waypoint.X {
		u.animation.SetOffsetY(2)
	} else if u.pos.Y < u.waypoint.Y {
		u.animation.SetOffsetY(0)
	} else if u.pos.Y > u.waypoint.Y {
		u.animation.SetOffsetY(3)
	}
}

func (u *Unit) IsDisposed() bool {
	return false
}
