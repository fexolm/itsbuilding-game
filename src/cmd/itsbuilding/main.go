package main

import (
	"github.com/fexolm/itsbuilding-game/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	const (
		WIDTH  = 1920 / 2
		HEIGHT = 1080 / 2
	)

	g := game.NewGame(WIDTH, HEIGHT)
	g.Start()

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("It's Building!")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
