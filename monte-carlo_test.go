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
	} else {
		fmt.Printf("Pi: %f\n", piEstimate)
	}
}

func TestMonteCarloIntegralX(t *testing.T) {
	estimate := monteCarlo(fx, Range{0, 10}, 100000)
	if math.Round(estimate) != 50.0 {
		t.Errorf("Wrong integral of y = x: %f", estimate)
	} else {
		fmt.Println("y=x OK")
	}
}

func TestMonteCarloIntegral10(t *testing.T) {
	estimate := monteCarlo(f10, Range{0, 10}, 100000)
	if math.Round(estimate) != 100.0 { // XXX 0!
		t.Errorf("Wrong integral of y = 10: %f", estimate)
	} else {
		fmt.Println("y=10 OK")
	}
}

func TestMonteCarloIntegralExp(t *testing.T) {
	estimate := monteCarlo(exp, Range{0, 1}, 10_000_000)
	if math.Round(estimate) != 2.0 { // XXX 1?!
		t.Errorf("Wrong integral of y = e^x over [0,1]: %f", estimate)
	}

	estimate = monteCarlo(exp, Range{0, 3}, 10_000_000)
	if math.Round(estimate) != 19.0 { // XXX 16?!
		t.Errorf("Wrong integral of y = e^x over [0,3]: %f", estimate)
	}

	estimate = monteCarlo(exp, Range{0, 4}, 10_000_000)
	if math.Round(estimate) != 54.0 { // XXX 50?!
		t.Errorf("Wrong integral of y = e^x over [0, 4]: %f", estimate)
	}

	estimate = monteCarlo(exp, Range{0, 10}, 100_000_000)
	if math.Round(estimate) != 22025 { // XXX very close. Need to use % error instead of an exact comparison.
		t.Errorf("Wrong integral of y = e^x over [0, 4]: %f", estimate)
	} else {
		fmt.Println("y=e^x OK for [0 10]")
	}
}

func TestMonteCarloIntegralOneOverX(t *testing.T) {
	estimate := monteCarlo(oneOverX, Range{0.1, 2}, 100000)
	if math.Round(estimate) != 3.0 { // XXX 0?!
		t.Errorf("Wrong integral of y = 1/x: %f", estimate)
	} else {
		fmt.Println("y=1/x OK")
	}
}

func TestMonteCarloIntegralSqrt(t *testing.T) {
	estimate := monteCarlo(sqrt, Range{0, 2}, 100000)
	if math.Round(estimate) != 2.0 {
		t.Errorf("Wrong integral of y = sqrt(x): %f", estimate)
	} else {
		fmt.Println("y=sqrt(x) OK")
	}
}

func TestMonteCarloIntegralSqr(t *testing.T) {
	estimate := monteCarlo(sqr, Range{0, 2}, 100000)
	if math.Round(estimate) != 3.0 {
		t.Errorf("Wrong integral of y = x^2: %f over [0,2]", estimate)
	} else {
		fmt.Println("y=x^2 OK over [0,2]")
	}
}

func TestMonteCarloIntegralSqrNegative(t *testing.T) {
	estimate := monteCarlo(sqr, Range{-2, 2}, 100000)
	if math.Round(estimate) != 5.0 { // XXX 0?!
		t.Errorf("Wrong integral of y = x^2 over [-2,2]: %f", estimate)
	} else {
		fmt.Println("y=x^2 OK over [-2,2]")
	}
}

func TestMonteCarloIntegralPow3Positive(t *testing.T) {
	estimate := monteCarlo(pow3, Range{0, 2}, 10000)
	if math.Round(estimate) != 4 {
		t.Errorf("Wrong integral of y = x^3 over [0, 2]]: %f", estimate)
	} else {
		fmt.Println("y=x^3 OK for [0,2]")
	}
}

func TestMonteCarloIntegralPow3Negative(t *testing.T) {
	estimate := monteCarlo(pow3, Range{-2, 0}, 10000)
	if math.Round(estimate) != -4 {
		t.Errorf("Wrong integral of y = x^3 over [-2, 0]: %f. Expected -4.", estimate)
	} else {
		fmt.Println("y=x^3 OK for [-2,0]")
	}
}

func TestMonteCarloIntegralPow3Zero(t *testing.T) {
	estimate := monteCarlo(pow3, Range{-2, 2}, 10000)
	if math.Round(estimate) != 0 {
		t.Errorf("Wrong integral of y = x^3 over [-2, 2]: %f. Expected 0.", estimate)
	} else {
		fmt.Println("y=x^3 OK for [-2,2]")
	}
}
