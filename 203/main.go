package main

import (
	"fmt"
	"math"
)

const (
	primeTableLength = 50000000 //+1!!!!!!!!!!!
	//lastPrime = Prime[primeTableLength - 1]
	lastPrime = 982451629
)

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// Squarefree Binomial Coefficients
// https://projecteuler.net/problem=203
func main() {
	row := int64(51)
	distinct := make(map[int64]bool)

	for n := int64(2); n < row; n++ {
		for k := int64(1); k < n; k++ {
			if isSquareFree(n, k) {
				distinct[choose(n, k)] = true
			}
		}
	}

	var total int64
	for key := range distinct {
		total += key
	}
	fmt.Println("203/ the sum of the distinct squarefree numbers in the first 51 rows of Pascal's triangle")
	fmt.Println(total + 1)
}

func isSquareFree(n, k int64) bool {
	factors := findfactors(choose(n, k))
	for i := 0; i < len(factors); i++ {
		if factors[i][1] >= 2 {
			return false
		}
	}
	return true
}

func choose(N, K int64) int64 {
	factors := make(map[int64]int64)

	if K == 0 || N == K || N <= 1 {
		return 1
	}

	if N < K {
		return 0
	}

	for n := N; n > N-K; n-- {
		nfactors := findfactors(n)
		for i := 0; i < len(nfactors); i++ {
			factors[nfactors[i][0]] += nfactors[i][1]
		}
	}

	for k := K; k >= 2; k-- {

		kfactors := findfactors(k)
		for i := 0; i < len(kfactors); i++ {
			factors[kfactors[i][0]] -= kfactors[i][1]
		}

	}

	answer := int64(1)

	for prime, multiplicity := range factors {

		for i := int64(0); i < multiplicity; i++ {
			answer *= prime
		}

	}
	return answer
}

func findfactors(n int64) [][2]int64 {
	factorList := factor(n)
	factors := [][2]int64{[2]int64{factorList[0], 1}}
	for i := 1; i < len(factorList); i++ {
		if factorList[i] == factors[len(factors)-1][0] {
			factors[len(factors)-1][1]++
		} else {
			factors = append(factors, [2]int64{factorList[i], 1})
		}
	}
	return factors
}

func factor(n int64) []int64 {
	var answer = []int64{}

	current := n

	i := int64(1)
	for !isPrime(current) {
		if current%prime(i) == 0 {
			answer = append(answer, prime(i))
			current = current / prime(i)
			i = 0
		}
		i++
	}

	answer = append(answer, current)

	return answer
}

func isPrime(n int64) bool {

	if n == 1 {
		return false
	}

	end := int64(math.Sqrt(float64(n)))

	//If we start computing beyond the table, this is stupid
	for i := int64(1); prime(i) <= end && i < primeTableLength; i++ {
		if n%prime(i) == 0 {
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

func prime(n int64) int64 {

	if n < 1 {
		return 0
	}

	primeTable[1] = 2
	primeTable[2] = 3

	if n < primeTableLength && primeTable[n] != 0 {
		return primeTable[n]
	}

	i := prime(n-1) + 1

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
