package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/maps"
	"log"
)

const (
	screenWidth  = 1000
	screenHeight = 1010
)

type Dot struct {
	x float32
	y float32
}

func main() {
	startingNumSamples := flag.Int("c", 1000, "Number of samples")
	deltaNumSamples := flag.Int("i", 1000, "Number of samples increment")
	minX := flag.Float64("n", 0, "Min x of the integration range")
	maxX := flag.Float64("x", 10, "Max x of the integration range")
	mode := flag.String("m", "pi", "Mode: pi, pi-text, func")
	function := flag.String("f", "x", fmt.Sprintf("%v", maps.Keys(functions)))
	flag.Parse()
	switch *mode {
	case "pi":
		ebiten.SetWindowSize(screenWidth, screenHeight)
		ebiten.SetWindowTitle("Calculating Pi with Monte Carlo")
		pi := makePi(*startingNumSamples, *deltaNumSamples)
		if err := ebiten.RunGame(pi); err != nil {
			log.Fatal(err)
		}
	case "pi-text":
		textModePi(*startingNumSamples)
	case "func":
		anyFunc(*function, *startingNumSamples, *minX, *maxX)
	default:
		log.Fatalf("Unknown mode %s", *mode)
	}
}
