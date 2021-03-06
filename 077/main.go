package main

import (
	"fmt"
	"math"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

var table map[[2]int]int = make(map[[2]int]int)

func main() {
	answer := 0
	for i := 0; ways([2]int{i, 1}) < 5000; i++ {
		answer = 1 + i
	}
	fmt.Println(answer)
}

func ways(value [2]int) int {
	if solution, ok := table[value]; ok {
		return solution
	}

	n := value[0]
	k := value[1]
	solution := 0

	prime := int(Prime(int64(k)))

	if prime > n {
		solution = 0
	} else if prime == n {
		solution = 1
	} else {

		for prime < n {

			solution += ways([2]int{n - prime, k})

			k++
			prime = int(Prime(int64(k)))

		}

		if n == prime {
			solution++
		}

	}

	table[value] = solution
	return solution
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
