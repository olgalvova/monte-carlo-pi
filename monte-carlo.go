package main

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"
)

type Function func(x float64) float64

type Range struct {
	min float64
	max float64
}

func anyFunc(funcName string, totalPoints int, min, max float64) {
	function, found := functions[funcName]
	if !found {
		log.Fatalf("Unknown function %s", funcName)
	}
	fmt.Printf("Starting Monte Carlo for function %s with %d points...\n", funcName, totalPoints)
	estimate := monteCarlo(function, Range{min, max}, totalPoints)
	fmt.Printf("Estimate: %f\n", estimate)
}

func monteCarlo(f Function, r Range, totalPoints int) float64 {
	pointsInside := 0
	width := r.max - r.min
	height := math.Max(f(r.max), f(r.min))
	for i := 0; i < totalPoints; i++ {
		x := rand.Float64()*width + r.min
		y := rand.Float64() * height

		if y <= f(x) {
			pointsInside++
		}
	}

	return width * height * float64(pointsInside) / float64(totalPoints)
}
