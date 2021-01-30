package main

import (
	"fmt"
	"math/big"
)

// A unit fraction contains 1 in the numerator. The decimal representation of the unit
// fractions with denominators 2 to 10 are given:
//
// 1/2	= 	0.5
// 1/3	= 	0.(3)
// 1/4	= 	0.25
// 1/5	= 	0.2
// 1/6	= 	0.1(6)
// 1/7	= 	0.(142857)
// 1/8	= 	0.125
// 1/9	= 	0.(1)
// 1/10	= 	0.1
// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
//
// Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.

// -----------------
// Using Fermat's little theorem to find the largest prime less than 1000
// http://en.wikipedia.org/wiki/Repeating_decimal#Fractions_with_prime_denominators

func main() {
	//find primes
	arr := make([]bool, 1000)
	arr[0], arr[1] = true, true
	prime := 3
	var k int
	for {
		for k = prime * 2; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
		} else {
			break
		}
	}

	//iterate over primes in reverse direction
	b := big.NewInt(0)
	var i int
	for k = prime; k > 0; k -= 2 {
		if !arr[k] && k%2 != 0 {
		inner:
			for i = 1; i < k; i++ {
				if b.Mod(b.Sub(b.Exp(big.NewInt(10), big.NewInt(int64(i)), nil), big.NewInt(1)), big.NewInt(int64(k))).Int64() == 0 {
					break inner
				}
			}
			if k-i == 1 {
				fmt.Println(k)
				return
			}
		}
	}
}
