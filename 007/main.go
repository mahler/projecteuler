package main

import (
	"fmt"
	"math"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

func main() {
	var primeNumber, primeValue int

	for {
		primeValue++
		if isPrime(primeValue) {
			primeNumber++
		}
		if primeNumber == 10001 {
			break
		}
	}

	fmt.Println("007/", primeValue, " is prime number", primeNumber)
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
