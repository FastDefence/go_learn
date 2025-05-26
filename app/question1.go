package main

import (
	"math"
)

func sqrt(n int) float64 {
	return math.Sqrt((float64(n)))
}

func pow(m int, n int) int {
	return int(math.Pow(float64(m), float64(n)))
}
