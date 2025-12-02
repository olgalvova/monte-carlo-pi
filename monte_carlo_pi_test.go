package main

import (
  "fmt"
  "testing"
)

func TestCalculatePi(t *testing.T) {
  	_, piEstimate := calculatePi(10_000_000)
	if piEstimate < 3.14 || piEstimate > 3.15 {
	   t.Errorf("Wrong Pi: %f", piEstimate)
	}
	fmt.Printf("Pi: %f\n", piEstimate)
}
