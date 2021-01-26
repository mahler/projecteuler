package main

import (
	"fmt"
	"strconv"
)

// Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
//
// 1634 = 1^4 + 6^4 + 3^4 + 4^4
// 8208 = 8^4 + 2^4 + 0^4 + 8^4
// 9474 = 9^4 + 4^4 + 7^4 + 4^4
// As 1 = 1^4 is not a sum it is not included.
//
// The sum of these numbers is 1634 + 8208 + 9474 = 19316.
//
// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

func main() {

	fiveNines := 9 * 9 * 9 * 9 * 9
	sum := 0
	for i := 2; i < 6*fiveNines; i++ {
		if isFifthPower(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("Sum:", sum)

}

func isFifthPower(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, digit := range s {
		d, _ := strconv.Atoi(string(digit))
		sum += d * d * d * d * d
	}
	if sum == n {
		return true
	}
	return false
}
