package main

import (
	"fmt"
)

// See https://projecteuler.net/problem=122
func main() {

	fmt.Println("122/ For 1 ≤ k ≤ 200, find ∑ m(k):")
	fmt.Println(int(1 / prob(0, 0, 15)))
}

func prob(won, drawn, rounds int) float64 {

	if drawn == rounds {
		if 2*won > drawn {
			return 1
		} else {
			return 0
		}
	}

	return prob(won+1, drawn+1, rounds)*1/float64(drawn+2) + prob(won, drawn+1, rounds)*float64(drawn+1)/float64(drawn+2)

}
