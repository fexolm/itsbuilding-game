package game

import (
	"github.com/fexolm/itsbuilding-game/src/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gscene"
	"github.com/solarlune/ldtkgo"
	"image"
	"io"
)

type Map struct {
	sprite *graphics.Sprite
}

func (m *Map) Init(s *gscene.Scene) {
	s.AddGraphics(m.sprite, 0)
}

func (m *Map) Update(delta float64) {
}

func (m *Map) IsDisposed() bool {
	return false
}

type ltdkMapGenerator struct {
	proj           *ldtkgo.Project
	tilesets       map[string]*ebiten.Image
	currentTileset *ebiten.Image
	renderTarget   *ebiten.Image
}

func newLtdkMapGenerator(proj *ldtkgo.Project) *ltdkMapGenerator {
	tilesets := make(map[string]*ebiten.Image)
	return &ltdkMapGenerator{proj: proj, tilesets: tilesets}
}

func (g *ltdkMapGenerator) createMap() *Map {
	for _, tileset := range g.proj.Tilesets {
		_, exists := g.tilesets[tileset.Path]

		if !exists {
			img, _, err := ebitenutil.NewImageFromReader(assets.OpenAsset("maps" + "/" + tileset.Path))
			if err != nil {
				panic("Failed to read tileset from " + tileset.Path)
			}
			g.tilesets[tileset.Path] = img
		}
	}

	level := g.proj.Levels[0]

	g.renderTarget = ebiten.NewImage(960, 540)
	for layerIndex := len(level.Layers) - 1; layerIndex >= 0; layerIndex-- {
		layer := level.Layers[layerIndex]

		if layer.Tileset != nil && layer.Tileset.Path != "" {
			g.currentTileset = g.tilesets[layer.Tileset.Path]
			tileIndex := 0

			for _, tileData := range layer.AutoTiles {
				g.drawTile(tileData, tileIndex, layer)
				tileIndex++
			}
		}
	}
	s := graphics.NewSprite()
	s.SetImage(g.renderTarget)
	s.SetCentered(false)

	return &Map{sprite: s}
}

func (g *ltdkMapGenerator) drawTile(tileData *ldtkgo.Tile, tileIndex int, layer *ldtkgo.Layer) {
	tile := g.currentTileset.SubImage(image.Rect(tileData.Src[0], tileData.Src[1], tileData.Src[0]+layer.GridSize, tileData.Src[1]+layer.GridSize)).(*ebiten.Image)

	geoM := ebiten.GeoM{}

	geoM.Translate(float64(-layer.GridSize/2), float64(-layer.GridSize/2))

	if tileData.FlipX() {
		geoM.Scale(-1, 1)
	}
	if tileData.FlipY() {
		geoM.Scale(1, -1)
	}

	geoM.Translate(float64(layer.GridSize/2), float64(layer.GridSize/2))

	opt := ebiten.DrawImageOptions{}
	opt.GeoM = geoM

	opt.GeoM.Translate(float64(tileData.Position[0]+layer.OffsetX), float64(tileData.Position[1]+layer.OffsetY))
	g.renderTarget.DrawImage(tile, &opt)
}

func LoadMap(filename string) *Map {
	asset := assets.OpenAsset(filename)
	bytes, err := io.ReadAll(asset)
	if err != nil {
		panic("Cannot read asset")
	}

	proj, err := ldtkgo.Read(bytes)

	gen := newLtdkMapGenerator(proj)
	return gen.createMap()
}
