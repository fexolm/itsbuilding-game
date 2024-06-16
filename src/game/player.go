package game

import (
	"github.com/fexolm/itsbuilding-game/src/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Player struct {
	u *Unit
}

func NewPlayer(m *Map) *Player {
	p := &Player{
		u: NewUnit(m, assets.OpenSprite("characters/worker_player.png")),
	}
	return p
}

func (p *Player) Update(delta float64) {
	p.u.Update(delta)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		p.u.SetWaypoint(gmath.Vec{X: float64(x), Y: float64(y)})
	}
}

func (p *Player) Init(s *gscene.Scene) {
	p.u.Init(s)
}

func (p *Player) IsDisposed() bool {
	return p.u.IsDisposed()
}
