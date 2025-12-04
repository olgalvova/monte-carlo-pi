package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCalculatePi(t *testing.T) {
	_, piEstimate := calculatePi(10_000_000)
	if piEstimate < 3.14 || piEstimate > 3.15 {
		t.Errorf("Wrong Pi: %f", piEstimate)
	}
	fmt.Printf("Pi: %f\n", piEstimate)
}

func TestMonteCarlo(t *testing.T) {
	estimate := monteCarlo(fx, Range{0, 10}, 100000000)
	if math.Round(estimate) != 50.0 {
		t.Errorf("Wrong integral of y = x: %f", estimate)
	}
}
