package main

import "fmt"

// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example,
// the 5-digit number, 15234, is 1 through 5 pandigital.
//
// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier,
// and product is 1 through 9 pandigital.
//
// Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
//
// HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

func main() {
	pandigitalProducts := map[int]int{}
	for i := 2; i < 99; i++ {
		for j := i; j < 9999/i; j++ {
			if isPandigital(i, j, i*j) {
				pandigitalProducts[i*j]++
			}
		}
	}
	pandigitalSum := 0
	for k := range pandigitalProducts {
		pandigitalSum += k
	}

	fmt.Println("032/ the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital?")
	fmt.Println(pandigitalSum)
}

func isPandigital(a, b, c int) bool {
	full := fmt.Sprintf("%d%d%d", a, b, c)
	if len(full) != 9 {
		return false
	}
	set := map[rune]int{}
	for _, r := range full {
		if r == '0' || set[r] > 0 {
			return false
		}
		set[r] = 1
	}
	return true
}
