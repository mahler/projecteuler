package main

import (
	"fmt"
)

var memo = make(map[[2]int]int64)

// NOTE: This is a more difficult version of Problem 114.
//
// A row measuring n units in length has red blocks with a minimum length of m units placed on it,
// such that any two red blocks (which are allowed to be different lengths) are separated by at least one black square.
//
// Let the fill-count function, F(m, n), represent the number of ways that a row can be filled.
//
// For example, F(3, 29) = 673135 and F(3, 30) = 1089155.
//
// That is, for m = 3, it can be seen that n = 30 is the smallest value for which the fill-count function first
// exceeds one million.
//
// In the same way, for m = 10, it can be verified that F(10, 56) = 880711 and F(10, 57) = 1148904, so n = 57 is the
// least value for which the fill-count function first exceeds one million.
//
// For m = 50, find the least value of n for which the fill-count function first exceeds one million.

func main() {

	var answer int
	for n := 50; ways(n, 50) < 1000000; n++ {
		answer = n + 1
	}

	fmt.Println("115/ For m = 50, find the least value of n for which the fill-count function first exceeds one million?")
	fmt.Println(answer)
}

func ways(squares, minlength int) int64 {
	if squares < minlength {
		return 1
	}

	if answer, ok := memo[[2]int{squares, minlength}]; ok {
		return answer
	}

	total := int64(1) //The empty configuration

	for size := minlength; size <= squares; size++ {
		for start := 0; start <= squares-size; start++ {
			answer := int64(1)

			answer *= ways(squares-start-size-1, minlength)

			total += answer
		}

	}

	memo[[2]int{squares, minlength}] = total
	return total
}
