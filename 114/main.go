package main

import (
	"fmt"
)

var memo = make(map[int]int64)

// A row measuring seven units in length has red blocks with a minimum length of three units placed on it,
// such that any two red blocks (which are allowed to be different lengths) are separated by at least one
// grey square. There are exactly seventeen ways of doing this.
//
// See https://projecteuler.net/problem=114
//
// How many ways can a row measuring fifty units in length be filled?
//
// NOTE: Although the example above does not lend itself to the possibility, in general it is permitted to
// mix block sizes. For example, on a row measuring eight units in length you could use red (3), grey (1),
// and red (4).

func main() {
	fmt.Println("114/ How many ways can a row measuring fifty units in length be filled?")
	fmt.Println(ways(50))
}

func ways(squares int) int64 {
	if squares < 3 {
		return 1
	}

	if answer, ok := memo[squares]; ok {
		return answer
	}

	total := int64(1) //The empty configuration

	for size := 3; size <= squares; size++ {
		for start := 0; start <= squares-size; start++ {
			answer := int64(1)

			answer *= ways(squares - start - size - 1)

			total += answer
		}

	}

	memo[squares] = total
	return total
}
