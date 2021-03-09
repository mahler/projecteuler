package main

import (
	"fmt"
)

// https://projecteuler.net/problem=86
func main() {
	seek := 1000000
	a, b := 2, 10000

	for b-a > 1 {

		c := (b + a) / 2

		if routes(c) >= seek {
			b = c
		} else {
			a = c
		}

	}

	fmt.Println("086/ Find the least value of M such that the number of solutions first exceeds one million.")
	fmt.Println(b)
}

func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func routes(M int) (total int) {

	maxm := M //could be more clever

	for n := 1; n < maxm; n++ {

		for m := n + 1; m < maxm; m += 2 {

			if greatestCommonDivisor(int64(m), int64(n)) == 1 {

				a := m*m - n*n
				b := 2 * m * n
				a, b = sort(a, b)

				//a is two sides case
				for k := 1; k*b <= M; k++ {
					total += (a * k) / 2
				}

				//b is two sides case
				if 2*a >= b {
					for k := 1; k*a <= M; k++ {
						total += ((2 * a * k) - (b * k) + 2) / 2
					}
				}
			}

		}

	}
	return
}

func greatestCommonDivisor(x, y int64) int64 {
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}

	if x > y {
		return x
	} else {
		return y
	}
}
