package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// https://projecteuler.net/problem=65
func main() {

	numerator, _ := ctdFrac(eList(100))
	result := StringDigitSum(numerator)
	fmt.Println("065/ Find the sum of digits in the numerator of the 100th convergent of the continued fraction for E:")
	fmt.Println(result)

}

func ctdFrac(list []int) (num string, den string) {
	current := big.NewRat(int64(list[len(list)-1]), 1)

	for i := len(list) - 2; i >= 0; i-- {
		contrib := big.NewRat(int64(list[i]), 1)

		current.Inv(current)
		current.Add(current, contrib)
	}

	return current.Num().String(), current.Denom().String()
}

func eList(n int) []int {
	answer := make([]int, n)

	answer[0] = 2

	for i := 1; i < n; i++ {
		answer[i] = 1
	}

	for i := 0; 3*i+2 < n; i++ {
		answer[3*i+2] = 2 * (i + 1)

	}

	return answer
}

func StringDigitSum(str string) (total int) {
	for i := 0; i < len(str); i++ {
		is, _ := strconv.Atoi(str[i : i+1])
		total += is
	}
	return
}
