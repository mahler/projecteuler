package main

import (
	"fmt"
	"math"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// https://projecteuler.net/problem=123
func main() {
	target := int64(10000000000)
	solution := int64(0)

	for n := int64(1); g(n) < target; n++ {
		solution = n + 1
	}

	fmt.Println("123/ Find the least value of n for which the remainder first exceeds 1010")
	fmt.Println(solution)
}

func isPrime(n int64) bool {

	if n == 1 {
		return false
	}

	end := int64(math.Sqrt(float64(n)))

	//If we start computing beyond the table, this is stupid
	for i := int64(1); Prime(i) <= end && i < primeTableLength; i++ {
		if n%Prime(i) == 0 {
			return false
		}
	}

	//If we need to pass the end of the prime table brute force
	if end > lastPrime {
		for i := int64(lastPrime); i <= end; i++ {
			if n%i == 0 {
				return false
			}
		}

	}

	return true
}

func Prime(n int64) int64 {

	if n < 1 {
		return 0
	}

	primeTable[1] = 2
	primeTable[2] = 3

	if n < primeTableLength && primeTable[n] != 0 {
		return primeTable[n]
	}

	i := Prime(n-1) + 1

	for !isPrime(i) {
		i++
	}

	if i < primeTableLength {
		primepilist[i] = n
	}

	if n < primeTableLength {
		primeTable[n] = i
	}
	return i
}

//Binomial theorem
func f(a, n int64) int64 {

	if n%2 == 0 {
		return 2
	}
	return (2 * n * a) % (a * a)
}

func g(n int64) int64 {
	return f(Prime(n), n)
}
