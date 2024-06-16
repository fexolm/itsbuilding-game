package game

import (
	"fmt"
	"github.com/fexolm/itsbuilding-game/src/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"os"
)

type GameController struct {
	scene *gscene.Scene
}

func NewGameController() *GameController {
	return &GameController{}
}

func (c *GameController) Init(ctx gscene.InitContext) {
	c.scene = ctx.Scene
	m := LoadMap("maps/office.ldtk")
	c.scene.AddObject(m)
	c.scene.AddObject(NewStaticObject(gmath.Vec{X: 60., Y: 60.}, assets.OpenSprite("table.png")))
	c.scene.AddObject(NewLaptop(gmath.Vec{X: 60., Y: 50.}))
	c.scene.AddObject(NewPlayer(m))

	fmt.Println("GameController Init")
}

func (c *GameController) Update(delta float64) {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
