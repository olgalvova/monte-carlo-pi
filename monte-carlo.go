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
	if r.max < r.min {
		log.Fatalf("Invalid min and max values: %f < %f", r.max, r.min)
	}
	pointsInsideNegative := 0
	pointsInsidePositive := 0
	// Find the box around the function grapth.
	width := r.max - r.min
	// TODO: it is not enough to find f at the two ends of the interval as the function may not be monotonous
	// We are going to assume it is for now.
	yMin := f(r.min)
	yMax := f(r.max)
	var height float64
	if yMax*yMin >= 0 {
		// same sign
		height = math.Max(math.Abs(yMax), math.Abs(yMin))
	} else {
		// different signs
		height = math.Abs(yMax) + math.Abs(yMin)
	}
	fmt.Printf("range: %v, width = %f, height = %f\n", r, width, height)
	for i := 0; i < totalPoints; i++ {
		x := rand.Float64()*width + r.min
		y := rand.Float64()*height + yMin
		fx := f(x)
		//fmt.Printf("x = %f, y = %f, f(x) = %f, ", x, y, fx)
		if f(x) >= 0 {
			if y > 0 && y <= fx {
				pointsInsidePositive++
				//fmt.Printf("inside positive\n")
			} else {
				//fmt.Printf("outside positive\n")
			}
		} else {
			// f(x) < 0
			if y < 0 && y >= fx {
				pointsInsideNegative++
				//fmt.Printf("inside negative\n")
			} else {
				//fmt.Printf("outside negative\n")
			}
		}
	}
	positiveIntegralPart := width * height * float64(pointsInsidePositive) / float64(totalPoints)
	negativeIntegralPart := width * height * float64(pointsInsideNegative) / float64(totalPoints)
	fmt.Printf("positiveIntegralPart = %f, negativeIntegralPart = %f\n", positiveIntegralPart, negativeIntegralPart)
	return positiveIntegralPart - negativeIntegralPart
}
