package main

import (
	"fmt"
	"math"
)

// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1+2+3+4+5+6+7+8+9+10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .
// 3025 - 385 = 2640
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func main() {
	var min, max float64
	min, max = 1, 100
	var sqNat, xSum float64

	for x := min; x <= max; x++ {
		sqNat += math.Pow(x, 2)
		xSum += float64(x)
	}

	sqSum := math.Pow(xSum, 2)
	fmt.Println(int(sqSum))
	fmt.Println(int(sqNat))

	fmt.Println()
	fmt.Println("Result:", int(sqSum)-int(sqNat))
}
