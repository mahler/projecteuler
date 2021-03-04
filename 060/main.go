package main

import (
	"fmt"
	"math"
	"strconv"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastprime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them in any
// order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum
// of these four primes, 792, represents the lowest sum for a set of four primes with this property.
//
// Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.

func main() {

	ceiling := int64(2000)
	current := int64(99999999999)

	for i := int64(1); i < ceiling; i++ {

		for j := i; j < ceiling; j++ {

			if works(i, j) {

				for k := j; k < ceiling; k++ {

					if works(i, k) && works(j, k) {

						for l := k; l < ceiling; l++ {

							if works(i, l) && works(j, l) && works(k, l) {

								for a := l; a < ceiling; a++ {
									if works(a, i) && works(a, j) && works(a, k) && works(a, l) {
										sum := prime(a) + prime(i) + prime(j) + prime(k) + prime(l)
										if sum < current {
											//fmt.Println(prime(i), prime(j), prime(k), prime(l), prime(a), "Sum:", sum)
											current = sum

										}

									}

								}

							}

						}

					}
				}

			}

		}
	}

	fmt.Println("060/ Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.")
	fmt.Println(current)

}

func works(i int64, j int64) bool {

	if !isprime(concatanInt(prime(i), prime(j))) {
		return false
	}
	if !isprime(concatanInt(prime(j), prime(i))) {
		return false
	}

	return true
}

func isprime(n int64) bool {

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
	if end > lastprime {
		for i := int64(lastprime); i <= end; i++ {
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

	for !isprime(i) {
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

func concatanInt(a int64, b int64) int64 {

	//wrong string conversion
	x, _ := strconv.Atoi(strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10))
	return int64(x)
}
