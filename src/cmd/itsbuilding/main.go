package main

import (
	"github.com/fexolm/itsbuilding-game/src/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gscene"
)

type Game struct {
	windowWidth  int
	windowHeight int

	sceneManager *gscene.Manager
}

func (g *Game) start() {
	g.sceneManager.ChangeScene(game.NewGameController())
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

func main() {
	const (
		WIDTH  = 960
		HEIGHT = 540
	)

	g := &Game{
		windowWidth:  WIDTH,
		windowHeight: HEIGHT,
		sceneManager: gscene.NewManager(),
	}

	g.start()

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("It's Building!")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
