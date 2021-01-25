package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// n! means n × (n − 1) × ... × 3 × 2 × 1
//
// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
// Find the sum of the digits in the number 100!

func main() {
	n := 100
	res := big.NewInt(int64(n))

	for x := n - 1; x > 0; x-- {
		bigX := big.NewInt(int64(x))
		res = res.Mul(res, bigX)
	}
	fmt.Printf("020/ Sum of %d! is %s\n", n, res)

	var sum int
	strRes := fmt.Sprint(res)
	for x := 0; x < len(strRes); x++ {
		i, _ := strconv.Atoi(strRes[x : x+1])
		sum += i
	}
	fmt.Println("Sum of digits:", sum)
}
