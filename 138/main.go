package main

import (
	"fmt"
	"math"
)

const target = 12

// https://projecteuler.net/problem=138
func main() {
	count := 0
	answer := int64(0)

	for n := int64(1); count < target; n++ {

		MCans := make([]int64, 0)

		if sqrt, ok := IntSqrt(5*n*n + 1); ok {
			MCans = append(MCans, 2*n+sqrt)
		}
		if sqrt, ok := IntSqrt(5*n*n - 1); ok {
			MCans = append(MCans, 2*n+sqrt)
		}
		if sqrt, ok := IntSqrt(5*n*n + 2); ok {
			MCans = append(MCans, n+sqrt)
		}
		if sqrt, ok := IntSqrt(5*n*n - 2); ok {
			MCans = append(MCans, n+sqrt)
		}

		for _, m := range MCans {
			a, b, c := m*m-n*n, 2*m*n, m*m+n*n

			count++
			answer += c

			fmt.Println(a, b, c)
		}

	}

	fmt.Println("138/")
	fmt.Println(answer)
}

func IntSqrt(n int64) (sqrt int64, square bool) {
	sqrt = int64(math.Sqrt(float64(n)))
	square = sqrt*sqrt == n
	return
}
