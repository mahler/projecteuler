package main

import (
	"fmt"
	"math/big"
)

// A googol (10100) is a massive number: one followed by one-hundred zeros; 100100 is almost unimaginably large:
// one followed by two-hundred zeros. Despite their size, the sum of the digits in each number is only 1.
//
// Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?

func DigitSum(num *big.Int) int {
	sum := 0
	str := num.String()
	for _, r := range str {
		sum += int(r - '0')
	}
	return sum
}

func main() {
	maxsum := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			big_a := big.NewInt(int64(a))
			result := big.NewInt(int64(a))
			for x := 1; x < b; x++ {
				result.Mul(result, big_a)
			}
			sum := DigitSum(result)
			if sum > maxsum {
				maxsum = sum
				//fmt.Println(a, b, sum)
			}
		}
	}
	fmt.Println("056/ Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?")
	fmt.Println(maxsum)
}
