package main

import (
	"fmt"
)

// Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
//
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:
//
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 3 fractions between 1/3 and 1/2.
//
// How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d ≤ 12,000?
func main() {

	count := 0
	for den := 4; den <= 12000; den++ {
		bignum := den / 2
		smallnum := (den / 3) + 1

		for num := smallnum; num <= bignum; num++ {

			if greatestCommonDivisor(int64(num), int64(den)) == 1 {
				count++
			}
		}
	}

	fmt.Println("073/ How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d ≤ 12,000?")
	fmt.Println(count)
}

func greatestCommonDivisor(x, y int64) int64 {
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}

	if x > y {
		return x
	} else {
		return y
	}
}
