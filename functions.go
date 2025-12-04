package main

import (
	"math"
)

var functions map[string]Function

func init() {
	functions = map[string]Function{
		"x":                fx,
		"10":               f10,
		"e^x":              exp,
		"1/x":              oneOverX,
		"sqrt(x)":          sqrt,
		"x^2":              sqr,
		"1/(x+1)(sqrt(x))": oneOverXPlusOneTimesSqrtX,
	}
}

// f(x) = 1/(x+1)(sqrt(x))
func oneOverXPlusOneTimesSqrtX(x float64) float64 {
	return 1 / (x + 1) * math.Pow(x, 0.5)
}

// f(x) = x. Integral is x^2 / 2
func fx(x float64) float64 {
	return x
}

// f(x) = k. Integral is k*x
func f10(x float64) float64 {
	return 10
}

// f(x) = e^x. Integral is the same.
func exp(x float64) float64 {
	return math.Exp(x)
}

// f(x) = 1/x. Integral is 1 in the interval [1, e]
func oneOverX(x float64) float64 {
	return 1 / x
}

func sqrt(x float64) float64 {
	return math.Pow(x, 0.5)
}

func sqr(x float64) float64 {
	return math.Pow(x, 2)
}
