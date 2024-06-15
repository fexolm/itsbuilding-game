package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"image/png"
	"io"
	"log"
)

//go:embed all:_data
var gameAssets embed.FS

func OpenAsset(path string) io.ReadCloser {
	f, err := gameAssets.Open("_data/" + path)
	if err != nil {
		panic(err)
	}
	return f
}

func OpenSprite(path string) *graphics.Sprite {
	file := OpenAsset(path)
	imgFile, err := png.Decode(file)
	defer file.Close()

	if err != nil {
		log.Fatalf("Failed to decode png: %v", err)
	}

	img := ebiten.NewImageFromImage(imgFile)

	sprite := graphics.NewSprite()
	sprite.SetImage(img)

	return sprite
}
