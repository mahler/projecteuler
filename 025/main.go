package main

import (
	"fmt"
	"math/big"
)

// The Fibonacci sequence is defined by the recurrence relation:
//
// Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:
//
// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
//
// The 12th term, F12, is the first term to contain three digits.
//
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

func main() {
	var i int

	for {
		i++
		fib := fibonacciBig(i)
		//fmt.Println(i, ":", fib)

		if len(string(fib.String())) > 999 {
			fmt.Println(i, ":", fib)
			break
		}
	}

}

func fibonacciBig(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := int(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	return n1
}
