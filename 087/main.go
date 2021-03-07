package main

import (
	"fmt"
	"math"
)

const hat = 50000000
const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// The smallest number expressible as the sum of a prime square, prime cube, and prime fourth power is 28. In fact, there are exactly four numbers below fifty that can be expressed in such a way:
//
// 28 = 2^2 + 2^3 + 2^4
// 33 = 3^2 + 2^3 + 2^4
// 49 = 5^2 + 2^3 + 2^4
// 47 = 2^2 + 3^3 + 2^4
//
// How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?
func main() {
	table := [hat]int{}
	pre := int64(0)

	for i := int64(1); pre < hat; i++ {
		pre = Prime(i) * Prime(i) * Prime(i) * Prime(i)
		pri := int64(0)

		for j := int64(1); pri < hat; j++ {
			pri = pre + Prime(j)*Prime(j)*Prime(j)

			pro := int64(0)

			for k := int64(1); pro < hat; k++ {
				pro = pri + Prime(k)*Prime(k)

				if pro < hat {
					table[pro]++
				}
			}
		}
	}

	total := 0
	for i := 0; i < hat; i++ {
		if table[i] > 0 {
			total++
		}
	}

	fmt.Println("087/ How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?")
	fmt.Println(total)
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
