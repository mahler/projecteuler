package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove
// digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work
// from right to left: 3797, 379, 37, and 3.
//
// Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
//
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

func main() {
	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13

	sum := 0
	cnt := 0

	for cnt < 11 {
		maxP += 2
		if isPrime(prime, maxP) {
			prime = append(prime, maxP)
			str := strconv.Itoa(maxP)

			isTuncatable := true
			for i := 1; i < len(str); i++ {
				curr, _ := strconv.Atoi(str[i:])
				if prime[sort.SearchInts(prime, curr)] != curr {
					isTuncatable = false
					break
				}

				curr, _ = strconv.Atoi(str[:len(str)-i])
				if prime[sort.SearchInts(prime, curr)] != curr {
					isTuncatable = false
					break
				}
			}

			if isTuncatable {
				cnt++
				sum += maxP
			}
		}
	}

	fmt.Println("037/ Find the sum of the only eleven primes that are both truncatable from left to right and right to left")
	fmt.Println(sum)
}

func isPrime(prime []int, x int) bool {
	bound := int(math.Sqrt(float64(x)))
	if x%bound == 0 {
		return false
	}

	for _, p := range prime {
		if p >= bound {
			return true
		}
		if x%p == 0 {
			return false
		}
	}

	return true
}
