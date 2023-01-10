package walker

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Walker interface {
	Walk()
	Draw(dst *ebiten.Image)
}

type WalkerType int

const (
	FourStep WalkerType = iota
	EightStep
	DownRight
)

type walker struct {
	size   float64
	typeOf WalkerType
	x      float64
	y      float64
	color  color.Color
}

// NewWalker returns a 4 steps walker. These walkers can move randomly up, down, left, or right.
func NewWalker(size, x, y float64, clr color.Color) Walker {
	return &walker{
		size:   size,
		typeOf: FourStep,
		x:      x,
		y:      y,
		color:  clr,
	}
}

// New8StepWalker walker that can walk in eight directions. These walkers can move randomly up, down, left, or right, up-right, down-right, up-left, down-left.
func New8StepWalker(size, x, y float64, clr color.Color) Walker {
	return &walker{
		size:   size,
		typeOf: EightStep,
		x:      x,
		y:      y,
		color:  clr,
	}
}

func NewDownRightWalker(size, x, y float64, clr color.Color) Walker {
	return &walker{
		size:   size,
		typeOf: DownRight,
		x:      x,
		y:      y,
		color:  clr,
	}
}

func (w *walker) Walk() {
	if w.typeOf == FourStep {
		dir := rand.Intn(4)
		switch dir {
		case 0:
			w.y--
		case 1:
			w.x++
		case 2:
			w.y++
		case 3:
			w.x--
		}
	}

	if w.typeOf == EightStep {
		// Generate a random number (-1, 0, 1)
		dirX := rand.Intn(3) - 1
		dirY := rand.Intn(3) - 1

		w.x += float64(dirX)
		w.y += float64(dirY)
	}

	if w.typeOf == DownRight {
		prob := rand.Float64()
		switch {
		// Probability to Move:
		// Up: 20%
		// Right: 30%
		// Down: 30%
		// Left: 20%
		case prob < .20:
			w.y--
		case prob > .20 && prob < .50:
			w.x++
		case prob > .50 && prob < .80:
			w.y++
		case prob > .80 && prob < 1:
			w.x--
		}
	}
}

func (w *walker) Draw(dst *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(w.x, w.y)

	ebitenutil.DrawCircle(dst, w.x, w.y, w.size, w.color)
}
