package main

import (
	"fmt"
	"math"
)

// If
func sqrt(x float64) string {
	if (x < 0) {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If with a short statement
func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}