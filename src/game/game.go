package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gscene"
)

type Game struct {
	windowWidth  int
	windowHeight int

	sceneManager *gscene.Manager
}

func (g *Game) Start() {
	g.sceneManager.ChangeScene(NewGameController())
}

func (g *Game) Update() error {
	g.sceneManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.windowWidth, g.windowHeight
}

func NewGame(w, h int) *Game {
	return &Game{
		windowWidth:  w,
		windowHeight: h,
		sceneManager: gscene.NewManager(),
	}
}
