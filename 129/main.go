package main

import (
	"fmt"
)

const target = 1000000

// A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit of length k; for example, R(6) = 111111.
//
// Given that n is a positive integer and GCD(n, 10) = 1, it can be shown that there always exists a value, k, for which R(k) is divisible by n, and let A(n) be the least such value of k; for example, A(7) = 6 and A(41) = 5.
//
// The least value of n for which A(n) first exceeds ten is 17.
//
// Find the least value of n for which A(n) first exceeds one-million.
func main() {
	i := target
	for i%2 == 0 || i%5 == 0 {
		i++
	}

	for a(i) < target {
		i++
		for i%2 == 0 || i%5 == 0 {
			i++
		}
	}

	fmt.Println("129/ The least value of n for which A(n) first exceeds one-million")
	fmt.Println(i)
}

func a(n int) int {
	remainder := 1
	i := 1
	for ; remainder != 0; i++ {
		remainder *= 10
		remainder++
		remainder %= n
	}
	return i
}
