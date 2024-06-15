package game

import (
	"fmt"
	"github.com/quasilyte/gscene"
)

type GameController struct {
	scene *gscene.Scene
}

func NewGameController() *GameController {
	return &GameController{}
}

func (c *GameController) Init(ctx gscene.InitContext) {
	c.scene = ctx.Scene

	c.scene.AddObject(NewPlayer())

	fmt.Println("GameController Init")
}

func (c *GameController) Update(delta float64) {
}
