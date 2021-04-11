package main

import (
	"fmt"
	"math"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64
var sigtab map[uint64]uint64

// https://projecteuler.net/problem=211
func main() {
	sigtab = make(map[uint64]uint64)
	sigtab[1] = 1

	total := uint64(0)
	count := 1

	for i := uint64(1); i < 64000000; i++ {
		sig := sigma(i)

		if isSquare(sig) {
			// fmt.Println(count, ":", i, "\t\t\t", sig, "\t\t\t", total)
			count++
			total += i
		}
	}
	fmt.Println("211/ Divisor Square Sum")
	fmt.Println(total)
}

func sigma(n uint64) uint64 {
	if ans, ok := sigtab[n]; ok {
		return ans
	}

	top := uint64(math.Sqrt(float64(n))) + 1
	prime := uint64(2)

	for i := int64(1); prime <= top; i++ {
		prime = uint64(Prime(i))
		if n%prime == 0 {
			j := 0

			oldn := n

			for n%prime == 0 {
				n /= prime
				j++
			}

			adjust := uint64(1)
			factor := prime * prime

			for k := 1; k <= j; k++ {
				adjust += factor
				factor *= prime * prime
			}

			ans := adjust * sigma(n)

			sigtab[oldn] = ans
			return sigma(oldn)
		}

	}

	// n prime:
	sigtab[n] = n*n + 1

	return sigma(n)
}

func isSquare(n uint64) bool {
	sqrt := uint64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n || (sqrt+1)*(sqrt+1) == n {
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

	for i := int64(1); Prime(i) <= end && i < primeTableLength; i++ {
		if n%Prime(i) == 0 {
			return false
		}
	}

	if end > lastPrime {
		for i := int64(lastPrime); i <= end; i++ {
			if n%i == 0 {
				return false
			}
		}

	}

	return true
}
