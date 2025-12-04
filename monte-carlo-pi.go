package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	textHeight = 15
	dotRadius  = 2
)

// calculatePi uses the Monte Carlo method to estimate Pi with a given number of points.
func calculatePi(totalPoints int) ([]Dot, float64) {
	pointsInsideCircle := 0
	dots := make([]Dot, 0, totalPoints)
	for i := 0; i < totalPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()
		// Image a circle with radius 1.
		// Check if the point is inside the unit quarter circle using Piphagorian theorem x^2 + y^2 <= 1^2
		// Where 1 is the hypotenuse if the dot is on the circle
		// We use the distance formula (squared) to avoid a costly square root operation.
		if x*x+y*y <= 1.0 {
			pointsInsideCircle++
		}
		dots = append(dots, Dot{float32(x), float32(y)})
	}

	// The ratio of areas (circle/square) is pi/4.
	// We multiply the observed ratio by 4 to estimate Pi.
	piEstimate := 4.0 * float64(pointsInsideCircle) / float64(totalPoints)
	return dots, piEstimate
}

// To run in text mode without the graphics.
func textModePi(totalPoints int) {
	fmt.Printf("Starting Monte Carlo Pi estimation with %d points...\n", totalPoints)

	_, piEstimate := calculatePi(totalPoints)

	fmt.Printf("Estimated Pi: %f\n", piEstimate)
	fmt.Printf("Actual math.Pi: %f\n", math.Pi)

	// Calculate and display the error percentage
	errorPct := math.Abs(piEstimate-math.Pi) / math.Pi * 100.0
	fmt.Printf("Error: %.4f%%\n", errorPct)
}

// Implements the ebitengine.Game interface.
type Pi struct {
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

func makePi(startingNumSamples, deltaNumSamples int) *Pi {
	return &Pi{
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
}

// Update updates the game state. This is called every frame (e.g., 60 times/second).
func (g *Pi) Update() error {
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
func (g *Pi) Draw(screen *ebiten.Image) {
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
func (g *Pi) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
