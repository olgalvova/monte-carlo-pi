package main

import (
	"fmt"
	"math"
	"math/rand/v2"
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
func textMode(totalPoints int) {
	fmt.Printf("Starting Monte Carlo Pi estimation with %d points...\n", totalPoints)

	_, piEstimate := calculatePi(totalPoints)

	fmt.Printf("Estimated Pi: %f\n", piEstimate)
	fmt.Printf("Actual math.Pi: %f\n", math.Pi)

	// Calculate and display the error percentage
	errorPct := math.Abs(piEstimate-math.Pi) / math.Pi * 100.0
	fmt.Printf("Error: %.4f%%\n", errorPct)
}

