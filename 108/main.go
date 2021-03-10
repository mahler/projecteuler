package main

import (
	"fmt"
)

// In the following equation x, y, and n are positive integers.
//
// [See https://projecteuler.net/problem=108]
//
// For n = 4 there are exactly three distinct solutions:
//
// [See https://projecteuler.net/problem=108]
//
// What is the least value of n for which the number of distinct solutions exceeds one-thousand?
//
// NOTE: This problem is an easier version of Problem 110; it is strongly advised that you solve this one first.
func main() {
	target := int64(4)

	counter := int64(0)
	record := int64(0)

	for target = 4; counter < 1000; target++ {

		counter = 0

		for den := target + 1; den <= 2*target; den++ {

			if (den*target)%(den-target) == 0 {
				counter++
			}

		}

		if counter > record {
			record = counter
			//fmt.Println(target, record)
		}

	}

	fmt.Println("108/ What is the least value of n for which the number of distinct solutions exceeds one-thousand?")
	fmt.Println(target - 1)
}
