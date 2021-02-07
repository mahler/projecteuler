package main

import (
	"fmt"
	"math"
)

// It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.
//
// 9 = 7 + 2×1^2
// 15 = 7 + 2×2^2
// 21 = 3 + 2×3^2
// 25 = 7 + 2×3^2
// 27 = 19 + 2×2^2
// 33 = 31 + 2×1^2
//
// It turns out that the conjecture was false.
//
// What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
func main() {
	primes := make([]bool, 1000000)
	primes[0], primes[1], primes[2] = false, false, true
	resultCandidate := 3
	var composite, found bool
	var tmp, tmp2 int
loop:
	for {
		composite = false
	inner:
		for j := 2; j < resultCandidate; j++ {
			if primes[j] && resultCandidate%j == 0 {
				composite = true
				break inner
			}
		}
		if composite {
			found = false
		inner2:
			for j := 2; j < resultCandidate; j++ {
				if primes[j] {
					tmp = (resultCandidate - j)
					if tmp2 = int(math.Sqrt(float64(tmp / 2))); tmp%2 == 0 && tmp2*tmp2 == tmp/2 {
						found = true
						break inner2
					}
				}
			}
			if !found {
				// If the test didn't fail, we got the result.
				break loop
			}
		} else {
			primes[resultCandidate] = true
		}
		// Next candidate...
		resultCandidate += 2
	}

	fmt.Println("046/ What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?")
	fmt.Println(resultCandidate)

}
