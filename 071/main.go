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
// It can be seen that 2/5 is the fraction immediately to the left of 3/7.
//
// By listing the set of reduced proper fractions for d ≤ 1,000,000 in ascending order of size, find the numerator of the fraction immediately to the left of 3/7.

func main() {
	d := 1000000
	if d%2 == 0 {
		d--
	}
	n := d/2 - 1

	for n*7 >= d*3 {
		n--
	}

	gcd := greatestCommonDivisor(d, n)
	ansN := n / gcd
	ansD := d / gcd

	fmt.Println("Fraction:", ansN, "/", ansD)
	fmt.Println()
	fmt.Println("071/ the numerator of the fraction immediately to the left of 3/7?")
	fmt.Println(ansN)
}

func greatestCommonDivisor(x, y int) int {
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
