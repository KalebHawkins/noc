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

type State int

const (
	Random State = iota
	Normal
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

type Circle struct {
	x, y, r float64
	color   color.Color
}

func (c *Circle) Draw(dst *ebiten.Image) {
	ebitenutil.DrawCircle(dst, c.x, c.y, c.r, c.color)
}

func (c *Circle) Update() {
	r := rand.NormFloat64()
	stdDev := 60.0
	mean := 320.0

	c.x = stdDev*r + mean
}

type Game struct {
	// 640 (scrWidth) / 32 (barWidth) = 20 bars.
	bars       [20]*Bar
	background *ebiten.Image
	State
	*Circle
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
		bars:       generateBars(),
		background: ebiten.NewImage(scrWidth, scrHeight),
		State:      0,
		Circle: &Circle{
			x:     scrWidth / 2,
			y:     scrHeight / 2,
			r:     10,
			color: color.RGBA{255, 255, 255, 25},
		},
	}
}

func (g *Game) Update() error {
	if g.State == Random {
		for _, b := range g.bars {
			b.Update()
		}
	}

	if g.State == Normal {
		g.Circle.Update()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.background.Fill(color.Black)
		g.State++
		if g.State > 1 {
			g.State = 0
		}
	}
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {

	if g.State == Random {
		g.background.Fill(color.Black)
		for _, b := range g.bars {
			b.Draw(g.background)
		}
	}

	if g.State == Normal {
		g.Circle.Draw(g.background)
	}

	scr.DrawImage(g.background, nil)

	ebitenutil.DebugPrint(scr, "Click to swap the animation.")

	switch g.State {
	case Random:
		ebitenutil.DebugPrintAt(scr, "Example: Random Distribution", 0, 20)
	case Normal:
		ebitenutil.DebugPrintAt(scr, "Example: Normal Distribution", 0, 20)
	}
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
