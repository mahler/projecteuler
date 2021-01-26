package main

//
// WARNING: The code in this version is easy to read, but very slow.
//

import "fmt"

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

	fmt.Println("VERY SLOW VERSION. See main.go for faster solution.")
	for {
		i++
		fib := fibonacci(i)
		//fmt.Println(i, ":", fib)

		if len(string(fib)) > 1000 {
			fmt.Println(i, ":", fib)
			break
		}
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
