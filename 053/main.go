package main

import (
	"fmt"
	"math/big"
)

// See https://projecteuler.net/problem=53
// There are exactly ten ways of selecting three from five, 12345:
//
// 123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
//
// How many, not necessarily distinct, are greater than one-million?
const upperLimit = 100

func main() {
	// The size of the factorials requires bigInt to contain the numbers.
	factorials := make([]*big.Int, upperLimit+1)

	factorials[0] = big.NewInt(1)
	for i := 1; i <= upperLimit; i++ {
		var f big.Int
		f.Mul(factorials[i-1], big.NewInt(int64(i)))
		factorials[i] = &f
	}

	count := 0
	for n := 1; n <= upperLimit; n++ {
		for r := 1; r <= n; r++ {
			var c big.Int
			c.Set(factorials[n])
			c.Div(&c, factorials[r])
			c.Div(&c, factorials[n-r])
			if c.Cmp(big.NewInt(1000000)) > 0 {
				count++
			}
		}
	}

	fmt.Println("053/ How many, not necessarily distinct, are greater than one-million?")
	fmt.Println(count)
}
