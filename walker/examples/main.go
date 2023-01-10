package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/KalebHawkins/noc/walker"
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
	Traditional4Step State = iota
	Traditional8Step
	// A walker that has a tendency to walk down and to the right.
	// Exersice Introduction - I.1
	DownRight
)

type Game struct {
	State
	background *ebiten.Image
	walker.Walker
}

func NewGame() *Game {
	return &Game{
		State:      Traditional4Step,
		background: ebiten.NewImage(scrWidth, scrHeight),
		Walker:     walker.NewWalker(1, scrWidth/2, scrHeight/2, color.White),
	}
}

func (g *Game) Update() error {
	var stateChanged bool
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.State++
		if g.State > DownRight {
			g.State = Traditional4Step
		}

		stateChanged = true
	}

	if stateChanged {
		switch g.State {
		case Traditional4Step:
			g.background.Fill(color.Black)
			g.Walker = walker.NewWalker(1, scrWidth/2, scrHeight/2, color.White)
		case Traditional8Step:
			g.background.Fill(color.Black)
			g.Walker = walker.New8StepWalker(1, scrWidth/2, scrHeight/2, color.White)
		case DownRight:
			g.background.Fill(color.Black)
			g.Walker = walker.NewDownRightWalker(1, scrWidth/2, scrHeight/2, color.White)
		}
	}

	g.Walker.Walk()
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	g.Walker.Draw(g.background)
	scr.DrawImage(g.background, nil)

	var state string
	switch g.State {
	case 0:
		state = "Traditional 4 Step"
	case 1:
		state = "Traditional 8 Step"
	case 2:
		state = "Down Right Walker (Exercise I.1)"
	}

	ebitenutil.DebugPrint(scr, fmt.Sprintf("Click the Left Mouse Button to Change the State.\n\nState: %s", state))

	if g.State == DownRight {
		ebitenutil.DebugPrintAt(scr, "DownRight Walker has the following properties:\nProbability to move:\nUp: 20%\nRight: 30%\nDown: 30%\nLeft: 20%\n", 0, 60)
	}
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func main() {
	rand.Seed(time.Now().Unix())

	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("Walker")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
