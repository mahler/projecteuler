package main

import (
	"fmt"
	"strconv"
)

// It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.
//
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.

func main() {
	var result int
	for i := 1; ; i++ {
		n, n2, n3, n4, n5, n6 := i, i*2, i*3, i*4, i*5, i*6
		if isPermutation(n, n2) && isPermutation(n, n3) && isPermutation(n, n4) && isPermutation(n, n5) && isPermutation(n, n6) {
			// First hit is the smallest number :)
			result = n
			break
		}
	}

	fmt.Println("052/ Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits?")
	fmt.Println(result)
}

func isPermutation(i, j int) bool {
	fti, ftj := map[rune]int{}, map[rune]int{}
	si, sj := strconv.Itoa(i), strconv.Itoa(j)
	if len(si) == len(sj) {
		for _, v := range si {
			fti[v] += 1
		}
		for _, v := range sj {
			ftj[v] += 1
		}
		for k, v := range fti {
			if ftj[k] != v {
				return false
			}
		}
		return true
	}
	return false
}
