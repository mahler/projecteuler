package main

import "fmt"

// The following iterative sequence is defined for the set of positive integers:
//
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following sequence:
//
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?

// NOTE: Once the chain starts the terms are allowed to go above one million.

func main() {
	var maxCollatz, startNumber int

	for n := 1; n <= 1000000; n++ {
		collatz := calcCollatz(n)

		if len(collatz) > maxCollatz {
			maxCollatz = len(collatz)
			startNumber = n
		}
	}

	fmt.Println("Max Collatz size:", maxCollatz)
	fmt.Println("Start number:", startNumber)
}

func calcCollatz(n int) []int {
	var collatz []int
	for {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		collatz = append(collatz, n)

		if n == 1 {
			break
		}
	}
	return collatz
}
