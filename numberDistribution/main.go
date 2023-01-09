package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	scrWidth  = 640
	scrHeight = 480
)

type Bar struct {
	x, y, w, h float64
}

func (b *Bar) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, b.x, b.y, b.w, -b.h, color.White)
}

func (b *Bar) Update() {
	r := rand.Float64()
	b.h += r
}

type Game struct {
	// 640 (scrWidth) / 32 (barWidth) = 20 bars.
	bars       [20]*Bar
	background *ebiten.Image
}

func generateBars() [20]*Bar {
	var bars [20]*Bar
	for i := 0; i < 20; i++ {
		bars[i] = &Bar{
			x: float64(i * 32),
			y: scrHeight,
			w: 30,
			h: 0,
		}
	}
	return bars
}

func NewGame() *Game {

	return &Game{
		background: ebiten.NewImage(scrWidth, scrHeight),
		bars:       generateBars(),
	}
}

func (g *Game) Update() error {
	for _, b := range g.bars {
		b.Update()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.bars = generateBars()
	}
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	g.background.Fill(color.Black)

	for _, b := range g.bars {
		b.Draw(g.background)
	}

	scr.DrawImage(g.background, nil)

	ebitenutil.DebugPrint(scr, "Click to restart the animation.\n\nExample: Normal (Uniform) Number Distribution")
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func main() {
	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("Random Number Distribution")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
