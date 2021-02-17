package main

import (
	"fmt"
)

const nlid = 10000
const maxa = 1000

// See https://projecteuler.net/problem=120
func main() {
	total := 0

	for a := 3; a <= maxa; a++ {

		rmax := 0

		//This is not very smart, but f is very fast so...
		for n := 1; n < nlid; n++ {
			if f(a, n) > rmax {
				rmax = f(a, n)
			}
		}

		total += rmax

	}

	fmt.Println("120/ For 3 ≤ a ≤ 1000, find ∑ rmax:")
	fmt.Println(total)
}

//Binomial theorem
func f(a, n int) int {

	if n%2 == 0 {
		return 2
	}
	return (2 * n * a) % (a * a)
}
