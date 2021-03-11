package main

import (
	"fmt"
	"math"
)

const target = 1000000000
const factors = 40
const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

func main() {
	total := int64(0)
	counter := 0

	for i := int64(1); counter < factors; i++ {
		for Prime(i)%2 == 0 || Prime(i)%5 == 0 {
			i++
		}
		if a(int(Prime(i))) {
			counter++
			total += Prime(i)
		}
	}

	fmt.Println("132/ Find the sum of the first forty prime factors of R(10^9)")
	fmt.Println(total)
}

func a(n int) bool {
	remainder := 1
	i := 1
	for ; remainder != 0; i++ {
		remainder *= 10
		remainder++
		remainder %= n
	}
	if target%i == 0 {
		return true
	}
	return false
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
