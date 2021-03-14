package main

import (
	"fmt"
	"math"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

const target = 2000

// https://projecteuler.net/problem=128
func main() {
	var current, pd int

	seen := 2

	for i := 5; seen < target; i++ {
		if n := i / 2; i%2 == 0 {
			current = side(n)
			pd = PD(diff(side(n), soups(n)))
		} else {
			current = up(n)
			pd = PD(diff(up(n), oops(n)))
		}

		if pd == 3 {
			seen++
		}
	}

	fmt.Println("128/ The 2000th tile in this sequence")
	fmt.Println(current)
}

func diff(an int, list [6]int) [6]int {
	for i := 0; i < 6; i++ {
		if an-list[i] > 0 {
			list[i] = an - list[i]
		} else {
			list[i] = list[i] - an
		}
	}
	return list
}

func PD(list [6]int) (count int) {
	for i := 0; i < 6; i++ {
		if IsPrime(int64(list[i])) {
			count++
		}
	}
	return
}

func up(n int) int {
	return 3*n*n - 3*n + 2
}

func side(n int) int {
	return up(n) - 1
}

func oops(n int) [6]int {
	return [6]int{
		up(n + 1),
		side(n + 2),
		side(n + 1),
		up(n - 1),
		up(n) + 1,
		up(n+1) + 1,
	}
}

func soups(n int) [6]int {
	return [6]int{
		side(n + 1),
		side(n+1) - 1,
		side(n) - 1,
		side(n - 1),
		up(n - 2),
		up(n - 1),
	}
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
