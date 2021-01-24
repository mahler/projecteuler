package main

import (
	"fmt"
	"math"
)

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
//
// Note: The prime finder is quite slow. It takes quite a while to find the number.

func main() {
	var primes []int

	for x := 1; x < 2000000; x++ {
		if isPrime(x) {
			primes = append(primes, x)
		}
	}
	// --------------------------------
	var primeSum int
	for _, p := range primes {
		primeSum += p
	}

	fmt.Println(primeSum)
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
