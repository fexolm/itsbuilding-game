package main

import (
	"github.com/fexolm/itsbuilding-game/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	const (
		WIDTH  = 800
		HEIGHT = 600
	)

	g := game.NewGame(WIDTH, HEIGHT)
	g.Start()

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("It's Building!")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
