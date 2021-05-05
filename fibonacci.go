package main

import (
	"math"
)

// O(2^n)
func fibonacci(n int) int {

	switch n {
	case 0:
		return 0

	case 1:
		return 1

	default:

		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func fastFib(n float64) int {
	return int((math.Pow(1+math.Sqrt(5), n) - (math.Pow(1-math.Sqrt(5), n))) / (math.Pow(2, n) * math.Sqrt(5)))

}
