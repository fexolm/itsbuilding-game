package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type StaticObject struct {
	pos    gmath.Vec
	sprite *graphics.Sprite
}

func NewStaticObject(pos gmath.Vec, sprite *graphics.Sprite) *StaticObject {
	obj := &StaticObject{
		pos, sprite,
	}
	obj.sprite.Pos.Base = &obj.pos
	return obj
}

func (p *StaticObject) Init(s *gscene.Scene) {
	s.AddGraphics(p.sprite, 0)
}

func (p *StaticObject) Update(delta float64) {
}

func (p *StaticObject) IsDisposed() bool {
	return p.sprite.IsDisposed()
}

func (p *StaticObject) Dispose() {
	p.sprite.Dispose()
}
