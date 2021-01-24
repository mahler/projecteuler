package main

import (
	"fmt"
	"strconv"
)

func main() {
	largestPalindrome := 0
	numberOne := 0
	numberTwo := 0
	for one := 999; one >= 100; one-- {
		for two := 999; two >= 100; two-- {
			product := one * two
			// It the product is smaller than the largest palindrome found, it's not a viable candidate.
			if largestPalindrome > product {
				break
			}
			if productIsPalindrome(product) {
				numberOne, numberTwo, largestPalindrome = one, two, product
			}

		}

	}
	fmt.Println(numberOne, "x", numberTwo, "=", largestPalindrome)
}

func productIsPalindrome(p int) bool {
	pStr := strconv.Itoa(p)

	for x := 0; x < len(pStr)/2; x++ {
		if pStr[x:x+1] != pStr[len(pStr)-x-1:len(pStr)-x] {
			return false
		}
	}
	// No palindrome test failed...
	return true
}
