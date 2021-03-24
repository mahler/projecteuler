package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	lid    = 12
	target = 4000000
	top    = 4
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

func main() {
	fac := make([][2]int64, lid)
	for i := int64(1); i < lid+1; i++ {
		fac[i-1] = [2]int64{Prime(i), 0}
	}

	var smallest big.Int
	smallest.SetString("9999999999999999", 10)

	for fac[0][1] = 0; fac[0][1] < top; fac[0][1]++ {
		for fac[1][1] = fac[0][1]; fac[1][1] >= 0; fac[1][1]-- {
			for fac[2][1] = fac[1][1]; fac[2][1] >= 0; fac[2][1]-- {
				for fac[3][1] = fac[2][1]; fac[3][1] >= 0; fac[3][1]-- {
					for fac[4][1] = fac[3][1]; fac[4][1] >= 0; fac[4][1]-- {
						for fac[5][1] = fac[4][1]; fac[5][1] >= 0; fac[5][1]-- {
							for fac[6][1] = fac[5][1]; fac[6][1] >= 0; fac[6][1]-- {
								for fac[7][1] = fac[6][1]; fac[7][1] >= 0; fac[7][1]-- {
									for fac[8][1] = fac[7][1]; fac[8][1] >= 0; fac[8][1]-- {
										for fac[9][1] = fac[8][1]; fac[9][1] >= 0; fac[9][1]-- {
											for fac[10][1] = fac[9][1]; fac[10][1] >= 0; fac[10][1]-- {
												for fac[11][1] = fac[10][1]; fac[11][1] >= 0; fac[11][1]-- {
													if solns(fac) > target && smallest.Cmp(evaluate(fac)) > 0 {
														smallest = *evaluate(fac)
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
			}
		}
	}
	fmt.Println("110/ What is the least value of  for which the number of distinct solutions exceeds four million?")
	fmt.Println(smallest.String())
}

func evaluate(factors [][2]int64) *big.Int {
	show := big.NewInt(1)
	for i := 0; i < len(factors); i++ {
		if factors[i][1] > 0 {
			prime := big.NewInt(factors[i][0])

			for j := int64(0); j < factors[i][1]; j++ {
				show = show.Mul(show, prime)
			}
		}
	}
	return show
}

func solns(factors [][2]int64) int64 {
	answer := int64(1)

	for i := 0; i < len(factors); i++ {
		answer *= 1 + (2 * factors[i][1])
	}

	return answer/2 + 1
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
