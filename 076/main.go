package main

import (
	"fmt"
)

var parttable = make(map[int]int)

// It is possible to write five as a sum in exactly six different ways:
//
// 4 + 1
// 3 + 2
// 3 + 1 + 1
// 2 + 2 + 1
// 2 + 1 + 1 + 1
// 1 + 1 + 1 + 1 + 1
//
// How many different ways can one hundred be written as a sum of at least two positive integers?
func main() {
	fmt.Println("076/ How many different ways can one hundred be written as a sum of at least two positive integers?")
	fmt.Println(IntPartitions(100) - 1)
}

func IntPartitions(n int) int {
	if ans, ok := parttable[n]; ok {
		return ans
	}

	if n == 0 {
		parttable[0] = 1
		return IntPartitions(0)
	}

	if n < 0 {
		return 0
	}

	sum := 0

	for k := 1; k <= n; k++ {
		var summand int
		if k%2 == 0 {
			summand = -1
		} else {
			summand = 1
		}

		summand *= IntPartitions(n-(k*(3*k-1)/2)) + IntPartitions(n-(k*(3*k+1)/2))

		sum += summand
	}

	parttable[n] = sum
	return IntPartitions(n)
}
