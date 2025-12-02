package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	screenWidth  = 1000
	screenHeight = 1010
	textHeight   = 15
	dotRadius    = 2
)

// Game implements the ebitengine.Game interface.
type Game struct {
	circleX          float32
	circleY          float32
	totalPoints      int
	totalPointsDelta int
	radius           float32
	squareColor      color.RGBA
	circeColor       color.RGBA
	dotColor         color.RGBA
	updateSpeedHz    int
	iteration        int
	dots             []Dot
	pi               float64
	errorPct         float64
}

type Dot struct {
	x float32
	y float32
}

// Update updates the game state. This is called every frame (e.g., 60 times/second).
func (g *Game) Update() error {
	// Generate the next set of dots, bigger and bigger each time.
	if g.iteration%g.updateSpeedHz == 0 {
		g.dots, g.pi = calculatePi(g.totalPoints)
		g.errorPct = math.Abs(g.pi-math.Pi) / math.Pi * 100.0
		g.totalPoints += g.totalPointsDelta
	}
	g.iteration++
	return nil
}

// Draw draws the game screen. This is called according the frame rate in Display settings. Up to 120 times/sec on mac.
func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, g.circleX-g.radius, g.circleY-g.radius+textHeight, 2*g.radius, 2*g.radius, g.squareColor, false)
	vector.FillCircle(screen, g.circleX, g.circleY+textHeight, g.radius, g.circeColor, false)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Samples: %d, Estimated Pi: %f. Error: %.4f%%\n", len(g.dots), g.pi, g.errorPct))
	for i := 0; i < len(g.dots); i++ {
		dotX := g.dots[i].x * g.radius * 2
		dotY := g.dots[i].y*g.radius*2 + textHeight
		vector.FillCircle(screen, dotX, dotY, dotRadius, g.dotColor, false)
	}
}

// Layout takes the outside size (e.g. window size) and returns the intrinsic game screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	startingNumSamples := 1000
	deltaNumSamples := 1000
	mode := "g" // graphics
	var err error
	if len(os.Args) > 1 {
		mode = os.Args[1]
		if len(os.Args) > 2 {
			startingNumSamples, err = strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalf("Error reading staringNumSamples: %v", err)
			}
			if len(os.Args) > 3 {
				deltaNumSamples, err = strconv.Atoi(os.Args[3])
				if err != nil {
					log.Fatalf("Error reading deltaNumSamples: %v", err)
				}
			}
		}
	}
	switch mode {
	case "g":
		ebiten.SetWindowSize(screenWidth, screenHeight)
		ebiten.SetWindowTitle("Calculating Pi with Monte Carlo")

		game := &Game{
			circleX:          screenWidth / 2,
			circleY:          screenHeight / 2,
			radius:           screenHeight / 2,
			totalPoints:      startingNumSamples,
			totalPointsDelta: deltaNumSamples,
			squareColor:      color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, // blue
			circeColor:       color.RGBA{R: 0xff, G: 0xff, B: 0x0, A: 0xff},  // yellow
			dotColor:         color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, // green
			updateSpeedHz:    60,
		}
		if err := ebiten.RunGame(game); err != nil {
			log.Fatal(err)
		}
	case "t":
		textMode(startingNumSamples)
	default:
		log.Fatalf("Unknown mode %s", mode)
	}
}
