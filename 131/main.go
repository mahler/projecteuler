package main

import (
	"fmt"
	"math"
)

const search = 1000000
const hat = 50000000
const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// There are some prime values, p, for which there exists a positive integer, n, such that the
// xpression n^3 + n^2p is a perfect cube.
//
// For example, when p = 19, 8^3 + 8^2×19 = 12^3.
//
// What is perhaps most surprising is that for each prime with this property the value of n is
// unique, and there are only four such primes below one-hundred.
//
// How many primes below one million have this remarkable property?
func main() {
	count := 0

	for start := int64(1); difference(start, 1) <= search; start++ {
		for jump := int64(1); difference(start, jump) <= search; jump++ {
			if isPrime(difference(start, jump)) {
				count++
				//fmt.Println(start, jump, difference(start, jump))
			}
		}
	}

	fmt.Println("131/ How many primes below one million have this remarkable property?")
	fmt.Println(count)
}

func difference(n, offset int64) int64 {
	return (offset * offset * offset) + (3 * offset * offset * n) + (3 * offset * n * n)
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
