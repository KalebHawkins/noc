// Exercise I.4
// Consider a simulation of paint splatter drawn as a collection of colored dots.
// Most of the paint clusters around a central location, but some dots do splatter
// out towards the edges. Can you use a normal distribution of random numbers to
// generate the locations of the dots? Can you also use a normal distribution
// of random numbers to generate a color palette?

package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	scrWidth  = 640
	scrHeight = 480
)

type Dot struct {
	x, y, r float64
	color   color.Color
}

func NewDot(x, y, r float64, clr color.Color) *Dot {
	return &Dot{
		x:     x,
		y:     y,
		r:     r,
		color: clr,
	}
}

func (d *Dot) Draw(dst *ebiten.Image) {
	ebitenutil.DrawCircle(dst, d.x, d.y, d.r, d.color)
}

func (d *Dot) Update(stdDev, mean float64) {
	d.x = stdDev*rand.NormFloat64() + mean
	d.y = stdDev*rand.NormFloat64() + mean

	red := stdDev*rand.NormFloat64() + mean
	grn := stdDev*rand.NormFloat64() + mean
	blu := stdDev*rand.NormFloat64() + mean

	d.color = color.RGBA{uint8(red), uint8(grn), uint8(blu), 10}
}

type Game struct {
	Dots       []*Dot
	StdDev     int32
	Mean       int32
	Background *ebiten.Image
}

func NewGame() *Game {
	dots := make([]*Dot, 100)
	for i := range dots {
		dots[i] = NewDot(0, 0, 5, color.RGBA{255, 255, 255, 1})
	}

	return &Game{
		Dots:       dots,
		StdDev:     45.0,
		Mean:       scrWidth / 2,
		Background: ebiten.NewImage(scrWidth, scrHeight),
	}
}

func (g *Game) Update() error {
	for _, d := range g.Dots {
		d.Update(float64(g.StdDev), float64(g.Mean))
	}
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	for _, d := range g.Dots {
		d.Draw(g.Background)
	}

	scr.DrawImage(g.Background, nil)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func main() {
	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("Paint Splatter")
	ebiten.SetScreenClearedEveryFrame(false)

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
