package game

import (
	"github.com/fexolm/itsbuilding-game/src/assets"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Laptop struct {
	pos          gmath.Vec
	laptopSprite *graphics.Sprite
	screenSprite *graphics.Sprite
}

func NewLaptop(pos gmath.Vec) *Laptop {
	obj := &Laptop{
		pos:          pos,
		laptopSprite: assets.OpenSprite("laptop.png"),
		screenSprite: assets.OpenSprite("screen/wallpaper1.png"),
	}
	obj.laptopSprite.Pos.Base = &obj.pos
	obj.screenSprite.Pos.Base = &obj.pos
	obj.screenSprite.Pos.Offset = gmath.Vec{X: 0., Y: -3.}
	return obj
}

func (p *Laptop) Init(s *gscene.Scene) {
	s.AddGraphics(p.laptopSprite, 0)
	s.AddGraphics(p.screenSprite, 0)
}

func (p *Laptop) Update(delta float64) {
}

func (p *Laptop) IsDisposed() bool {
	return p.laptopSprite.IsDisposed()
}

func (p *Laptop) Dispose() {
	p.laptopSprite.Dispose()
	p.screenSprite.Dispose()
}
