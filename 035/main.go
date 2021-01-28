package main

import (
	"fmt"
	"strconv"
)

// The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
//
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
//
// How many circular primes are there below one million?

func main() {
	arr := make([]bool, 1000000)
	arr[1] = true
	prime := 3
	circularPrimes := 13
	var k, tmp, localCount int
	var str string

primeloop:
	for {
		for k = prime * 2; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			str = strconv.Itoa(prime)
			if prime > 100 {
				localCount = 1
				for i := 0; i < len(str)-1; i++ {
					str = str[1:] + str[0:1]
					tmp, _ = strconv.Atoi(str)
					if tmp > prime {
						// Not a circular prime
						continue primeloop
					} else if !arr[tmp] && tmp%2 != 0 {
						// Not a circular prime
						localCount++
					} else {
						continue primeloop
					}
				}
				circularPrimes += localCount
			}
		} else {
			break
		}
	}

	fmt.Println("035/ How many circular primes are there below one million?")
	fmt.Println(circularPrimes)
}
