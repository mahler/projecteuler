package main

import (
	"fmt"
)

const top = 1000000
const maxA = top/4 + 1

// Given the positive integers, x, y, and z, are consecutive terms of an arithmetic progression,
// the least value of the positive integer, n, for which the equation, x^2 − y^2 − z^2 = n, has
// exactly two solutions is n = 27:
//
// 34^2 − 27^2 − 20^2 = 12^2 − 9^2 − 6^2 = 27
//
// It turns out that n = 1155 is the least value which has exactly ten solutions.
//
// How many values of n less than one million have exactly ten distinct solutions?
func main() {
	count := make(map[int64]int)

	for a := int64(1); a < maxA; a++ {

		for c := 2*a - 1; c >= 0; c-- {

			n := 4*a*a - c*c

			if n > top {
				break
			}

			if a-c > 0 {
				count[n]++
			}
			if a+c > 0 && c != 0 {
				count[n]++

			}
		}

	}

	total := 0
	for _, mult := range count {

		if mult == 10 {

			total++
		}
	}

	fmt.Println("135/ How many values of n less than one million have exactly ten distinct solutions?")
	fmt.Println(total)
}
