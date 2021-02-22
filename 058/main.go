package main

import (
	"fmt"
	"math"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.
//
// 37 36 35 34 33 32 31
// 38 17 16 15 14 13 30
// 39 18  5  4  3 12 29
// 40 19  6  1  2 11 28
// 41 20  7  8  9 10 27
// 42 21 22 23 24 25 26
// 43 44 45 46 47 48 49
//
// It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting is that 8 out of the 13 numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.
//
// If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed. If this process is continued, what is the side length of the square spiral for which the ratio of primes along both diagonals first falls below 10%?
func main() {
	fmt.Println("Hello, World", Prime(10000))

	primes := 3
	diagonal := 4
	var solution int
	for i := 2; float64(primes)/float64(diagonal) > .1; i++ {
		if IsPrime(int64(tl(i))) {
			primes++
		}
		if IsPrime(int64(tr(i))) {
			primes++
		}
		if IsPrime(int64(bl(i))) {
			primes++
		}
		if IsPrime(int64(br(i))) {
			primes++
		}
		diagonal += 4
		solution = 2*i + 1
	}
	fmt.Println("058/ what is the side length of the square spiral for which the ratio of primes along both diagonals first falls below 10%?")
	fmt.Println(solution)
}

func br(n int) int {
	return (1 + 2*n) * (1 + 2*n)
}

func bl(n int) int {
	return br(n) - 2*n
}

func tl(n int) int {
	return bl(n) - 2*n
}

func tr(n int) int {
	return tl(n) - 2*n
}

func IsPrime(n int64) bool {

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

	for !IsPrime(i) {
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
