package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=197
func main() {
	start := -1.
	seen := make(map[float64]bool)

	for !seen[start] {
		seen[start] = true
		start = f(start)
	}

	fmt.Println("197/")
	fmt.Println(start + f(start))
}

func f(x float64) float64 {
	return math.Floor(math.Pow(2, 30.403243784-(x*x))) / 1000000000
}
